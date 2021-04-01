package user

import (
	"github.com/1024casts/snake/internal/ecode"
	"github.com/1024casts/snake/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/1024casts/snake/pkg/log"
)

// VCode 获取验证码
// @Summary 根据手机号获取校验码
// @Description Get an user by username
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param area_code query string true "区域码，比如86"
// @Param phone query string true "手机号"
// @Success 200 {object} app.Response
// @Router /vcode [get]
func VCode(c *gin.Context) {
	// 验证区号和手机号是否为空
	if c.Query("area_code") == "" {
		log.Warn("vcode area code is empty")
		Response.Error(c, ecode.ErrAreaCodeEmpty)
		return
	}

	phone := c.Query("phone")
	if phone == "" {
		log.Warn("vcode phone is empty")
		Response.Error(c, ecode.ErrPhoneEmpty)
		return
	}

	// TODO: 频率控制，以防攻击

	// 生成短信验证码
	verifyCode, err := service.VCodeSvc.GenLoginVCode(phone)
	if err != nil {
		log.Warnf("gen login verify code err, %v", errors.WithStack(err))
		Response.Error(c, ecode.ErrGenVCode)
		return
	}

	// 发送短信
	err = service.VCodeSvc.SendSMS(phone, verifyCode)
	if err != nil {
		log.Warnf("send phone sms err, %v", errors.WithStack(err))
		Response.Error(c, ecode.ErrSendSMS)
		return
	}

	Response.Success(c, nil)
}
