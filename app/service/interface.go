/**
@author: 铁柱
@date:2021/6/30
@note:
**/
package service

import "github.com/gin-gonic/gin"

type ManageTestUser interface {
	Generate(quantity int) error
	ClearAll() error
}

type File interface {
	AddToIpfs(path, saveFileName string) (string, error)
	GetFileByCid(cid string) string
	Upload(context *gin.Context, savePath string) (r bool, finnalSavePath, saveFileName string)
}
