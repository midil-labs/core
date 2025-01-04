package request

import "github.com/midil-labs/core/shared/dtos/common"

type QueryOption = common.Option[QueryParams]
type BodyOption common.Option[Resource]


func WithFilter(fields map[string][]string) QueryOption {
	return func(q *QueryParams) {
		q.Filter.Fields = fields
	}
}

func WithSort(fields ...string) QueryOption {
	return func(q *QueryParams) {
		if q.Sort.Fields == nil {
			q.Sort.Fields = make([]string, 0)
		}
		q.Sort.Fields = append(q.Sort.Fields, fields...)
	}
}

func WithPagination(pageSize, pageNumber int) QueryOption {
	return func(q *QueryParams) {
		q.Page.PageSize = pageSize
		q.Page.PageNumber = pageNumber
	}
}

func WithToOneRelationship(name, relType, relID string, opts ...common.MetaOption) BodyOption {
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
			},
		}

		common.ApplyOptions[common.NonStandardMeta](&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}

func WithToManyRelationship(name string, resources []common.ResourceIdentifier, opts ...common.MetaOption) BodyOption {
	return func(r *Resource) {
		if r.Relationships == nil {
			r.Relationships = make(map[string]common.Relationship)
		}
		relationship := common.Relationship{
			Data: common.RelationshipData{
				Resources: resources,
			},
		}
		common.ApplyOptions[common.NonStandardMeta](&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}
