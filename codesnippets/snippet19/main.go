package main

import (
	"log"
	"gopkg.in/gographics/imagick.v2/imagick"
)

/*
go get gopkg.in/gographics/imagick.v2/imagick
Converting a PDF to JPG (using ImageMagick)
将pdf文件转为jpg文件
*/

func main() {
	pdfName := "test.pdf"
	imgName := "test.jpg"
	if err := ConvertingPdfToJpg(pdfName, imgName); err != nil {
		log.Fatal(err)
	}
}
func ConvertingPdfToJpg(pdfName string, imgName string) error {

	//初始化
	imagick.Initialize()
	defer imagick.Terminate()

	//返回一个魔术棒对象，该对象提供所有API
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	//确保图片的质量
	if err := mw.SetResolution(300, 300); err != nil {
		return err
	}

	//读取pdf文件
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	//设置图片处理通道
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
		return err
	}

	//设置图片质量 100=max-quality
	if err := mw.SetCompressionQuality(95); err != nil {
		return err
	}

	//仅仅选择pdf的第一页
	mw.SetIteratorIndex(0)

	//设置格式
	if err := mw.SetFormat("jpg"); err != nil {
		return err
	}

	//保存图片
	return mw.WriteImage(imgName)
}
