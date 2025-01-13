package common


type ID = string

type Type  = string

// ResourceIdentifier is a JSON:API resource identifier object.
type ResourceIdentifier struct {
	ID   *ID  	`json:"id"`
	Type Type  	`json:"type" required:"true"`
}

// NonStandardMeta is a map of non-standard meta-information.
type NonStandardMeta map[string]interface{}

// RelationshipData holds either one or many ResourceIdentifiers.
type RelationshipData struct {
	Resource  *ResourceIdentifier   `json:"resource,omitempty"`
	Resources []ResourceIdentifier  `json:"resources,omitempty"`
}

type RelatedLink struct {
	Href        string            `json:"href,omitempty"`
	Title       string            `json:"title,omitempty"`
	DescribedBy string            `json:"describedby,omitempty"`
	Meta        NonStandardMeta   `json:"meta,omitempty"`
}


type Links struct {
	Self    string      		`json:"self,omitempty"`
	Related *RelatedLink 		`json:"related,omitempty"`
}
func (l *Links) ApplyOptions(opts ...LinksOption) {
	ApplyOptions(l, opts...)
}

// Relationship is a JSON:API relationship object.
type Relationship struct {
	Data  RelationshipData 			`json:"data,omitempty"`
	Meta  NonStandardMeta           `json:"meta,omitempty"`
	Links *Links                    `json:"links,omitempty"`
}

func (r *Relationship) ApplyOptions(opts ...RelationshipOption) {
	ApplyOptions(r, opts...)
}