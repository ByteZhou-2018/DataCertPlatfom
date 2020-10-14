package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

type FilesController struct {
	beego.Controller
}

func (f FilesController) Get() {
	f.TplName = "files.html"
}
func (f FilesController) Post() {

	//file,header,err := f.GetFile("file")
	//file,header,err := f.Ctx.Request.FormFile("file")

	//files,err := f.GetFiles("file")

	files := f.Ctx.Request.MultipartForm.File["file"]

	//  f.Ctx.Request.Form
	f.Ctx.Request.FormFile("file")//获取单个上传的文件
	len := len(files)
	for i := 0; i < len; i++ {
		//打开上传文件
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			f.Data["Error"] = err.Error()
			f.TplName = "404.html"
			return
		}
		//创建上传目录
		 err = os.Mkdir("./upload", os.ModePerm)
		if err != nil {
			f.Data["Error"] = err.Error()
			f.TplName = "404.html"
			return
		}
		//创建上传文件
		cur, err := os.Create("./upload/" + files[i].Filename)

		defer cur.Close()
		if err != nil {
			//log.Fatal(err)
			f.Data["Error"] = err.Error()
			f.TplName = "404.html"
			return
		}

		io.Copy(cur, file)
		fmt.Println("文件",i,"的文件名:	\t",files[i].Filename) //输出上传的文件名
		fmt.Println("文件",i,"的文件大小:	\t",files[i].Size)//输出上传文件的大小
		fmt.Println("文件",i,"的文件Header:	\t",files[i].Header)//输出上传文件的大小
	fmt.Println("文件",i,"的文件内存地址:	\t",&files[i])//输出上传文件的大小
}


	f.Ctx.WriteString("提交文件成功！")


}

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