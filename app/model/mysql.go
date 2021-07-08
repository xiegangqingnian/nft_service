/**
@author: 铁柱
@date:2021/6/24
@note:
**/
package model

import "nft_service/app/global/variable"

const (
	_ = iota
	AllNft
	AllInSelling
	AllInGame
	AllByOwner
)

// init table
func Init() {
	if !variable.GormDbMysql.Migrator().HasTable(&Nft{}) {
		err := variable.GormDbMysql.AutoMigrate(&Nft{}, &Pending{}, &MyCollect{}, &MyGame{}, &TestUser{}, &File{})
		if err != nil {
			panic(err)
		}
	}
}
