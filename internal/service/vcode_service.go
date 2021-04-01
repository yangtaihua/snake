package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/1024casts/snake/pkg/log"
	"github.com/1024casts/snake/pkg/redis"
)

// 验证码服务，主要提供生成验证码和获取验证码

const (
	verifyCodeRedisKey = "app:login:vcode:%d" // 验证码key
	maxDurationTime    = 10 * time.Minute     // 验证码有效期
)

// GenLoginVCode 生成校验码
func (s *Service) GenLoginVCode(phone string) (int, error) {
	// step1: 生成随机数
	vCodeStr := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// step2: 写入到redis里
	// 使用set, key使用前缀+手机号 缓存10分钟）
	key := fmt.Sprintf("app:login:vcode:%s", phone)
	err := redis.RedisClient.Set(key, vCodeStr, maxDurationTime).Err()
	if err != nil {
		return 0, errors.Wrap(err, "gen login code from redis set err")
	}

	vCode, err := strconv.Atoi(vCodeStr)
	if err != nil {
		return 0, errors.Wrap(err, "string convert int err")
	}

	return vCode, nil
}

// 手机白名单
var phoneWhiteLit = []int64{
	13010102020,
}

// isTestPhone 这里可以添加测试号，直接通过
func isTestPhone(phone int64) bool {
	for _, val := range phoneWhiteLit {
		if val == phone {
			return true
		}
	}
	return false
}

// CheckLoginVCode 验证校验码是否正确
func (s *Service) CheckLoginVCode(phone int64, vCode int) bool {
	if isTestPhone(phone) {
		return true
	}

	oldVCode, err := s.GetLoginVCode(phone)
	if err != nil {
		log.Warnf("[vcode_service] get verify code err, %v", err)
		return false
	}

	if vCode != oldVCode {
		return false
	}

	return true
}

// GetLoginVCode 获得校验码
func (s *Service) GetLoginVCode(phone int64) (int, error) {
	// 直接从redis里获取
	key := fmt.Sprintf(verifyCodeRedisKey, phone)
	vcode, err := redis.RedisClient.Get(key).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, errors.Wrap(err, "redis get login vcode err")
	}

	verifyCode, err := strconv.Atoi(vcode)
	if err != nil {
		return 0, errors.Wrap(err, "strconv err")
	}

	return verifyCode, nil
}
