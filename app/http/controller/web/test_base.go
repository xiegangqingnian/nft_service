/**
@author: 铁柱
@date:2021/6/30
@note:
**/
package web

import (
	"github.com/gin-gonic/gin"
	"nft_service/app/global/consts"
	service2 "nft_service/app/service"
	"nft_service/app/service/contract_event"
	"nft_service/app/utils/response"
)

type GenerateParam struct {
	Quantity int `json:"quantity"`
}

// @Tags Test Base
// Generate Addr
// @Summary 生成测试账户
// @Description 生成测试账户
// @Accept  json
// @Produce  json
// @Param quantity body GenerateParam true "生成地址数量"
// @Success 200 {object} string
// @Router /generate [post]
func GenerateTestUsers(c *gin.Context) {

	var param GenerateParam
	if err := c.BindJSON(&param); err != nil {
		response.ErrorParam(c, err)
		return
	}
	var service service2.ManageTestUser
	service = &contract_event.TestUserService{}

	err := service.Generate(param.Quantity)
	if err != nil {
		response.ErrorSystem(c, "", err)
		return
	}

	response.Success(c, consts.CurdStatusOkMsg, "")
}

// @Tags Test Base
// Generate Addr
// @Summary 清除测试账户
// @Description 清除测试账户
// @Success 200 {object} string
// @Router /clean-all-user [post]
func CleanAllUser(c *gin.Context) {

	var service service2.ManageTestUser
	service = &contract_event.TestUserService{}

	err := service.ClearAll()
	if err != nil {
		response.ErrorSystem(c, "", err)
		return
	}

	response.Success(c, consts.CurdStatusOkMsg, "")
}
