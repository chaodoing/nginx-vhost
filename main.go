package main

import (
	"flag"
	"fmt"
	"os"
)

// 控制台输出颜色
var console Console = Console{}
var page Page = Page{}
var configure Configure = Configure{
	NGINX: "/opt/nginx/sbin/nginx",
	VHOST: "/opt/nginx/vhost/",
}
var mutual Mutual = Mutual{}

var (
	vhost  string
	nginx string
	help   bool
)

func main() {
	flag.StringVar(&vhost, "v", configure.VHOST, "配置 nginx vhost 目录位置")
	flag.BoolVar(&help, "h", false, "查看帮助信息")
	flag.StringVar(&nginx, "nginx", configure.NGINX, "配置 nginx 可执行文件位置")
	flag.Parse()
	configure.VHOST = vhost
	configure.NGINX = nginx
	if !mutual.isFileExist(configure.VHOST) {
		fmt.Println("虚拟站点配置目录找不到", console.Yellow("[-v vhost_dir]"))
		os.Exit(2)
	}
	if !mutual.isFileExist(configure.NGINX) {
		fmt.Println(console.Red("nginx运行文件找不到"))
		os.Exit(2)
	}
	if help {
		flag.PrintDefaults()
	}
	mutual.ServerName()
	mutual.RootPath()
	mutual.ErrorPath()
	mutual.ErrorPage()
	mutual.LogPath()
	mutual.build()
	mutual.reloadNginx()
}
