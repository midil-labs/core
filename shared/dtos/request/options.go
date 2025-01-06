package request

import "github.com/midil-labs/core/shared/dtos/common"

type QueryOption = common.Option[QueryParams]
type BodyOption common.Option[Resource]
// type HeaderOption common.Option[Header]


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

		common.ApplyOptions(&relationship.Meta, opts...)

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
		common.ApplyOptions(&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}

func WithToManyRelationshipFromMap(name string, resources []map[string]string, opts ...common.MetaOption) BodyOption {
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
		common.ApplyOptions(&relationship.Meta, opts...)

		r.Relationships[name] = relationship
	}
}