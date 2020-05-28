package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
)

var (
	VHOST = "/opt/nginx/vhost/"
	NGINX = "/opt/nginx/sbin/nginx"
)

var err error
var input = bufio.NewScanner(os.Stdin)

func main() {
	if len(os.Args) <= 1 {
		siteName()
		path()
		errpage()
		createErrorPage()
		logpath()
		create()
	}
}
func createErrorPage() {
	var (
		file400 string = Configure.Error + "/400.html"
		file401 string = Configure.Error + "/401.html"
		file403 string = Configure.Error + "/403.html"
		file404 string = Configure.Error + "/404.html"
		file500 string = Configure.Error + "/500.html"
		file501 string = Configure.Error + "/501.html"
		file502 string = Configure.Error + "/502.html"
		file503 string = Configure.Error + "/503.html"
		file504 string = Configure.Error + "/504.html"
	)
	render(file400, NotFound)
	render(file401, Unauthorized)
	render(file403, Forbidden)
	render(file404, NotFound)
	render(file500, InternalServerError)
	render(file501, NotImplemented)
	render(file502, BadGateway)
	render(file503, ServiceUnavailable)
	render(file504, GatewayTimeout)
}
func render(file string, content pageContent) {
	tpl, err    := template.New("errpage").Parse(ErrorPage)
	ioFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(Red("写入错误页面失败"))
	}
	tpl.Execute(ioFile, content)
	defer ioFile.Close()
}
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	
	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

func create() {
	var (
		file string
	)
	tpl, err := template.New("vhost").Parse(Nginx)
	if err != nil {
		fmt.Println(Red("解析站点配置文件失败"))
	}
	file = VHOST + "/" + Configure.Site + ".conf"
	
	fileio, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	
	if err != nil {
		fmt.Println(Red("读写站点配置文件失败"))
	}
	err = tpl.Execute(fileio, Configure)
	if err != nil {
		fmt.Println(Red("创建站点配置文件失败"))
	}
	defer fileio.Close()
	fmt.Println("创建站点", Blue("[OK]"))
	fmt.Println(Yellow("现在您可以重启站点"))
}

func siteName() {
	fmt.Print(Green("请输入站点名称:"))
	input.Scan()
	Configure.Site = input.Text()
	
	if len(Configure.Site) <= 0 {
		fmt.Println("站点名称不能为空", Red("[Error]"))
		os.Exit(2)
	}
}
func path() {
	fmt.Print(Green("请输入站点主路径:"))
	input.Scan()
	Configure.Path = input.Text()
	if len(Configure.Path) <= 0 {
		fmt.Println("站点主路径不能为空", Red("[Error]"))
		os.Exit(2)
	}
	err = os.MkdirAll(Configure.Path+"/public", os.ModePerm)
	if err != nil {
		fmt.Println("创建站点目录", Red("[Fail]"))
		os.Exit(2)
	}
	fmt.Println("创建站点目录", Blue("[Ok]"))
}
func errpage() {
	fmt.Print(Green("请输入错误页面目录名称:"))
	input.Scan()
	if len(input.Text()) <= 0 {
		fmt.Println("使用默认错误目录", Yellow("[error_page]"))
		Configure.Error = Configure.Path + "/error_page"
	} else {
		Configure.Error = Configure.Path + "/" + input.Text()
	}
	
	err = os.MkdirAll(Configure.Error, os.ModePerm)
	if err != nil {
		fmt.Println("创建错误页面目录", Configure.Error, Red("[Fail]"))
		os.Exit(2)
	}
	fmt.Println("创建错误页面目录", Configure.Error, Blue("[Ok]"))
}
func logpath() {
	
	fmt.Print(Green("请输入日志目录名称:"))
	input.Scan()
	if len(input.Text()) <= 0 {
		fmt.Println("创建认日志目录", Yellow("[logs]"))
		Configure.Log = Configure.Path + "/logs"
	} else {
		Configure.Log = Configure.Path + "/" + input.Text()
	}
	err = os.MkdirAll(Configure.Log, os.ModePerm)
	if err != nil {
		fmt.Println("创建认日志目录", Configure.Log, Red("[Fail]"))
		os.Exit(2)
	}
	fmt.Println("创建认日志目录", Configure.Log, Blue("[Ok]"))
}
