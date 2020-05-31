package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

//go version 来测试有没有安装好go。
// go build
//要找到项目目录，然后执行。
//main.go是入口文件，main()函数是程序的入口。
//cd是change dirctory 改变文件夹的意思。
//mac中执行 ./编译后的文件名。
//go env -w GOPROXY=https://goproxy.cn,direct 设置代理服务器，下载开发包。
//在其他路径下执行go build,要在后面加上项目的路径（从src后开始到项目文件夹，不包扩src）。
