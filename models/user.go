package models

type UserInfo struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	FollowerCount   int    `json:"follower_count"`
}
