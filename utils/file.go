package utils

import (
	"io"
	"os"
)

//打开并创建目录
func OpenDir(dirPath string) (error) {
	_,err := os.Open(dirPath)
	if err != nil {
		if err := os.Mkdir(dirPath,os.ModePerm);err != nil{
			return err
		}
	}
	return nil
}
func  SaveFile(filename string,file io.Reader)(int64,error)  {
	savefile,err :=os.OpenFile(filename,os.O_CREATE|os.O_RDWR,os.ModePerm)
	if err != nil{
		return -1,err
	}
	length,err := io.Copy(savefile,file)
	return length,err
}
//func OpenFile(filePath string)(io.Reader,error) {
//	file,err :=os.OpenFile(filePath,os.O_CREATE|os.O_RDWR,os.ModePerm)
//	if err != nil{
//		return nil,err
//	}
//	return file,nil
//}
