package common

import ("github.com/midil-labs/core/shared/utils/goutils"
		 "encoding/json"
		"fmt")


// MetaOption is a function that applies a configuration to a NonStandardMeta.
type MetaOption = goutils.Option[NonStandardMeta]

// Option is a function that applies a configuration to a Relationship.
type RelationshipOption = goutils.Option[Relationship]

// RelatedLinkOption is a function that applies a configuration to a RelatedLink.
type RelatedLinkOption = goutils.Option[RelatedLink]

type LinksOption = goutils.Option[Links]

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

func WithToManyRelationship(name string, resources []*ResourceIdentifier, opts ...RelationshipOption) RelationshipOption {
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

func (c *RelationshipData) UnmarshalJSON(b []byte) error {
	var resource ResourceIdentifier
	if err := json.Unmarshal(b, &resource); err == nil && resource.ID != nil && *resource.ID != "" {
		c.Resource = &resource
		return nil
	}

	var resources []*ResourceIdentifier
	if err := json.Unmarshal(b, &resources); err == nil {
		c.Resources = resources
		return nil
	}

	return fmt.Errorf("data field is neither a resource object nor a valid array of objects")
}

func (c RelationshipData) MarshalJSON() ([]byte, error) {
	if c.Resource != nil {
		return json.Marshal(c.Resource)
	}
	return json.Marshal(c.Resources)
}