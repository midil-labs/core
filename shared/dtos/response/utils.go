package response

import (
	"fmt"
	"github.com/midil-labs/core/shared/dtos"
)

func BuildPaginationLinks(baseURL string, currentPage, totalPages int64) *PaginationLinks {
	links := &PaginationLinks{
		Self: fmt.Sprintf("%s?page=%d", baseURL, currentPage),
	}

	if currentPage > 1 {
		links.First = fmt.Sprintf("%s?page=1", baseURL)
	}

	if totalPages > 0 {
		links.Last = fmt.Sprintf("%s?page=%d", baseURL, totalPages)
	}

	if currentPage > 1 {
		links.Prev = fmt.Sprintf("%s?page=%d", baseURL, currentPage-1)
	}

	if currentPage < totalPages {
		links.Next = fmt.Sprintf("%s?page=%d", baseURL, currentPage+1)
	}

	return links
}


func BuildPagination(currentPage, totalCount, pageSize int64) Pagination {
	totalPages := (totalCount + pageSize - 1) / pageSize

	return Pagination{
		CurrentPage: currentPage,
		PrevPage:    max(1, currentPage-1),
		NextPage:    min(totalPages, currentPage+1),
		TotalPages:  totalPages,
		TotalCount:  totalCount,
	}
}


func BuildSingleResourceResponse[T dtos.DTOInterface](
	id, resourceType string, 
	attributes T, 
	relationships map[string]Relationship,
	baseURL string,
) *SingleResourceResponse[T] {

	if id == "" || resourceType == "" {
		return nil
	}

	resource := &Resource[T]{
		ResourceIdentifier: ResourceIdentifier{
			ID:   id,
			Type: resourceType,
		},
		Attributes: attributes,
		Links: &Links{
			Self: fmt.Sprintf("%s/%s/%s", baseURL, resourceType, id),
		},
	}

	if len(relationships) > 0 {
		resource.Relationships = relationships
	}

	return &SingleResourceResponse[T]{
		Data: resource,
	}
}


func BuildMultipleResourcesResponse[T dtos.DTOInterface](
	resources []T, 
	resourceType string, 
	baseURL string, 
	currentPage, totalCount, pageSize int64,
) *MultipleResourcesResponse[T] {

	var data []Resource[T]
	for i, item := range resources {
		data = append(data, Resource[T]{
			ResourceIdentifier: ResourceIdentifier{
				ID:   fmt.Sprintf("%d", i+1),
				Type: resourceType,
			},
			Attributes: item,
			Links: &Links{
				Self: fmt.Sprintf("%s/%s/%d", baseURL, resourceType, i+1),
			},
		})
	}

	response := &MultipleResourcesResponse[T]{
		Data: data,
		Links: BuildPaginationLinks(baseURL, currentPage, (totalCount+pageSize-1)/pageSize),
		Meta: &Meta{
			Pagination: BuildPagination(currentPage, totalCount, pageSize),
		},
	}

	return response
}


func BuildErrorResponse(
	errorCode string, 
	title string, 
	detail string, 
	source *ErrorSource,
) *ErrorResponse {
	return &ErrorResponse{
		Errors: []ErrorObject{
			{
				Status: errorCode,
				Title:  title,
				Detail: detail,
				Source: source,
			},
		},
	}
}


func BuildRelationship(
	relatedID string, 
	relatedType string, 
	baseURL string,
) Relationship {
	return Relationship{
		Data: RelationshipData{
			Resource: &ResourceIdentifier{
				ID:   relatedID,
				Type: relatedType,
			},
		},
		Links: &Links{
			Self: fmt.Sprintf("%s/%s/%s", baseURL, relatedType, relatedID),
		},
	}
}