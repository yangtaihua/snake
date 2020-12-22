// 单元测试工具

package testutil

import (
	"github.com/jinzhu/gorm"

	"github.com/1024casts/snake/internal/model"
	"github.com/1024casts/snake/internal/service"
	"github.com/1024casts/snake/pkg/conf"
	"github.com/1024casts/snake/pkg/redis"
)

// App 结构体，主要是为了方便实例一个app
type App struct {
	DB *gorm.DB
}

// Initialize 初始化
func (app *App) Initialize() {
	cfg := "../../../../conf/config.sample.yaml"
	if err := conf.Init(cfg); err != nil {
		panic(err)
	}

	// init log
	conf.InitLog(conf.Conf)

	// init db
	model.Init(conf.Conf)
	app.DB = model.DB

	redis.InitTestRedis()

	// init service
	svc := service.New(conf.Conf)

	// set global service
	service.Svc = svc
}
