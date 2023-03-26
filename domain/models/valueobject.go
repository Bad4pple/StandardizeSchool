package models

type Customer struct {
	OrganizationName string   `json:"organization_name" bson:"organization_name"`
	Phone            string   `json:"phone" bson:"phone"`
	Email            string   `json:"email" bson:"email"`
	SocialLinks      []string `json:"social_links" bson:"social_links"`
	Address          string   `json:"address" bson:"address"`
}
