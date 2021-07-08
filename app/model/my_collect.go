/**
@author: 铁柱
@date:2021/6/23
@note:
**/
package model

import "gorm.io/gorm"

// 收藏
type MyCollect struct {
	gorm.Model
	Address string `gorm:"column:address;comment:'地址'"`
	NftDbId uint64 `gorm:"column:nft_db_id;comment:'nft id'"`
	TokenID string `gorm:"column:token_id;comment:'token id'"`
}

func (u *MyCollect) TableName() string {
	return "my_collects"
}
