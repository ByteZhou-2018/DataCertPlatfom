package models

type User struct {
	Id int `form:"id"`
	Name string `form:"name"`
	Phone string`form:"phone"`
	Password string`form:"password"`
}
