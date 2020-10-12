package db_mysql

import (
	"DataCertPlatfom/Hash"
	"DataCertPlatfom/models"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/msyql"
)
var Db *sql.DB
func OpenDB()  {
	if Db != nil{
		return
	}
	config := beego.AppConfig
	db_dirver_name := config.String("db_dirver_name")
	db_admin_name := config.String("db_admin_name")
	db_admin_pwd := config.String("db_admin_pwd")
	db_port := config.String("db_port")
	db_name := config.String("db_name")
	database,err := sql.Open(db_dirver_name,db_admin_name+":"+db_admin_pwd+"@tcp("+db_port+")/"+db_name+"?charset=utf8")
	//fmt.Println(db_dirver_name,db_admin_name+":"+db_admin_pwd+"@tcp("+db_port+")/"+db_name+"?charset=utf8")


	if err != nil {
		fmt.Println(err.Error())
	}
	Db =database
	//appName :=config.String("appname")
	//fmt.Println("程序名称",appName)
	//httpPort,err := config.Int("httpport")
	//if err != nil {
	//	panic("程序错误")
	//
	//}
	//fmt.Println("监听端口",httpPort)
	//"mysql", "root:123456@tcp(127.0.0.1:3306)/weather?charset=utf8"
	//fmt.Println(db)
	//"mysql", "root:123456@tcp(127.0.0.1:3306)/weather?charset=utf8"
	//			root:123456@tcp(127.0.0.1:3306)/weather?charset=utf8
	//fmt.Println(db_admin_name+":"+db_admin_pwd+"@tcp("+db_port+")/"+db_name+"?charset=utf8")
}
func Inseret(u models.User)(int64,error){
	u.Password = Hash.HASH(u.Password,"md5",false)
	result,err := Db.Exec("insert into user_info(username,sex,phone,password,email)" +
		"value (?,?,?,?,?)",u.Name,u.Sex,u.Phone,u.Password,u.Email)
	//result, err := Db.Exec("insert into user_info(phone,password,)"+
	//	"values(?,?)",u.Phone,u.Password,)
	if err != nil {
		return -1,err
	}
	rows,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return rows,nil
}
func Query_user_info(username,pwd string) (bool) {
	//查询数据
	rows := Db.QueryRow("select username,password from user_info where username = ? and password = ?",
		username,pwd)

	var user_db models.User

	err := rows.Scan(&user_db.Name,&user_db.Password)
	if err != nil {
		fmt.Println("rows.Sacn err :",err.Error())
	}
	//fmt.Println(user_db)
	if username ==user_db.Name && pwd == user_db.Password {
		return true
	}
	return false
}