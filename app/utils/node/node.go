/**
@author: 铁柱
@date:2021/6/17
@note:
**/
package node

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"nft_service/app/global/variable"
	"nft_service/app/utils/yml_config"
	"nft_service/app/utils/yml_config/ymlconfig_interf"
)

// chain node
var NodeClient *ethclient.Client
var configYml ymlconfig_interf.YmlConfigInterf

func Init() {
	configYml = yml_config.CreateYamlFactory()
	NodeClient = createNodeClient()
}

func createNodeClient() *ethclient.Client {
	url := configYml.GetString("Node.Host") + ":" + configYml.GetString("Node.Port")
	Client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	variable.ZapLog.Info("connect node success, url : " + url)
	return Client
}
