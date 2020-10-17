package models

import "DataCertPlatfom/db_mysql"
import "fmt"
import "DataCertPlatfom/Hash"
type User struct {
	//Id int `form:"id"`
	Name string`form:"username"`
	Sex string `form:"sex"`
	Phone string`form:"phone"`
	Password string`form:"password1"`
	Email string `form:"email"`
}
//用户注册方法 注册成功返回rows
func (u User)AddUser()(int64,error){

	//1.密码脱敏 哈希包 md5加密再存入数据库
	u.Password = Hash.HASH(u.Password,"md5",false)
	//2.将用户数据存入数据库
	result,err := db_mysql.Db.Exec("insert into user_info(username,sex,phone,password,email)" +
		"value (?,?,?,?,?)",u.Name,u.Sex,u.Phone,u.Password,u.Email)
	if err != nil {
		return -1,err
	}
	//获取数据库的操作所影响的行数
	rows,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return rows,nil
}
//用户登录方法  登录是否成功取决于返回的bool返回的类型
func (u User)QueryUser() (*User,error) {
	u.Password = Hash.HASH(u.Password,"md5",false)
	//查询数据
	row := db_mysql.Db.QueryRow("select username,phone from user_info where username = ? and password = ?",
		u.Name, u.Password)

	err := row.Scan(&u.Name,&u.Phone)
	if err != nil {
		fmt.Println("rows.Sacn err :",err.Error())
		return nil,err

	}
	return  &u,nil
}
func QueryUserId(phone string)(int,error)  {
	 row := db_mysql.Db.QueryRow("select id from user_info where  phone = ? ",phone)
	 var userid int
	 if err :=row.Scan(&userid);err!=nil{
		 return -1,err

	 }
	 return  userid,nil
}