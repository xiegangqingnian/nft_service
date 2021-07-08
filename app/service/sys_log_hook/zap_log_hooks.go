/**
@author: 铁柱
@date:2021/6/17
@note: 日志钩子
**/
package sys_log_hook

import (
	 "go.uber.org/zap/zapcore"
)


/*
	// 参数 entry 介绍
	// entry  参数就是单条日志结构体，主要包括字段如下：
	// Level      日志等级
	// Time       当前时间
	// LoggerName  日志名称
	// Message    日志内容
	// Caller     各个文件调用路径
	// Stack      代码调用栈
*/
func ZapLogHandler(entry zapcore.Entry) error {

	go func(paramEntry zapcore.Entry) {
		//fmt.Println(" 可以在这里继续处理系统日志....")
		//fmt.Printf("%#+v\n", paramEntry)
	}(entry)
	return nil
}