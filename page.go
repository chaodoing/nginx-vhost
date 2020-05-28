package main

import (
	"fmt"
	"html/template"
	"os"
)

type Content struct {
	Title    string
	Keywords string
	Code     int
	Message  string
}
type Page struct{}

func (this *Page) render(file string, content Content) {
	tpl, err := template.New("page").Parse(ErrorPage)
	ioFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(console.Red("写入错误页面失败"))
		os.Exit(2)
	}
	tpl.Execute(ioFile, content)
	defer ioFile.Close()
	fmt.Println("创建页面[ ", console.Red(file), " ]成功")
}

// 400
func (this *Page) BadRequest(file string) {
	this.render(file, Content{
		Title:    "400 Bad Request",
		Keywords: "Bad Request",
		Code:     400,
		Message:  "400 Bad Request",
	})
}

// 401
func (this *Page) Unauthorized(file string) {
	this.render(file, Content{
		Title:    "401 Unauthorized",
		Keywords: "Unauthorized",
		Code:     401,
		Message:  "401 Unauthorized",
	})
}

// 403
func (this *Page) Forbidden(file string) {
	this.render(file, Content{
		Title:    "403 Forbidden",
		Keywords: "Forbidden",
		Code:     403,
		Message:  "403 Forbidden",
	})
}

// 404
func (this *Page) NotFound(file string) {
	this.render(file, Content{
		Title:    "404 Not Found",
		Keywords: "Not Found",
		Code:     404,
		Message:  "404 Not Found",
	})
}

// 500
func (this *Page) InternalServerError(file string) {
	this.render(file, Content{
		Title:    "500 Internal Server Error",
		Keywords: "Internal Server Error",
		Code:     500,
		Message:  "500 Internal Server Error",
	})
}

// 501
func (this *Page) NotImplemented(file string) {
	this.render(file, Content{
		Title:    "501 Not Implemented",
		Keywords: "Not Implemented",
		Code:     501,
		Message:  "501 Not Implemented",
	})
}

// 502
func (this *Page) BadGateway(file string) {
	this.render(file, Content{
		Title:    "502 Bad Gateway",
		Keywords: "Bad Gateway",
		Code:     502,
		Message:  "502 Bad Gateway",
	})
}

// 503
func (this *Page) ServiceUnavailable(file string) {
	this.render(file, Content{
		Title:    "503 Service Unavailable",
		Keywords: "Service Unavailable",
		Code:     503,
		Message:  "503 Service Unavailable",
	})
}

// 504
func (this *Page) GatewayTimeout(file string) {
	this.render(file, Content{
		Title:    "504 Gateway Timeout",
		Keywords: "Gateway Timeout",
		Code:     504,
		Message:  "504 Gateway Timeout",
	})
}
