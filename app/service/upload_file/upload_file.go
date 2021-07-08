/**
@author: 铁柱
@date:2021/6/17
@note:
**/
package upload_file

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"nft_service/app/global/my_errors"
	"nft_service/app/global/variable"
	"nft_service/app/model"
	"nft_service/app/utils/ipfs"
	"nft_service/app/utils/md5_encrypt"
	"os"
	"path"
	"strings"
	"time"
)

type FileService struct {
}

func (fs *FileService) Upload(context *gin.Context, savePath string) (r bool, finnalSavePath, saveFileName string) {

	newSavePath, newReturnPath := generateYearMonthPath(savePath)

	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）

	//  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var saveErr error
	if sequence := variable.SnowFlake.GetId(); sequence > 0 {
		saveFileName := fmt.Sprintf("%d%s", sequence, file.Filename)
		saveFileName = md5_encrypt.MD5(saveFileName) + path.Ext(saveFileName)

		if saveErr = context.SaveUploadedFile(file, newSavePath+saveFileName); saveErr == nil {
			//  上传成功,返回资源的相对路径，这里请根据实际返回绝对路径或者相对路径
			//finnalSavePath = gin.H{
			//	"path": strings.ReplaceAll(newReturnPath+saveFileName, variable.BasePath, ""),
			//}
			finnalSavePath = strings.ReplaceAll(newReturnPath+saveFileName, variable.BasePath, "")
			return true, finnalSavePath, saveFileName
		}
	} else {
		saveErr = errors.New(my_errors.ErrorsSnowflakeGetIdFail)
		variable.ZapLog.Error("文件保存出错：" + saveErr.Error())
	}
	return false, "", ""

}

func (fs *FileService) AddToIpfs(path, saveFileName string) (string, error) {

	sh := ipfs.GetIpfsClient()
	f, err := os.Open("./" + path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	cid, err := sh.Add(f)
	if err != nil {
		return cid, err
	}

	err = sh.FilesCp(context.Background(), "/ipfs/"+cid, "/"+saveFileName)
	if err != nil {
		return cid, err
	}
	fmt.Println("cid***********", cid)
	err = model.AddFile(cid, saveFileName, path)

	return cid, err
}

func (fs *FileService) GetFileByCid(cid string) string {
	return ""
}

// 文件上传可以设置按照 xxx年-xx月 格式存储
func generateYearMonthPath(savePathPre string) (string, string) {
	returnPath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileReturnPath")
	curYearMonth := time.Now().Format("2006_01")
	newSavePathPre := savePathPre + curYearMonth
	newReturnPathPre := returnPath + curYearMonth
	// 相关路径不存在，创建目录
	if _, err := os.Stat(newSavePathPre); err != nil {
		if err = os.MkdirAll(newSavePathPre, 0666); err != nil {
			variable.ZapLog.Error("文件上传创建目录出错" + err.Error())
			return "", ""
		}
	}
	return newSavePathPre + "/", newReturnPathPre + "/"
}
