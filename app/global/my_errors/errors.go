/**
@author: 铁柱
@date:2021/6/17
@note:
**/
package my_errors

const (
	//系统部分
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	ErrorsPublicNotExists           string = "public 目录不存在"
	ErrorsConfigYamlNotExists       string = "config.yml 配置文件不存在"
	ErrorsConfigGormNotExists       string = "gorm_v2.yml 配置文件不存在"
	ErrorsStorageLogsNotExists      string = "storage/logs 目录不存在"
	ErrorsConfigInitFail            string = "初始化配置文件发生错误"
	ErrorsSoftLinkCreateFail        string = "自动创建软连接失败,请以管理员身份运行客户端(开发环境为goland等，生产环境检查命令执行者权限)"
	ErrorsSoftLinkDeleteFail        string = "删除软软连接失败"

	ErrorsFuncEventAlreadyExists   string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister     string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall         string = "注册的函数无法正确执行"
	ErrorsBasePath                 string = "初始化项目根目录失败"
	ErrorsNoAuthorization          string = "token鉴权未通过，请通过token授权接口重新获取token,"
	ErrorsParseTokenFail           string = "解析token失败"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"
	ErrorsCasbinNoAuthorization    string = "Casbin 鉴权未通过，请在后台检查 casbin 设置参数"
	ErrorsGormNotInitGlobalPointer string = "%s 数据库全局变量指针没有初始化，请在配置文件 Gormv2.yml 设置 Gormv2.%s.IsInitGolobalGormMysql = 1, 并且保证数据库配置正确 \n"
	// 数据库部分
	ErrorsDbDriverNotExists   string = "数据库驱动类型不存在,目前支持的数据库类型：mysql、sqlserver、postgresql，您提交数据库类型："
	ErrorsDialectorDbInitFail string = "gorm dialector 初始化失败,dbType:"

	//redis部分
	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuthFail     string = "Redis Auth 鉴权失败，密码错误"
	ErrorsRedisGetConnFail  string = "Redis 从连接池获取一个连接失败，超过最大重试次数"
	// 验证器错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"
	//token部分
	ErrorsTokenInvalid      string = "无效的token"
	ErrorsTokenNotActiveYet string = "token 尚未激活"
	ErrorsTokenMalFormed    string = "token 格式不正确"

	ErrPrivateKeyToPublicKey string = "私钥转公钥错误"
	ErrGetIntance            string = "获取合约instance失败"
	ErrGetNonce              string = "获取nonce失败"
	ErrGetGasPrice           string = "获取gasprice"
	ErrSendContract          string = "发送合约交易失败"

	//snowflake
	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"
	// sms
	ErrorsSmsGetCodeFail string = "存在非法数字，无效的手机号"

	// websocket
	ErrorsWebsocketOnOpenFail                 string = "websocket onopen 发生阶段错误"
	ErrorsWebsocketUpgradeFail                string = "websocket Upgrade 协议升级, 发生错误"
	ErrorsWebsocketReadMessageFail            string = "websocket ReadPump(实时读取消息)协程出错"
	ErrorsWebsocketBeatHeartFail              string = "websocket BeatHeart心跳协程出错"
	ErrorsWebsocketBeatHeartsMoreThanMaxTimes string = "websocket BeatHeart 失败次数超过最大值"
	ErrorsWebsocketSetWriteDeadlineFail       string = "websocket  设置消息写入截止时间出错"
	ErrorsWebsocketWriteMgsFail               string = "websocket  Write Msg(send msg) 失败"
	// rabbitMq
	ErrorsRabbitMqReconnectFail string = "RabbitMq消费者端掉线后重连失败，超过尝试最大次数"

	//文件上传
	ErrorsFilesUploadOpenFail string = "打开文件失败，详情："
	ErrorsFilesUploadReadFail string = "读取文件32字节失败，详情："

	// casbin 初始化可能的错误
	ErrorCasbinCanNotUseDbPtr         string = "casbin 的初始化基于gorm 初始化后的数据库连接指针，程序检测到 gorm 连接指针无效，请检查数据库配置！"
	ErrorCasbinCreateAdaptFail        string = "casbin NewAdapterByDBUseTableName 发生错误："
	ErrorCasbinCreateEnforcerFail     string = "casbin NewEnforcer 发生错误："
	ErrorCasbinNewModelFromStringFail string = "NewModelFromString 调用时出错："

	// json格式化错误
	ErrorJsonUnmarshalFail string = "json数据Unmarshal失败"
	ErrorJsonmarshalFail   string = "json数据marshal失败"

	// redis
	ErrorRedisGetFail       string = "reids get 失败"
	ErrorRedisSetFail       string = "reids set 失败"
	ErrorRedisSetExpireFail string = "redis set expire 失败"

	ErrorRegisterCaptchaInvaild string = "短信验证码 无效"
	ErrorRegisterPassInvaild    string = "密码需字母数字组合，长度为8-20"

	ErrorMysqlSearchFail string = "数据库查找错误"

	ErrorsIDNumberFail string = "无效的身份证号"
)


