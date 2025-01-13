package response

import "github.com/midil-labs/core/shared/jsonapi/common"


type ResponseOption[T DataType] common.Option[JSONAPIResponse[T]]
type ResourceOption common.Option[Resource]

type ListResourceOption common.Option[JSONAPIResponse[ListResource]] // ListResource response option


// WithMeta merges the provided `meta` into different places based on `scope`.
// Valid scopes: "top", "resource", "relationship", "link".
func WithMeta[T DataType](meta map[string]any, scope Scope) ResponseOption[T] {
	return func(r *JSONAPIResponse[T]) {
		if !scope.IsValid() {
			panic(ErrInvalidScope)
		}

		switch scope {
		case ScopeTop:
			if r.Meta == nil {
				r.Meta = make(map[string]any)
			}
			for k, v := range meta {
				r.Meta[k] = v
			}

		case ScopeResource:
			switch dataAny := any(r.Data).(type) {
			case *Resource:
				mergeResourceMeta(dataAny, meta)
			case ListResource:
				for i := range dataAny {
					mergeResourceMeta(dataAny[i], meta)
				}
				r.Data = any(dataAny).(T)
			}

		case ScopeRelationship:
			switch dataAny := any(r.Data).(type) {
			case *Resource:
				mergeRelationshipMeta(dataAny, meta)
			case ListResource:
				for i := range dataAny {
					mergeRelationshipMeta(dataAny[i], meta)
				}
				r.Data = any(dataAny).(T)
			}

		case ScopeLink:
			switch dataAny := any(r.Data).(type) {
			case *Resource:
				mergeLinksMeta(dataAny, meta)
			case ListResource:
				for i := range dataAny {
					mergeLinksMeta(dataAny[i], meta)
				}
				r.Data = any(dataAny).(T)
			}
		}
	}
}


func WithToOneRelationship(
	name, relType, relID string,
	links *common.Links,
	meta common.NonStandardMeta,
) ResponseOption[*Resource] {
	return func(resp *JSONAPIResponse[*Resource]) {
		r := resp.Data
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

func WithRelationshipMeta(name string, meta common.NonStandardMeta) ResponseOption[*Resource] {
    return func(resp *JSONAPIResponse[*Resource]) {
        if resp.Data.Relationships == nil {
            resp.Data.Relationships = make(map[string]Relationship)
        }
        resp.Data.Relationships[name] = Relationship{
            Meta: meta,
        }
    }
}


// WithToManyRelationship sets a to-many relationship on a single Resource response.
func WithToManyRelationship(
	name string,
	resources []common.ResourceIdentifier,
	links *common.Links,
	meta common.NonStandardMeta,
) ResponseOption[*Resource] {
	return func(resp *JSONAPIResponse[*Resource]) {
		r := resp.Data
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

// WithLinks sets the `Links` field in a single Resource response.
func WithLinks(self string, related *common.RelatedLink) ResponseOption[*Resource] {
	return func(resp *JSONAPIResponse[*Resource]) {
		if resp.Data.Links == nil {
			resp.Data.Links = &common.Links{}
		}
		resp.Data.Links.Self = self
		resp.Data.Links.Related = related
	}
}


// WithIncluded overwrites the `Included` slice with a new set of included resources.
func WithNewIncludedResources[T DataType](included ListResource) ResponseOption[T] {
    return func(rr *JSONAPIResponse[T]) {
        rr.Included = included
    }
}

// WithIncludedResources appends included resources to the existing slice.
func WithAppendIncludedResources[T DataType](resources ListResource) ResponseOption[T] {
    return func(rr *JSONAPIResponse[T]) {
        rr.Included = append(rr.Included, resources...)
    }
}

// WithPagination sets a Pagination object inside `Meta["Pagination"]` for a multi-resource response.
func WithPagination(currentPage, prevPage, nextPage, totalPages, totalCount int64) ResponseOption[ListResource] {
	return func(rr *JSONAPIResponse[ListResource]) {
		if rr.Meta == nil {
			rr.Meta = common.NonStandardMeta{}
		}
		rr.Meta["Pagination"] = Pagination{
			CurrentPage: currentPage,
			PrevPage:    prevPage,
			NextPage:    nextPage,
			TotalPages:  totalPages,
			TotalCount:  totalCount,
		}
	}
}

// WithPaginationLinks sets pagination-related links in the top-level `Links`.
func WithPaginationLinks(self, first, last, prev, next string) ResponseOption[ListResource] {
	return func(rr *JSONAPIResponse[ListResource]) {
		if rr.Links == nil {
			rr.Links = &PaginationLinks{}
		}
		rr.Links.Self = self
		rr.Links.First = first
		rr.Links.Last = last
		rr.Links.Prev = prev
		rr.Links.Next = next
	}
}
