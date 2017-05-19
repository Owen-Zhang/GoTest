package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/axgle/mahonia"
)

func main() {
	//GetEnvirment()
	//ExecComandOpenFile()
	ExecCommandExeFile()
}

//GetEnvirment 获取环境变量
func GetEnvirment() {
	goroot := os.Getenv("GOROOT")
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	fmt.Println(goroot)
}

//ExecComandOpenFile 通过Command执行程序(用NotePad++打开一个文件)
func ExecComandOpenFile() {
	exeFilePath := "D:\\Program\\Notepad++\\notepad++.exe"
	targetPath := "D:\\Work\\Test\\WebAPI\\WebAPI-Test\\WebAPI-Test\\Web.config"
	comm := exec.Command(exeFilePath, targetPath)
	if err := comm.Run(); err == nil {
		fmt.Println("Sucess")
	} else {
		fmt.Println(err)
	}
}

//ExecCommandExeFile 通过Command执行exe程序，并传入相关的参数
func ExecCommandExeFile() {
	//在这个exe文件中，输出了中文，抛出了异常,以及将参数也输出了
	exeFilePath := "D:\\Work\\Test\\ConsoleTest\\ConsoleTest\\bin\\Debug\\ConsoleTest.exe"
	parameters := []string{"5678", "Test Parameters List"}
	//这里有个规则，如果参数支持任意个传入，在此可以用数组，只不过传入时要加上三个...
	comm := exec.Command(exeFilePath, parameters...)
	var outInfo bytes.Buffer
	var outErr bytes.Buffer

	comm.Stdout = &outInfo
	comm.Stderr = &outErr
	if err := comm.Run(); err == nil {
		fmt.Println("Sucess")
	} else {
		fmt.Println(err)
	}

	//使用第三方的库来处理中文乱码问题
	encoder := mahonia.NewDecoder("gbk")
	fmt.Println(encoder.ConvertString(outInfo.String()))
	fmt.Println(encoder.ConvertString(outErr.String()))
}
