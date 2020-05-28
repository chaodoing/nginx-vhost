package main

import "fmt"

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

type Console struct {}

func (this *Console) textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}

func (this *Console) Black(str string) string {
	return this.textColor(textBlack, str)
}

func (this *Console) Red(str string) string {
	return this.textColor(textRed, str)
}
func (this *Console) Yellow(str string) string {
	return this.textColor(textYellow, str)
}
func (this *Console) Green(str string) string {
	return this.textColor(textGreen, str)
}
func (this *Console) Cyan(str string) string {
	return this.textColor(textCyan, str)
}
func (this *Console) Blue(str string) string {
	return this.textColor(textBlue, str)
}
func (this *Console) Purple(str string) string {
	return this.textColor(textPurple, str)
}
func (this *Console) White(str string) string {
	return this.textColor(textWhite, str)
}