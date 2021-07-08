/**
@author: 铁柱
@date:2021/6/17
@note: init system
**/
package bootstrap

import (
	"log"
	"nft_service/app/global/my_errors"
	"nft_service/app/global/variable"
	"nft_service/app/model"
	"nft_service/app/service/contract_event"
	"nft_service/app/service/sys_log_hook"
	"nft_service/app/utils/casbin_v2"
	"nft_service/app/utils/gorm_v2"
	"nft_service/app/utils/node"
	"nft_service/app/utils/snow_flake"
	"nft_service/app/utils/websocket/core"
	"nft_service/app/utils/yml_config"
	"nft_service/app/utils/zap_factory"
	"os"
)

// check all file,such as  log
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	if _, err := os.Stat(variable.BasePath + "/config/gorm_v2.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigGormNotExists + err.Error())
	}
	//2.检查public目录是否存在
	if _, err := os.Stat(variable.BasePath + "/public/"); err != nil {
		log.Fatal(my_errors.ErrorsPublicNotExists + err.Error())
	}
	//3.检查storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}
	// 4.自动创建软连接、更好的管理静态资源
	if _, err := os.Stat(variable.BasePath + "/public/storage"); err == nil {
		if err = os.Remove(variable.BasePath + "/public/storage"); err != nil {
			log.Fatal(my_errors.ErrorsSoftLinkDeleteFail + err.Error())
		}
	}
	if err := os.Symlink(variable.BasePath+"/storage/app", variable.BasePath+"/public/storage"); err != nil {
		log.Fatal(my_errors.ErrorsSoftLinkCreateFail + err.Error())
	}
}

func init() {

	checkRequiredFolders()

	//启动针对配置文件(confgi.yml、gorm_v2.yml)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYml = yml_config.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()
	// config>gorm_v2.yml 启动文件变化监听事件
	variable.ConfigGormv2Yml = variable.ConfigYml.Clone("gorm_v2")
	variable.ConfigGormv2Yml.ConfigFileChangeListen()

	// 全局ZapLog对象，并载入钩子可处理后续日志
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	// init gorm mysql 全局 *gorm.Db
	if variable.ConfigGormv2Yml.GetInt("Gormv2.Mysql.IsInitGolobalGormMysql") == 1 {
		if dbMysql, err := gorm_v2.GetOneMysqlClient(); err != nil {
			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbMysql = dbMysql
			//create table
			model.Init()
		}
	}

	// SnowFlake
	variable.SnowFlake = snow_flake.CreateSnowflakeFactory()

	// websocket Hub中心启动
	if variable.ConfigYml.GetInt("Websocket.Start") == 1 {
		// websocket 管理中心hub全局初始化一份
		variable.WebsocketHub = core.CreateHubFactory()
		if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
			go Wh.Run()
		}
	}

	// SaiYou sms

	variable.SmsConfig = make(map[string]string)
	variable.SmsConfig["appid"] = variable.ConfigYml.GetString("Sms.Appid")
	variable.SmsConfig["appkey"] = variable.ConfigYml.GetString("Sms.Appkey")
	variable.SmsConfig["signType"] = variable.ConfigYml.GetString("Sms.SignType")
	variable.SmsLength = variable.ConfigYml.GetInt("Sms.Length")

	// 10 casbin 依据配置文件设置参数(IsInit=1)初始化
	if variable.ConfigYml.GetInt("Casbin.IsInit") == 1 {
		var err error
		if variable.Enforcer, err = casbin_v2.InitCasbinEnforcer(); err != nil {
			log.Fatal(err.Error())
		}
	}

	// node

	// ipfs
	//if variable.ConfigYml.GetInt("IPFS.IsInit") == 1 {
	//	var err error
	//	if variable.Shell,err = ipfs.InitIpfs(variable.ConfigYml.GetString("Ipfs.Host")); err != nil {
	//		log.Fatal(err.Error())
	//	}
	//}

	variable.ContractAddr = variable.ConfigYml.GetString("Node.ContractAddr")

	variable.IpfsHost = variable.ConfigYml.GetString("Ipfs.Host")
	variable.InfuraIpfsProjectID = variable.ConfigYml.GetString("Ipfs.InfuraIpfsProjectID")
	variable.InfuraProjectSecret = variable.ConfigYml.GetString("Ipfs.InfuraProjectSecret")

	// eth node
	node.Init()

	go contract_event.GetAllEventLog()
}
