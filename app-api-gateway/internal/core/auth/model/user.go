package model

type User struct {
	UserID      string   `json:"user_id"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	Address     []string `json:"address"`
}

type UserMeta struct {
	UserMetaID   string `json:"user_meta_id"`
	UserID       string `json:"user_id"`
	FacebookUrl  string `json:"facebook_url"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`
	LinkedinUrl  string `json:"linkedin_url"`
	WebsiteUrl   string `json:"website_url"`
	DealershipID string `json:"dealership_id"`
}
