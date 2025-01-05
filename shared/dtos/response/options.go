package response

import "github.com/midil-labs/core/shared/dtos/common"

type ResourceOption = common.Option[Resource]
type ResponseOption[T DataType] common.Option[JSONAPIResponse[T]]


func WithToOneRelationship(name, relType, relID string, links *common.Links, meta common.NonStandardMeta) ResourceOption {
    return func(r *Resource) {
        if r.Relationships == nil {
            r.Relationships = make(map[string]Relationship)
        }
        r.Relationships[name] = Relationship{
            Data: RelationshipData{
                Resource: &common.ResourceIdentifier{
                    Type: relType,
                    ID:   &relID,
                },
            },
            Links: links,
            Meta:  meta,
        }
    }
}

func WithToManyRelationship(name string, resources []common.ResourceIdentifier, links *common.Links, meta common.NonStandardMeta) ResourceOption {
    return func(r *Resource) {
        if r.Relationships == nil {
            r.Relationships = make(map[string]Relationship)
        }
        r.Relationships[name] = Relationship{
            Data: RelationshipData{
                Resources: resources,
            },
            Links: links,
            Meta:  meta,
        }
    }
}

func WithRelationships(relationships map[string]Relationship) ResourceOption {
	return func(r *Resource) {
		r.Relationships = relationships
	}
}


func WithLinks(self string, related *common.RelatedLink) ResourceOption {
	return func(r *Resource) {
		if r.Links == nil {
			r.Links = &common.Links{}
		}
		r.Links.Self = self
		r.Links.Related = related
	}
}

func WithIncluded[T DataType](included ListResource) ResponseOption[T] {
	return func(rr *JSONAPIResponse[T]) {
		rr.Included = included
	}
}


func WithMeta(meta map[string]any) ResourceOption {
	return func(r *Resource) {
		r.Meta = meta
	}
}

func WithPagination(currentPage, prevPage, nextPage, totalPages, totalCount int64) ResponseOption[ListResource] {
    return func(rr *JSONAPIResponse[ListResource]) {
        if rr.Meta == nil {
            rr.Meta = common.NonStandardMeta{}
        }
        rr.Meta["Pagination"] = Pagination{
            CurrentPage: currentPage,
			PrevPage: prevPage,
			NextPage: nextPage,
            TotalPages: totalPages,
            TotalCount: totalCount,
        }
    }
}


func WithPaginationLinks(self, first, last, prev, next string) ResponseOption[ListResource] {
	return func(r *JSONAPIResponse[ListResource]) {
		if r.Links == nil {
			r.Links = &PaginationLinks{}
		}
		r.Links.Self = self
		r.Links.First = first
		r.Links.Last = last
		r.Links.Prev = prev
		r.Links.Next = next
	}
}


func WithIncludedResources[T DataType](resources ListResource) ResponseOption[T] {
    return func(r *JSONAPIResponse[T]) {
        r.Included = append(ListResource(nil), resources...)
    }
}


func WithDataResource(id common.ID, resourceType common.Type, attributes map[string]interface{}, opts ...ResourceOption) ResponseOption[*Resource] {
    return func(r *JSONAPIResponse[*Resource]) {
        resource := &Resource{
            ResourceIdentifier: common.ResourceIdentifier{
                ID:   &id,
                Type: resourceType,
            },
            Attributes:   attributes,
			Relationships: make(map[string]Relationship),
			Meta:         make(map[string]any),
        }
        for _, opt := range opts {
            opt(resource)
        }
        r.Data = resource
    }
}


func WithManyDataResources(resources ...Resource) ResponseOption[ListResource] {
    return func(r *JSONAPIResponse[ListResource]) {
        r.Data = append(r.Data, resources...)
    }
}
