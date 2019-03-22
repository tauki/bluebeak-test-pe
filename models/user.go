package models

type UserInfo struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	FollowerCount   int    `json:"follower_count"`
}

type Reviews struct {
	Points              string `json:"points"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	TasterName          string `json:"taster_name"`
	TasterTwitterHandle string `json:"taster_twitter_handle"`
	Price               *int   `json:"price"` // pointer to support nil
	Designation         string `json:"designation"`
	Variety             string `json:"variety"`
	Region1             string `json:"region_1"`
	Region2             string `json:"region_2"`
	Province            string `json:"province"`
	Country             string `json:"country"`
	Winery              string `json:"winery"`
}
