/**
@author: 铁柱
@date:2021/6/17
@note:
**/
package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nft_service/app/global/consts"
	"nft_service/app/global/variable"
	"nft_service/app/model"
	service "nft_service/app/service"
	"nft_service/app/service/upload_file"

	//"nft_service/app/service/upload_file"
	"nft_service/app/utils/response"
)

func UploadFile(context *gin.Context) {
	savePath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")
	var service service.File
	service = &upload_file.FileService{}

	if r, finnalSavePath, _ := service.Upload(context, savePath); r == true {
		response.Success(context, consts.CurdStatusOkMsg, finnalSavePath)
	} else {
		response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
	}
}

// @Tags NFT
// add file to ipfs
// @Summary 上传文件到ipfs
// @Description 上传文件到ipfs 返回cid
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} string
// @Router /file/up-to-ipfs [post]
func AddToIpfs(context *gin.Context) {

	savePath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")

	var service service.File
	service = &upload_file.FileService{}
	var err error
	if r, finnalSavePath, saveFileName := service.Upload(context, savePath); r == true {
		cid, err := service.AddToIpfs(finnalSavePath, saveFileName)
		if err == nil {
			response.Success(context, consts.CurdStatusOkMsg, gin.H{
				"cid": cid,
			})
			return
		}
	}
	response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, err)
}

type DownloadParam struct {
	Cid string `json:"cid"`
}

// @Tags NFT
// @Summary 下载文件
// @Description 通过cid下载文件
// @Accept  json
// @Produce  json
// @Param param body DownloadParam  true "cid"
// @Success 200 {object}	gin.Context
// @Router /file/download [post]
func Download(context *gin.Context) {

	var param DownloadParam
	if err := context.BindJSON(&param); err != nil {
		response.ErrorParam(context, err)
		return
	}

	fileName, filePath, err := model.GetFileNamePathByCid(param.Cid)
	if err != nil || fileName == "" || filePath == "" {
		response.Fail(context, consts.FilesDownloadFailCode, consts.FilesDownloadFailMsg, err)
	}

	context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	context.Writer.Header().Add("Content-Type", "application/octet-stream")
	context.File("./" + filePath)
}
