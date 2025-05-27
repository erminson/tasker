package model

type User struct {
	Login     string
	Name      *string
	CreatedAt string
	UpdatedAt string
	Id        int64
	Points    int64
}
