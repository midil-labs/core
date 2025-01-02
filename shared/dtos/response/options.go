package response

import "github.com/midil-labs/core/shared/dtos"

type ResourceOption[T dtos.DTOInterface] func(*Resource[T])
type ResponseOption[T dtos.DTOInterface] func(*ResourceResponse[T])


func WithToOneRelationship[T dtos.DTOInterface](name, relType, relID string, links *Links, meta NonStandardMeta) ResourceOption[T] {
    return func(r *Resource[T]) {
        if r.Relationships == nil {
            r.Relationships = make(map[string]Relationship)
        }
        r.Relationships[name] = Relationship{
            Data: RelationshipData{
                Resource: &ResourceIdentifier{
                    Type: relType,
                    ID:   relID,
                },
            },
            Links: links,
            Meta:  meta,
        }
    }
}

func WithToManyRelationship[T dtos.DTOInterface](name string, resources []ResourceIdentifier, links *Links, meta NonStandardMeta) ResourceOption[T] {
    return func(r *Resource[T]) {
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


func WithRelationships[T dtos.DTOInterface](relationships map[string]Relationship) ResourceOption[T] {
	return func(r *Resource[T]) {
		r.Relationships = relationships
	}
}


func WithLinks[T dtos.DTOInterface](self string, related *RelatedLink) ResourceOption[T] {
	return func(r *Resource[T]) {
		if r.Links == nil {
			r.Links = &Links{}
		}
		r.Links.Self = self
		r.Links.Related = related
	}
}

func WithIncluded[T dtos.DTOInterface](included []Resource[dtos.DTOInterface]) ResponseOption[T] {
	return func(rr *ResourceResponse[T]) {
		rr.Included = included
	}
}


func WithMeta[T dtos.DTOInterface](meta map[string]any) ResourceOption[T] {
	return func(r *Resource[T]) {
		r.Meta = meta
	}
}

func WithPagination[T dtos.DTOInterface](currentPage, prevPage, nextPage, totalPages, totalCount int64) ResponseOption[T] {
    return func(rr *ResourceResponse[T]) {
        if rr.Meta == nil {
            rr.Meta = NonStandardMeta{}
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


func WithPaginationLinks[T dtos.DTOInterface](self, first, last, prev, next string) ResponseOption[T] {
	return func(r *ResourceResponse[T]) {
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




func WithResource[T dtos.DTOInterface](id string, resourceType string, attributes T, opts ...ResourceOption[T]) ResponseOption[T] {
    return func(r *ResourceResponse[T]) {
        resource := &Resource[T]{
            ResourceIdentifier: ResourceIdentifier{
                ID:   id,
                Type: resourceType,
            },
            Attributes:   attributes,
			Relationships: make(map[string]Relationship),
			Meta:         make(map[string]any),
        }
        for _, opt := range opts {
            opt(resource)
        }
        r.resource = resource
		r.resources = nil
    }
}


func WithMultipleResources[T dtos.DTOInterface](res ...Resource[T]) ResponseOption[T] {
    return func(r *ResourceResponse[T]) {
        r.resources = append(r.resources, res...)
        r.resource  = nil
    }
}
