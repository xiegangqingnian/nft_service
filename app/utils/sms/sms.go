/**
@author: 铁柱
@date:2021/4/20
@note:
**/
package sms

import (
	"cw_backend/app/global/variable"
	"github.com/dev-submail/submail_go_sdk/submail/sms"
)

/**
 * @Description: 调用三方sdk 给指定号码发送 code
 * @param to
 * @param code
 * @return string
 */
func SmsSend(to, code string) string {

	res := sms.CreateXsend(variable.SmsConfig)

	res.SetTo(to)
	res.SetProject("xg0VI3")
	res.AddVar("code",code)
	res.AddVar("time","30")

	return res.Xsend()
}