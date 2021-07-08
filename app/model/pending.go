/**
@author: 铁柱
@date:2021/6/17
@note:
**/
package model

import "gorm.io/gorm"

// 待售列表
type Pending struct {
	gorm.Model
	Price string `gorm:"column:price;comment:'当前待售价格'"`
	Item1 string `gorm:"column:item1;comment:'备用1'"`
	Item2 string `gorm:"column:item2;comment:'备用2'"`
}

func (u *Pending) TableName() string {
	return "pendings"
}

//func GetPending() []Pending {
//
//}
