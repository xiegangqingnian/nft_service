/**
@author: 铁柱
@date:2021/6/24
@note:
**/
package nft

import (
	"errors"
	"nft_service/app/model"
)

/**
 * @Description: add nft to db
 * @param tokenId
 * @param owner
 * @param cid
 * @return bool
 */
func AddNft(tokenId, owner, cid string) error {
	if model.CheckNftExistByTokenId(tokenId) {
		return errors.New("tokenId existed")
	}
	return model.AddNft(tokenId, owner, cid)
}

/**
 * @Description:  search interface
 * @param page
 * @param pageSize
 * @param typa	1 all nft   2 in sell   3 in game  4 by onwer
 * @param key     typa == 4  ,key is nft owner
 * @return interface{}
 */
func SearchNft(page, pageSize, typa int, key ...string) (res interface{}) {

	switch typa {
	case model.AllNft:
		res = model.GetAllNft(page, pageSize)
	case model.AllInSelling:
		res = model.GetAllNftInSell(page, pageSize)
	case model.AllInGame:
		res = model.GetAllNftInGame(page, pageSize)
	case model.AllByOwner:
		res = model.GetAllNftByOwner(page, pageSize, key[0])
	default:
		res = ""
	}
	return
}
