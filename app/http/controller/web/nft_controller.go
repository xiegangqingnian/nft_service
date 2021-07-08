package web

import (
	"github.com/gin-gonic/gin"
	"nft_service/app/global/consts"
	"nft_service/app/service/nft"
	"nft_service/app/utils/response"
)

type AddNftParam struct {
	TokenID string `json:"token_id" binding:"required"`
	Owner   string `json:"owner" binding:"required"`
	Cid     string `json:"cid" binding:"required"`
}

// @Tags NFT
// Add NFT TO DB
// @Summary 存入nft信息
// @Description 存入nft信息
// @Accept  json
// @Produce  json
// @Param param body AddNftParam true "add nft"
// @Success 200 {object} string
// @Router /nft/add [post]
func AddNft(c *gin.Context) {
	var param AddNftParam
	if err := c.BindJSON(&param); err != nil {
		response.ErrorParam(c, err)
		return
	}
	if err := nft.AddNft(param.TokenID, param.Owner, param.Cid); err != nil {
		response.Fail(c, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, err)
	} else {
		response.Success(c, consts.CurdStatusOkMsg, "")
	}
}

type SearchParm struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
	Typa     int    `json:"typa" binding:"required"`
	Key      string `json:"key"`
}

// @Tags NFT
// Get NFT
// @Summary 获取nft信息
// @Description 获取nft信息  typa	1 查询所有的nft   2 所有在售的nft   3所有在游戏中的  4根据key查询自己的nft
// @Accept  json
// @Produce  json
// @Param param body SearchParm true "nft id"
// @Success 200 {object} string
// @Router /nft/search [post]
func GetAllNft(c *gin.Context) {
	var param SearchParm
	if err := c.BindJSON(&param); err != nil {
		response.ErrorParam(c, err)
		return
	}
	if param.Typa == 4 && param.Key == "" {
		response.ErrorParam(c, "invalid key")
		return
	}

	var res interface{}
	res = nft.SearchNft(param.Page, param.PageSize, param.Typa, param.Key)

	response.Success(c, consts.CurdStatusOkMsg, res)
}
