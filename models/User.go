package models

type User struct {
	//Id int `form:"id"`
	Name string`form:"username"`
	Sex string `form:"sex"`
	Phone string`form:"phone"`
	Password string`form:"password1"`
	Email string `form:"email"`
}
