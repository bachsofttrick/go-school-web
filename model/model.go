package model

// School website
type SchoolWeb struct {
	ID          int    `json:"id" csv:"id"`
	Site        string `json:"site" csv:"site"`
	Name        string `json:"name" csv:"name"`
	CreatedDate int64  `json:"created_date" csv:"created_date"`
	Comment     string `json:"comment" csv:"comment"`
}
