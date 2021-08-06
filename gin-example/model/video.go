package model

type Video struct {
	Id          int    `uri:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}
