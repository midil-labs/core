package request

import (
	"io"
	"net/http"
	"net/url"

	"github.com/midil-labs/core/shared/jsonapi/common"
	"github.com/midil-labs/core/shared/utils/goutils"
)

type QueryOption = goutils.Option[Query]

type RequestOption[T RequestType] goutils.Option[JSONAPIRequest[T]]

type ResourceOption  = goutils.Option[Resource]

type BodyOption[T RequestType] goutils.Option[T]


func WithFilter(key string, values ...string) QueryOption {
    return func(q *Query) {
        if q.Filter == nil {
            q.Filter = make(Filter)
        }
        q.Filter[key] = append(q.Filter[key], values...)
    }
}



// WithSort adds sorting criteria to the query.
func WithSort(fields []string) QueryOption {
	return func(q *Query) {
		q.Sort = append(q.Sort, fields...)
	}
}


// WithFields adds sparse fieldsets to the query.
func WithFields(key string, value ...string) QueryOption {
	return func(q *Query) {
		q.Fields[key] = value
	}
}


// WithInclude adds relationships to include in the response.
func WithInclude(fields ...string) QueryOption {
    return func(q *Query) {
        q.Include = append(q.Include, fields...)
    }
}


func WithPagination(pageSize, pageNumber int) QueryOption {
	return func(q *Query) {
		q.Page.Size = pageSize
		q.Page.Number = pageNumber
	}
}

// WithLID sets the LID field on the Resource.
// It takes a string and assigns it to the LID field after converting it to common.ID.
func WithLID(lid string) ResourceOption {
    return func(r *Resource) {
        id := common.ID(lid)
        r.LID = &id
    }
}

func WithToOneRelationship(name, relType, relID string, opts ...common.MetaOption) ResourceOption {
	return func(r *Resource) {
		if r.Relationships == nil {
			r.Relationships = make(map[string]common.Relationship)
		}

		relationship := common.Relationship{
			Data: common.RelationshipData{
				Resource: &common.ResourceIdentifier{
					Type: relType,
					ID:   &relID,
				},
				Resources: nil,
			},
		}

		goutils.ApplyOptions(&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}


func WithToManyRelationship(name string, resources []common.ResourceIdentifier, opts ...common.MetaOption) ResourceOption {
	return func(r *Resource) {
		if r.Relationships == nil {
			r.Relationships = make(map[string]common.Relationship)
		}
		relationship := common.Relationship{
			Data: common.RelationshipData{
				Resources: resources,
				Resource: nil,
			},
		}
		goutils.ApplyOptions(&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}


func WithToManyRelationshipFromMap[T RequestType](name string, resources []map[string]string, opts ...common.MetaOption) ResourceOption {
	return func(r *Resource) {
		if r.Relationships == nil {
			r.Relationships = make(map[string]common.Relationship)
		}

		var resourceIdentifiers []common.ResourceIdentifier
		for _, res := range resources {
			id := res["id"]
			resourceIdentifiers = append(resourceIdentifiers, common.ResourceIdentifier{
				Type: res["type"],
				ID:   &id,
			})
		}

		relationship := common.Relationship{
			Data: common.RelationshipData{
				Resources: resourceIdentifiers,
			},
		}
		goutils.ApplyOptions(&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}


func WithQuery[T RequestType](query map[string][]string) RequestOption[T] {
	return func(r *JSONAPIRequest[T]) {
		if r.Query == nil {
			r.Query = &Query{}
		}
		query := ParseQueryParams(query)
		r.Query = &query
	}
}

func WithHeaders[T RequestType](header map[string][]string) RequestOption[T] {
	return func(r *JSONAPIRequest[T]) {
		if r.Header == nil {
			r.Header = make(map[string][]string)
		}
		for k, v := range header {
			r.Header[k] = v
		}
	}
}

func WithBody[T RequestType](body io.Reader) RequestOption[T] {
	return func(r *JSONAPIRequest[T]) {
		parsedBody, err := ParseBody[T](body)
		if err != nil {
			panic(err)
		}
		r.Body = parsedBody
	}
}