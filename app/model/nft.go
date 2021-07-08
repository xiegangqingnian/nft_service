/**
@author: 铁柱
@date:2021/6/18
@note:
**/
package model

import (
	"nft_service/app/global/variable"
	"time"
)

// all nft
type Nft struct {
	TokenID      string `gorm:"primarykey;column:token_id;comment:'token id'"`
	Owner        string `gorm:"column:owner;comment:'所属者'"`
	Cid          string `gorm:"column:cid;comment:'cid'"`
	ApproveCount string `gorm:"column:approve_count;comment:'赞数'"`
	DefuseCount  string `gorm:"column:defuse_count;comment:'踩数'"`
	IsRunning    bool   `gorm:"column:is_running;comment:'是否在游戏中'"`
	IsOnSelling  bool   `gorm:"column:is_on_selling;comment:'是否代售中'"`
	Amount       string `gorm:"column:amount;comment:'nft绑定金额'"`
	Round        uint64 `gorm:"column:round;comment:'nft轮次'"`
	Price        string `gorm:"column:price;comment:'nft 当前售卖价格'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *Nft) TableName() string {
	return "nfts"
}

func AddNft(tokenId, owner, cid string) error {
	var nft = Nft{
		TokenID:   tokenId,
		Owner:     owner,
		Cid:       cid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return variable.GormDbMysql.Create(&nft).Error
}

func CheckNftExistByTokenId(tokenId string) bool {
	var nfts []Nft
	variable.GormDbMysql.Where("token_id = ?", tokenId).Find(&nfts)
	if len(nfts) > 0 {
		return true
	}
	return false
}

func GetAllNft(page, pageSize int) interface{} {
	var nfts []Nft
	variable.GormDbMysql.Limit(pageSize).Offset((page - 1) * pageSize).Order("created_at desc").Find(&nfts)
	return nfts
}

func GetAllNftInSell(page, pageSize int) interface{} {
	var nfts []Nft
	variable.GormDbMysql.Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Where("is_on_selling = ?", true).Find(&nfts)
	return nfts
}

func GetAllNftInGame(page, pageSize int) interface{} {
	var nfts []Nft
	variable.GormDbMysql.Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Where("is_running = ?", true).Find(&nfts)
	return nfts
}

func GetAllNftByOwner(page, pageSize int, owner string) interface{} {
	var nfts []Nft
	variable.GormDbMysql.Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Where("owner = ?", true).Find(&nfts)
	return nfts
}
