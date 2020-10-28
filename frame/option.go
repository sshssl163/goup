package frame

import "github.com/gin-gonic/gin"

// An Option configures
type Option interface {
	apply(*bootstarpServerConfig)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*bootstarpServerConfig)

func (f optionFunc) apply(log *bootstarpServerConfig) {
	f(log)
}

// BeforeInit 在框架初始化之前运行
func BeforeInit(f func()) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.beforeInit = f
	})
}

// BeforeServerRun 在web服务启动之前运行
func BeforeServerRun(f func()) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.beforeServerRun = f
	})
}

// CustomRouter 自定义路由，用于突破框架的 json api 的局限性
func CustomRouter(f func(r *gin.Engine)) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.customRouter = f
	})
}

// Version 自定义 /version path 返回内容
func Version(f func(c *gin.Context)) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.versionHandler = f
	})
}

// ReportApi 上报接口到文档中心
func ReportApi(addr string) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.reportApiDocAddr = addr
	})
}

// Middlewares 添加自定义的 middleware
func Middlewares(list []gin.HandlerFunc) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.middlewareList = list
	})
}

// 配置中心本地缓存默认父目录使能
func EnableconfigCenterDefalutLocalCacheDir(enable bool) Option {
	return optionFunc(func(cfg *bootstarpServerConfig) {
		cfg.defaultLocalCachePathEnable = enable
	})
}
