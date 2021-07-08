/**
@author: 铁柱
@date:2021/7/1
@note:
**/
package model

import (
	"nft_service/app/global/variable"
	"time"
)

type File struct {
	Cid      string    `gorm:"primarykey;column:cid;size:42;comment:'cid'"`
	FileName string    `gorm:"column:filename;size:64;comment:'文件名'"`
	Path     string    `gorm:"column:path;size:64;comment:'存储路径'"`
	CreateAt time.Time `gorm:"column:create_at;size:64;comment:'创建时间'"`
}

func (u *File) TableName() string {
	return "files"
}

func AddFile(cid, fileName, path string) error {
	f := File{
		Cid:      cid,
		FileName: fileName,
		Path:     path,
		CreateAt: time.Now(),
	}
	return variable.GormDbMysql.Create(&f).Error
}

func GetFileNamePathByCid(cid string) (string, string, error) {
	var f File
	err := variable.GormDbMysql.Where("cid = ?", cid).First(&f).Error
	return f.FileName, f.Path, err
}
