package model

type User struct {
	Login     string
	Name      *string
	CreatedAt string
	UpdatedAt string
	Id        int64
	Points    int64
}

type LoginUser struct {
	Login    string
	PassHash string
	Id       int64
}
