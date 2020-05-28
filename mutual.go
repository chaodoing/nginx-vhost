package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"os/exec"
)

var input = bufio.NewScanner(os.Stdin)

type Mutual struct{}

func (this *Mutual) isFileExist(path string) (bool) {
	//fileInfo, err := os.Stat(path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	//我这里判断了如果是0也算不存在
	//if fileInfo.Size() == 0 {
	//	return false
	//}
	if err == nil {
		return true
	}
	return false
}
func (this *Mutual) build() {
	var file string
	tpl, err := template.New("vhost").Parse(Vhost)
	if err != nil {
		fmt.Println("点配置文件", console.Red("[Fail]"))
		os.Exit(2)
	}
	file = configure.VHOST + "/" + configure.ServerName + ".conf"
	ioFile, ioErr := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if ioErr != nil {
		fmt.Println("点配置文件", console.Red("[Fail]"))
	}
	err = tpl.Execute(ioFile, StdIn{
		ServerName: configure.ServerName,
		RootPath: configure.RootPath,
		LogPath: configure.LogPath,
		ErrPath: configure.ErrorPath,
	})
	defer ioFile.Close()
}
func (this *Mutual) ErrorPage() {
	page.BadRequest(configure.RootPath + configure.ErrorPath + "/400.html")
	page.Unauthorized(configure.RootPath + configure.ErrorPath + "/401.html")
	page.Forbidden(configure.RootPath + configure.ErrorPath + "/403.html")
	page.NotFound(configure.RootPath + configure.ErrorPath + "/404.html")
	page.InternalServerError(configure.RootPath + configure.ErrorPath + "/500.html")
	page.NotImplemented(configure.RootPath + configure.ErrorPath + "/501.html")
	page.BadRequest(configure.RootPath + configure.ErrorPath + "/502.html")
	page.ServiceUnavailable(configure.RootPath + configure.ErrorPath + "/503.html")
	page.GatewayTimeout(configure.RootPath + configure.ErrorPath + "/504.html")
}
func (this *Mutual) ServerName() {
	fmt.Print(console.Green("请输入站点名称: "))
	input.Scan()
	configure.ServerName = input.Text()
	if len(configure.ServerName) <= 0 {
		fmt.Println("站点名称不能为空", console.Red("[Error]"))
		os.Exit(2)
	}
}
func (this *Mutual) RootPath() {
	fmt.Print(console.Green("请输入站点主路径:"))
	input.Scan()
	configure.RootPath = input.Text()
	if len(configure.RootPath) <= 0 {
		fmt.Println("站点主路径不能为空", console.Red("[Error]"))
		os.Exit(2)
	}
	err := os.MkdirAll(configure.RootPath+"/public", os.ModePerm)
	if err != nil {
		fmt.Println("创建站点目录", console.Red("[Fail]"))
		os.Exit(2)
	}
	fmt.Println("创建站点目录", console.Blue("[Ok]"))
}
func (this *Mutual) LogPath() {
	fmt.Print(console.Green("请输入日志目录名称:"))
	input.Scan()
	if len(input.Text()) <= 0 {
		fmt.Println("创建认日志目录", console.Yellow("[logs]"))
		configure.LogPath = configure.RootPath + "/logs"
	} else {
		configure.LogPath = configure.RootPath + "/" + input.Text()
	}
	err := os.MkdirAll(configure.LogPath, os.ModePerm)
	if err != nil {
		fmt.Println("创建认日志目录", configure.LogPath, console.Red("[Fail]"))
		os.Exit(2)
	}
	fmt.Println("创建认日志目录", configure.LogPath, console.Blue("[Ok]"))
}
func (this *Mutual) ErrorPath()  {
	fmt.Print(console.Green("请输入错误页面目录名称:"))
	input.Scan()
	if len(input.Text()) <= 0 {
		fmt.Println("使用默认错误目录", console.Yellow("[error_page]"))
		configure.ErrorPath = "/error_page"
	} else {
		configure.ErrorPath = "/" + input.Text()
	}
	err := os.MkdirAll(configure.RootPath + configure.ErrorPath, os.ModePerm)
	if err != nil {
		fmt.Println("创建错误页面目录", configure.ErrorPath, console.Red("[Fail]"))
		os.Exit(2)
	}
	fmt.Println("创建错误页面目录", configure.ErrorPath, console.Blue("[Ok]"))
}
func (this *Mutual) reloadNginx()  {
	out, err := exec.Command( configure.NGINX, "-t" ).Output()
	if err != nil {
		fmt.Println("检查nginx配置", console.Red("[Error]"))
	}
	fmt.Println(console.White("检查结果如下:"))
	os.Stdout.Write(out)
	out, err = exec.Command(configure.NGINX, "-s", "reload").Output()
	if err != nil {
		fmt.Println("重新启动Nginx", console.Red("[Fail]"))
	}
	os.Stdout.Write(out)
	fmt.Println("重新启动Nginx", console.Green("[Success]"))
}