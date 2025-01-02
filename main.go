package main


import (
	"fmt"
	"github.com/midil-labs/core/shared/dtos/response" // Ensure this path is correct and the package exists
	"encoding/json"
)




type UserAttributes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UserAttributes) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

type PostAttributes struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p PostAttributes) Validate() error {
	if p.Title == "" {
		return fmt.Errorf("title is required")
	}
	if p.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}

type CommentAttributes struct {
	Content string `json:"content"`
}

func (c CommentAttributes) Validate() error {
	if c.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}



func main() {
	userResource := response.Resource[UserAttributes]{
		ResourceIdentifier: response.ResourceIdentifier{
			ID:   "123",
			Type: "users",
		},
		Attributes: UserAttributes{
			Name:  "John Doe",
			Email: "john@example.com",
		},
		Links: &response.Links{
			Self: "/users/123",
		},
		Meta: map[string]interface{}{
			"created_at": "2024-01-01T00:00:00Z",
		},
	}

	// Example 2: Creating a resource with relationships
	postResource := response.Resource[PostAttributes]{
		ResourceIdentifier: response.ResourceIdentifier{
			ID:   "456",
			Type: "posts",
		},
		Attributes: PostAttributes{
			Title:   "My First Post",
			Content: "Hello, world!",
		},
		Relationships: map[string]response.Relationship{
			"author": {
				Data: response.RelationshipData{
					Resource: &response.ResourceIdentifier{
						ID:   "123",
						Type: "users",
					},
				},
				Links: &response.Links{
					Self: "/posts/456/relationships/author",
					Related: &response.RelatedLink{
						Href: "/users/123",
					},
				},
			},
			"comments": {
				Data: response.RelationshipData{
					Resources: []response.ResourceIdentifier{
						{ID: "789", Type: "comments"},
						{ID: "790", Type: "comments"},
					},
				},
				Links: &response.Links{
					Self: "/posts/456/relationships/comments",
					Related: &response.RelatedLink{
						Href: "/posts/456/comments",
					},
				},
			},
		},
	}

	// Example 3: Creating a single resource response
	singleUserResponse := response.SingleResourceResponse[UserAttributes]{
		Data: &userResource,
		Meta: map[string]interface{}{
			"pagination": response.Pagination{
				CurrentPage: 1,
				TotalPages:  1,
				TotalCount:  1,
			},
		},
	}

	// Example 4: Creating a multiple resources response
	resourcesPostsResponse := response.MultipleResourcesResponse[PostAttributes]{
		Data: []response.Resource[PostAttributes]{
			postResource,
			{
				ResourceIdentifier: response.ResourceIdentifier{
					ID:   "457",
					Type: "posts",
				},
				Attributes: PostAttributes{
					Title:   "Another Post",
					Content: "Another hello!",
				},
			},
		},
		Links: &response.PaginationLinks{
			Self:  "/posts",
			First: "/posts?page=1",
			Last:  "/posts?page=2",
			Next:  "/posts?page=2",
		},
	}

	// Serialize to JSON
	includedResourcesJSON, _ := json.MarshalIndent(includedResourcesResponse, "", "  ")

	fmt.Println("\nIncluded Resources Response:")
	fmt.Println(string(includedResourcesJSON))

	// Example 6: Creating a resource with links

	fmt.Println("\nUser Resource:")
	fmt.Println(userResource)
	// Serialize to JSON
	userJSON, _ := json.MarshalIndent(userResource, "", "  ")
	postJSON, _ := json.MarshalIndent(postResource, "", "  ")
	singleResponseJSON, _ := json.MarshalIndent(singleUserResponse, "", "  ")
	resourcesResponseJSON, _ := json.MarshalIndent(resourcesPostsResponse, "", "  ")

	fmt.Println("User Resource:")
	fmt.Println(string(userJSON))
	fmt.Println("\nPost Resource:")
	fmt.Println(string(postJSON))
	fmt.Println("\nSingle User Response:")
	fmt.Println(string(singleResponseJSON))
	fmt.Println("\nresources Posts Response:")
	fmt.Println(string(resourcesResponseJSON))

}