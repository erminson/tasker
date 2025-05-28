package model

type User struct {
	Name   string
	Points int64
}

type UserInfo struct {
	Name      string
	Login     string
	Points    int64
	CreatedAt string
	UpdatedAt string
}
