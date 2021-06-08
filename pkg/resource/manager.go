// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"fmt"
	"github.com/rode/es-index-manager/indexmanager"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/esutil"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering"
	"github.com/rode/rode/config"
	pb "github.com/rode/rode/proto/v1alpha1"
	"github.com/rode/rode/protodeps/grafeas/proto/v1beta1/common_go_proto"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate counterfeiter -generate

const (
	genericResourcesDocumentKind = "generic-resources"

	genericResourceDocumentJoinField   = "join"
	genericResourceRelationName        = "resource"
	genericResourceVersionRelationName = "version"
)

//counterfeiter:generate . Manager
type Manager interface {
	BatchCreateGenericResources(ctx context.Context, occurrences []*grafeas_proto.Occurrence) error
	BatchCreateGenericResourceVersions(ctx context.Context, occurrences []*grafeas_proto.Occurrence) error
	ListGenericResources(ctx context.Context, request *pb.ListGenericResourcesRequest) (*pb.ListGenericResourcesResponse, error)
	ListGenericResourceVersions(ctx context.Context, request *pb.ListGenericResourceVersionsRequest) (*pb.ListGenericResourceVersionsResponse, error)
	GetGenericResource(ctx context.Context, resourceId string) (*pb.GenericResource, error)
}

type manager struct {
	logger       *zap.Logger
	esClient     esutil.Client
	esConfig     *config.ElasticsearchConfig
	indexManager indexmanager.IndexManager
	filterer     filtering.Filterer
}

func NewManager(logger *zap.Logger, esClient esutil.Client, esConfig *config.ElasticsearchConfig, indexManager indexmanager.IndexManager, filterer filtering.Filterer) Manager {
	return &manager{
		logger:       logger,
		esClient:     esClient,
		esConfig:     esConfig,
		indexManager: indexManager,
		filterer:     filterer,
	}
}

// BatchCreateGenericResources creates generic resources from a list of occurrences. This method is intended to be invoked
// as a side effect of BatchCreateOccurrences.
func (m *manager) BatchCreateGenericResources(ctx context.Context, occurrences []*grafeas_proto.Occurrence) error {
	log := m.logger.Named("BatchCreateGenericResources")

	genericResources := map[string]*pb.GenericResource{}
	var resourceIds []string
	for _, occurrence := range occurrences {
		uriParts, err := parseResourceUri(occurrence.Resource.Uri)
		if err != nil {
			return err
		}

		genericResource := &pb.GenericResource{
			Name:    uriParts.name,
			Type:    uriParts.resourceType,
			Created: occurrence.CreateTime,
		}
		resourceId := uriParts.prefixedName

		if _, ok := genericResources[resourceId]; ok {
			continue
		}
		genericResources[resourceId] = genericResource
		resourceIds = append(resourceIds, resourceId)
	}

	multiGetResponse, err := m.esClient.MultiGet(ctx, &esutil.MultiGetRequest{
		DocumentIds: resourceIds,
		Index:       m.indexManager.AliasName(genericResourcesDocumentKind, ""),
	})
	if err != nil {
		return err
	}

	var bulkItems []*esutil.BulkRequestItem
	for i, resourceId := range resourceIds {
		if multiGetResponse.Docs[i].Found {
			log.Debug("skipping resource creation because it already exists", zap.String("resourceId", resourceId))
			continue
		}

		log.Debug("Adding resource to bulk request", zap.String("resourceId", resourceId))

		bulkItems = append(bulkItems, &esutil.BulkRequestItem{
			Operation:  esutil.BULK_INDEX,
			Message:    genericResources[resourceId],
			DocumentId: resourceId,
			Join: &esutil.EsJoin{
				Field: genericResourceDocumentJoinField,
				Name:  genericResourceRelationName,
			},
		})
	}

	if len(bulkItems) == 0 {
		return nil
	}

	bulkResponse, err := m.esClient.Bulk(ctx, &esutil.BulkRequest{
		Index:   m.indexManager.AliasName(genericResourcesDocumentKind, ""),
		Refresh: m.esConfig.Refresh.String(),
		Items:   bulkItems,
	})
	if err != nil {
		return err
	}

	var bulkItemErrors []error
	for _, item := range bulkResponse.Items {
		if item.Index.Error != nil {
			bulkItemErrors = append(bulkItemErrors, fmt.Errorf("error creating generic resource [%d] %s: %s", item.Index.Status, item.Index.Error.Type, item.Index.Error.Reason))
		}
	}

	if len(bulkItemErrors) > 0 {
		return fmt.Errorf("failed to create all generic resources: %v", bulkItemErrors)
	}

	return nil
}

// BatchCreateGenericResourceVersions creates generic resource versions from a list of occurrences. This method is intended
// to be invoked as a side effect of BatchCreateOccurrences, after BatchCreateGenericResources
func (m *manager) BatchCreateGenericResourceVersions(ctx context.Context, occurrences []*grafeas_proto.Occurrence) error {
	log := m.logger.Named("BatchCreateGenericResourceVersions")

	genericResourceVersions := map[string]*pb.GenericResourceVersion{}
	var versionIds []string

	// build a list of generic resource versions from occurrences. these may or may not already exist
	for _, occurrence := range occurrences {
		newVersions := genericResourceVersionsFromOccurrence(occurrence)

		// check if we already know about these versions
		for versionId, newVersion := range newVersions {
			if existingVersion, ok := genericResourceVersions[versionId]; ok {
				// if we do, update the names and create time, if needed
				if newVersion.Created != nil {
					existingVersion.Created = newVersion.Created
				}

				if newVersion.Names != nil {
					existingVersion.Names = newVersion.Names
				}
			} else {
				genericResourceVersions[versionId] = newVersion
				versionIds = append(versionIds, versionId)
			}
		}
	}

	// check which generic resource versions already exist
	multiGetResponse, err := m.esClient.MultiGet(ctx, &esutil.MultiGetRequest{
		DocumentIds: versionIds,
		Index:       m.indexManager.AliasName(genericResourcesDocumentKind, ""),
	})
	if err != nil {
		return err
	}

	// versions that don't exist need to be created
	// versions that do exist may need to be updated
	var bulkItems []*esutil.BulkRequestItem
	for i, versionId := range versionIds {
		version := genericResourceVersions[versionId]
		uriParts, err := parseResourceUri(versionId)
		if err != nil {
			return err
		}

		bulkItem := &esutil.BulkRequestItem{
			Operation:  esutil.BULK_INDEX,
			Message:    version,
			DocumentId: versionId,
			Join: &esutil.EsJoin{
				Field:  genericResourceDocumentJoinField,
				Name:   genericResourceVersionRelationName,
				Parent: uriParts.prefixedName,
			},
		}

		if multiGetResponse.Docs[i].Found {
			// if the version already exists, we may have to update(index) it.
			// if we have a list of names, update the existing version with the new one
			// otherwise, we have nothing to do here
			if len(version.Names) != 0 {
				log.Debug("updating generic resource version", zap.Any("version", version))
			} else {
				bulkItem = nil
			}
		} else {
			// if the version doesn't exist, we need to create it
			// before creating the version, add a timestamp if it doesn't exist
			if version.Created == nil {
				version.Created = timestamppb.Now()
			}

			log.Debug("creating generic resource version", zap.Any("version", version))
		}

		if bulkItem != nil {
			bulkItems = append(bulkItems, bulkItem)
		}
	}

	// perform bulk create / update
	var bulkItemErrors []error
	if len(bulkItems) != 0 {
		bulkResponse, err := m.esClient.Bulk(ctx, &esutil.BulkRequest{
			Index:   m.indexManager.AliasName(genericResourcesDocumentKind, ""),
			Refresh: m.esConfig.Refresh.String(),
			Items:   bulkItems,
		})
		if err != nil {
			return err
		}

		for _, item := range bulkResponse.Items {
			if item.Index.Error != nil {
				bulkItemErrors = append(bulkItemErrors, fmt.Errorf("error indexing generic resource [%d] %s: %s", item.Index.Status, item.Index.Error.Type, item.Index.Error.Reason))
			}
		}
	}

	if len(bulkItemErrors) > 0 {
		return fmt.Errorf("failed to create/update some generic resource versions: %v", bulkItemErrors)
	}

	return nil
}

func (m *manager) ListGenericResources(ctx context.Context, request *pb.ListGenericResourcesRequest) (*pb.ListGenericResourcesResponse, error) {
	// generic resources and their versions are stored in the same index. we need to specify the join field
	// in this query in order to only select generic resources. we use a "must" here so we can combine this query with
	// the one generated by the filter, if provided
	queries := filtering.Must{
		&filtering.Query{
			Term: &filtering.Term{
				genericResourceDocumentJoinField: genericResourceRelationName,
			},
		},
	}

	if request.Filter != "" {
		filterQuery, err := m.filterer.ParseExpression(request.Filter)
		if err != nil {
			return nil, err
		}

		queries = append(queries, filterQuery)
	}

	searchRequest := &esutil.SearchRequest{
		Index: m.indexManager.AliasName(genericResourcesDocumentKind, ""),
		Search: &esutil.EsSearch{
			Query: &filtering.Query{
				Bool: &filtering.Bool{
					Must: &queries,
				},
			},
		},
	}

	if request.PageSize != 0 {
		searchRequest.Pagination = &esutil.SearchPaginationOptions{
			Size:  int(request.PageSize),
			Token: request.PageToken,
		}
	}

	searchResponse, err := m.esClient.Search(ctx, searchRequest)
	if err != nil {
		return nil, err
	}

	var genericResources []*pb.GenericResource
	for _, hit := range searchResponse.Hits.Hits {
		var genericResource pb.GenericResource
		err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(hit.Source, &genericResource)
		if err != nil {
			return nil, err
		}

		genericResource.Id = hit.ID
		genericResources = append(genericResources, &genericResource)
	}

	return &pb.ListGenericResourcesResponse{
		GenericResources: genericResources,
		NextPageToken:    searchResponse.NextPageToken,
	}, nil
}

// ListGenericResourceVersions handles the main logic for fetching versions associated with a generic resource.
func (m *manager) ListGenericResourceVersions(ctx context.Context, request *pb.ListGenericResourceVersionsRequest) (*pb.ListGenericResourceVersionsResponse, error) {
	// generic resources and their versions are stored in the same index. we need to specify the "has_parent" field
	// in this query in order to only select generic resource versions with the provided generic resource id as the parent.
	// we use a "must" here so we can combine this query with the one generated by the filter, if provided
	queries := filtering.Must{
		&filtering.Query{
			HasParent: &filtering.HasParent{
				ParentType: genericResourceRelationName,
				Query: &filtering.Query{
					Term: &filtering.Term{
						"_id": request.Id,
					},
				},
			},
		},
	}

	if request.Filter != "" {
		filterQuery, err := m.filterer.ParseExpression(request.Filter)
		if err != nil {
			return nil, err
		}

		queries = append(queries, filterQuery)
	}

	searchRequest := &esutil.SearchRequest{
		Index: m.indexManager.AliasName(genericResourcesDocumentKind, ""),
		Search: &esutil.EsSearch{
			Query: &filtering.Query{
				Bool: &filtering.Bool{
					Must: &queries,
				},
			},
			Sort: map[string]esutil.EsSortOrder{
				"created": esutil.EsSortOrderDescending,
			},
		},
	}

	if request.PageSize != 0 {
		searchRequest.Pagination = &esutil.SearchPaginationOptions{
			Size:  int(request.PageSize),
			Token: request.PageToken,
		}
	}

	searchResponse, err := m.esClient.Search(ctx, searchRequest)
	if err != nil {
		return nil, err
	}

	var genericResourceVersions []*pb.GenericResourceVersion
	for _, hit := range searchResponse.Hits.Hits {
		var genericResourceVersion pb.GenericResourceVersion
		err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(hit.Source, &genericResourceVersion)
		if err != nil {
			return nil, err
		}

		genericResourceVersions = append(genericResourceVersions, &genericResourceVersion)
	}

	return &pb.ListGenericResourceVersionsResponse{
		Versions:      genericResourceVersions,
		NextPageToken: searchResponse.NextPageToken,
	}, nil
}

func (m *manager) GetGenericResource(ctx context.Context, resourceId string) (*pb.GenericResource, error) {
	log := m.logger.Named("GetGenericResource").With(zap.String("resource", resourceId))

	response, err := m.esClient.Get(ctx, &esutil.GetRequest{
		Index:      m.indexManager.AliasName(genericResourcesDocumentKind, ""),
		DocumentId: resourceId,
	})
	if err != nil {
		return nil, err
	}

	if !response.Found {
		log.Debug("generic resource not found")

		return nil, nil
	}

	var resource pb.GenericResource
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(response.Source, &resource)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

// genericResourceVersionsFromOccurrence will create a map of generic versions, keyed by their IDs, from an occurrence.
// if the occurrence is a build occurrence, generic resource versions will also be created for each built artifact.
func genericResourceVersionsFromOccurrence(o *grafeas_proto.Occurrence) map[string]*pb.GenericResourceVersion {
	result := make(map[string]*pb.GenericResourceVersion)

	result[o.Resource.Uri] = &pb.GenericResourceVersion{
		Version: o.Resource.Uri,
	}

	if o.Kind == common_go_proto.NoteKind_BUILD {
		details := o.Details.(*grafeas_proto.Occurrence_Build)

		for _, artifact := range details.Build.Provenance.BuiltArtifacts {
			result[artifact.Id] = &pb.GenericResourceVersion{
				Version: artifact.Id,
				Names:   artifact.Names,
				Created: o.CreateTime,
			}
		}
	}

	return result
}