package request

import (
	"io"

	"github.com/midil-labs/core/shared/jsonapi/common"
	"github.com/midil-labs/core/shared/jsonapi/jsonapierror"
	"github.com/midil-labs/core/shared/utils/goutils"
)

type QueryOption = goutils.Option[Query]

type ErrableOption[T any] func(*T) jsonapierror.ErrorObjects

type RequestOption[T BodyType] ErrableOption[JSONAPIRequest[T]]

type BodyOption = goutils.Option[Body]

type BodyTypeOption[T BodyType] goutils.Option[T]

// WithFilter adds a filter to the query.
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
		if q.Fields == nil {
			q.Fields = make(Fields)
		}
		q.Fields[key] = value
	}
}

// WithInclude adds relationships to include in the response.
func WithInclude(fields ...string) QueryOption {
	return func(q *Query) {
		q.Include = append(q.Include, fields...)
	}
}

// WithPagination adds pagination parameters to the query.
func WithPagination(pageSize, pageNumber int) QueryOption {
	return func(q *Query) {
		q.Page.Size = pageSize
		q.Page.Number = pageNumber
	}
}

// WithLID sets the LID field on the Resource.
func WithLID(lid string) BodyOption {
	return func(r *Body) {
		id := common.ID(lid)
		r.LID = &id
	}
}

// WithToOneRelationship adds a to-one relationship to the body.
func WithToOneRelationship(name, relType, relID string, opts ...common.MetaOption) BodyOption {
	return func(r *Body) {
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

// WithToManyRelationship adds a to-many relationship to the body.
func WithToManyRelationship(name string, resources []*common.ResourceIdentifier, opts ...common.MetaOption) BodyOption {
	return func(r *Body) {
		if r.Relationships == nil {
			r.Relationships = make(map[string]common.Relationship)
		}
		relationship := common.Relationship{
			Data: common.RelationshipData{
				Resources: resources,
				Resource:  nil,
			},
		}
		goutils.ApplyOptions(&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}

// WithToManyRelationshipFromMap adds a to-many relationship to the body from a map.
func WithToManyRelationshipFromMap[T BodyType](name string, resources []map[string]string, opts ...common.MetaOption) BodyOption {
	return func(r *Body) {
		if r.Relationships == nil {
			r.Relationships = make(map[string]common.Relationship)
		}

		var resourceIdentifiers []*common.ResourceIdentifier
		for _, res := range resources {
			id := res["id"]
			resourceIdentifiers = append(resourceIdentifiers, &common.ResourceIdentifier{
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

// WithQuery parses query parameters and adds them to the request.
func WithQuery[T BodyType](query map[string][]string) RequestOption[T] {
	return func(r *JSONAPIRequest[T]) jsonapierror.ErrorObjects {
		if r.Query == nil {
			r.Query = &Query{}
		}
		parsedQuery, errs := ParseQueryParams(query)
		if len(errs) > 0 {
			return errs
		}
		r.Query = &parsedQuery
		return nil
	}
}

// WithHeaders adds headers to the request.
func WithHeaders[T BodyType](header map[string][]string) RequestOption[T] {
	return func(r *JSONAPIRequest[T]) jsonapierror.ErrorObjects {
		if r.Header == nil {
			r.Header = make(map[string][]string)
		}
		for k, v := range header {
			r.Header[k] = v
		}
		return nil
	}
}

// WithBody parses the request body and adds it to the request.
func WithBody[T BodyType](body io.Reader) RequestOption[T] {
	return func(r *JSONAPIRequest[T]) jsonapierror.ErrorObjects {
		parsedBody, errs := ParseBody[T](body)
		if len(errs) > 0 {
			return errs
		}
		r.Body = parsedBody
		return nil
	}
}
