package main

import (
	"log"
)

/*
go get gopkg.in/gographics/imagick.v2/imagick
Converting a PDF to JPG (using ImageMagick)
将pdf文件转为jpg文件
*/

func main() {
	pdfName := "test.pdf"
	jpgName := "test.jpg"
	if err := ConvertingPdfToJpg(pdfName, jpgName); err != nil {
		log.Fatal(err)
	}
}
func ConvertingPdfToJpg(pdfName string, jpgName string) error {

}
