package upload

import (
	"github.com/gin-gonic/gin"
	"ginCms/controllers"
	"ginCms/comm"
	"ginCms/comm/setting"
	"ginCms/utils"
)

func ImageUpload(c *gin.Context) {
	ctx := controllers.Context{c}
	file, err := ctx.FormFile("file")
	if err != nil {
		comm.Log("error").Println(err)
		ctx.Fail(500, "10001", "上传失败")
	}
	//fmt.Println(file.Filename)
	//fmt.Println(file.Size)
	imgPath := setting.AppSetting.ImageUploadPath
	if t, _ := utils.CheckCreateDir(imgPath); t {
		ctx.SaveUploadedFile(file, imgPath+"/"+file.Filename)
	} else {
		ctx.Fail(500, "10001", "图片上传失败！")
	}
}
