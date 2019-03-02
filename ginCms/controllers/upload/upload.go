package upload

import (
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/humorliang/go-demo/ginCms/controllers"
	"github.com/humorliang/go-demo/ginCms/comm"
	"github.com/humorliang/go-demo/ginCms/comm/setting"
	"github.com/humorliang/go-demo/ginCms/utils"
)

func ImageUpload(c *gin.Context) {
	ctx := controllers.Context{c}
	file, err := ctx.FormFile("file")
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "上传失败")
	}
	fmt.Println(file.Filename)
	//fmt.Println(file.Size)
	imgPath := setting.AppSetting.ImageUploadPath
	//if v, ok := ctx.Get("userId"); ok {
	//	if t, _ := utils.CheckCreateDir(imgPath); t {
	//		ctx.SaveUploadedFile(file, imgPath+"/"+fmt.Sprint(v)+"_"+file.Filename)
	//	} else {
	//		ctx.Fail(500, "10001", "图片上传失败！")
	//	}
	//} else {
	//	ctx.Fail(400, "10001", "获取用户Id失败！")
	//}
	fmt.Println(imgPath)
	if t, _ := utils.CheckCreateDir(imgPath); t {
		//保存的文件名
		newFileStr := imgPath + "/" + file.Filename
		err := ctx.SaveUploadedFile(file, newFileStr)
		if err != nil {
			ctx.Fail(500, "10001", "图片上传失败！")
			return
		}

	} else {
		ctx.Fail(500, "10001", "图片上传失败！")
	}

}
