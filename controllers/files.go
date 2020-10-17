package controllers

import (
	"DataCertPlatfom/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"strings"
	"time"
)

type FilesController struct {
	beego.Controller
}

func (f *FilesController) Get() {
	f.TplName = "files.html"

}

func (f *FilesController) Post() {

	title := f.Ctx.Request.PostFormValue("title")
	phone := f.Ctx.Request.PostFormValue("phone")
	fmt.Println(phone)
	//username := f.Ctx.Request.PostFormValue("phone")
 	userId,err := models.QueryUserId(phone)
 	if err != nil{
 		f.Ctx.WriteString("认证失败,请稍后重试!")
 		fmt.Println(err.Error())
		return
	}
	fmt.Println(title,phone)


	files, err := f.GetFiles("file")
	if err != nil {
		f.Data["Error"] = err.Error()
		f.TplName = "404.html"
		return
	}

	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")
	if err != nil {
		f.Data["Error"] = err.Error()
		f.TplName = "404.html"
		return
	}

	hashInterface := md5.New()
	//创建上传目录
	dirName := "static/upload"
	if _, err = os.Open(dirName);err!= nil{//先打开,没有则先创建
		if err = os.Mkdir(dirName, os.ModePerm);err != nil {
			f.Data["Error"] = err.Error()
			f.TplName = "404.html"
			return
		}
	}


	//判断每个文件的类型和大小。。
	for i := 0; i < len(files); i++ {
		isJpg := strings.HasSuffix(files[i].Filename, ".jpg")
		isPng := strings.HasSuffix(files[i].Filename, ".png")
		if !isJpg && !isPng {
			f.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
			return
		}
		if files[i].Size/1024 > fileSize {
			f.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
			return
		}

		//
		file, err := files[i].Open() //打开获取到的文件
		defer  file.Close()


		savePath := "static/upload" + "/" + files[i].Filename
		savefile, err := os.OpenFile(savePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			f.Ctx.WriteString("创建文件失败")
			fmt.Println(err.Error())
			return
		}
//		文件hash计算
			defer hashInterface.Reset()

		if _, err = io.Copy(savefile, file);err != nil {
			f.Data["Error"] = err.Error()
			f.TplName = "404.html"
			return
		}

		if _,err = io.Copy(hashInterface,savefile);err != nil{
			f.Data["Error"] = err.Error()
			f.TplName = "404.html"
			fmt.Println(err.Error())
			return
		}

		bytes := hashInterface.Sum(nil)
		//fmt.Println(hex.EncodeToString(bytes))

		thisFileInfo := models.UploadFile{
			//Id:        0,
			UserId:  userId  ,
			FileName:  files[i].Filename,
			FileSize:  files[i].Size,
			FileCert: hex.EncodeToString(bytes) ,
			FileTitle: title,
			CertTime:  time.Now().String(),
		}
		_,err = thisFileInfo.AddFiles()
		if err != nil{
			fmt.Println(err.Error())
			f.Ctx.WriteString("电子数据认证失败,请休息一会儿再试!")
			return
		}
	}
	records,err := models.QueryRecordsByUserId(userId)
	if err != nil{
		fmt.Println(err.Error())
		f.Ctx.WriteString("认证信息获取失败,请稍后重试!")
	}
	f.Data["Filesinfo"] = records
	f.TplName =" home.html"
	f.Ctx.WriteString("恭喜！！！ 提交完成！")
}

//file,header,err := f.GetFile("file")
//files := f.Ctx.Request.MultipartForm.File["file"]
//创建上传目录
//dirName := "static/upload"
//_, err = os.Open(dirName)//先打开,失败则拆先创建
//if err != nil {
//	err = os.Mkdir(dirName, os.ModePerm)
//	if err != nil {
//		f.Data["Error"] = err.Error()
//		f.TplName = "404.html"
//		return
//	}
//}
//file, err := files[i].Open()//打开获取到的文件

//保存多个文件时,使用 io.copy 避免内存读取负载
//destfile, err := os.Create(dirName + "/"+ files[i].Filename)
//if err != nil {
//	f.Data["Error"] = err.Error()
//	f.TplName = "404.html"
//}
//defer destfile.Close()

//files, header, err := f.Ctx.Request.FormFile("file")
//if err != nil {
//	f.Data["Error"] = err.Error()
//	f.TplName = "404.html"
//	return
//}
////f.Ctx.Request.FormFile("file")
//defer files.Close()
//fmt.Printf("%T\n", files)
//r.ParseMultipartForm(32 << 20)
//获取上传的文件组
//files := r.MultipartForm.File["uploadfile"]

//单个文件上传保存:
//fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
//// 创建保存单个文件
//destFile, err := os.Create("./upload/" + header.Filename)
//if err != nil {
//	log.Printf("Create failed: %s\n", err)
//	return
//}
//defer destFile.Close()
//
//// 读取表单文件，写入保存文件
//_, err = io.Copy(destFile, formFile)
//if err != nil {
//	log.Printf("Write file failed: %s\n", err)
//	return
//}
//fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的

//
//files := f.Ctx.Request.MultipartForm.File["file"]
//
//len := len(files)
//for i := 0; i < len; i++ {
//	//打开上传文件
//	file, err := files[i].Open()
//	defer file.Close()
//	if err != nil {
//		f.Data["Error"] = err.Error()
//		f.TplName = "404.html"
//		return
//	}
//	//创建上传目录
//	err = os.Mkdir("./upload", os.ModePerm)
//	if err != nil {
//		f.Data["Error"] = err.Error()
//		f.TplName = "404.html"
//		return
//	}
//	//创建上传文件
//	cur, err := os.Create("./upload/" + files[i].Filename)
//	if err != nil {
//		f.Data["Error"] = err.Error()
//		f.TplName = "404.html"
//		return
//	}
//	defer cur.Close()
//
//	_, err = io.Copy(cur, file)
//	if err != nil {
//		f.Data["Error"] = err.Error()
//		f.TplName = "404.html"
//		return
//	}
//	fmt.Println("文件", i, "的文件名:	\t", files[i].Filename)    //输出上传的文件名
//	fmt.Println("文件", i, "的文件大小:	\t", files[i].Size)       //输出上传文件的大小
//	fmt.Println("文件", i, "的文件Header:	\t", files[i].Header) //输出上传文件的大小
//	fmt.Println("文件", i, "的文件内存地址:	\t", &files[i])         //输出上传文件的大小
//}
//f.Ctx.WriteString("上传文件成功！")
