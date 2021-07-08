/**
@author: 铁柱
@date:2021/6/28
@note:
**/
package model

import (
	"gorm.io/gorm"
	"nft_service/app/global/variable"
	"time"
)

// 待售列表
type TestUser struct {
	gorm.Model
	Addr string `gorm:"column:addr;size:42;comment:'地址'"`
	Key  string `gorm:"column:key;size:64;comment:'私钥'"`
}

func (u *TestUser) TableName() string {
	return "test_users"
}

func AddUsers(addrs *[]TestUser) (success error) {
	success = variable.GormDbMysql.Create(addrs).Error
	return
}

func CleanAll() error {
	var user TestUser

	return variable.GormDbMysql.Where("created_at < ?", time.Now()).Unscoped().Delete(&user).Error
}
