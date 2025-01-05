package common

type Option[T any] func(*T)

// MetaOption is a function that applies a configuration to a NonStandardMeta.
type MetaOption = Option[NonStandardMeta]

// Option is a function that applies a configuration to a Relationship.
type RelationshipOption = Option[Relationship]

// RelatedLinkOption is a function that applies a configuration to a RelatedLink.
type RelatedLinkOption = Option[RelatedLink]

type LinksOption = Option[Links]


func WithRelatedMetadata(meta NonStandardMeta) RelationshipOption {
	return func(r *Relationship) {
		r.Meta = meta
	}
}

func WithRelatedLinks(self string, related *RelatedLink) RelationshipOption {
	return func(r *Relationship) {
		if r.Links == nil {
			r.Links = &Links{}
		}
		r.Links.Self = self
		r.Links.Related = related
	}
}

func WithToOneRelationship(name, relType string, relID *string, opts ...RelationshipOption) RelationshipOption {
	return func(r *Relationship) {
		if r.Data.Resource == nil {
			r.Data.Resource = &ResourceIdentifier{}
		}
		r.Data.Resource.Type = relType
		r.Data.Resource.ID = relID
		r.ApplyOptions(opts...)
		r.Data.Resources = nil
	}
}

func WithToManyRelationship(name string, resources []ResourceIdentifier, opts ...RelationshipOption) RelationshipOption {
	return func(r *Relationship) {
		r.Data.Resources = append(r.Data.Resources, resources...)
		r.ApplyOptions(opts...)
		r.Data.Resource = nil
	}
}

func WithRelatedLink(href, title, describedBy string, meta NonStandardMeta) LinksOption {
	return func(rl *Links) {
		rl.Related.Href = href
		rl.Related.Title = title
		rl.Related.DescribedBy = describedBy
		rl.Related.Meta = meta
	}
}