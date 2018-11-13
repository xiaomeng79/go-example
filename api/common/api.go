package common

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/micro/go-web"
	"github.com/xiaomeng79/go-example/cinit"
	"github.com/xiaomeng79/go-example/internal/metrics"
	"github.com/xiaomeng79/go-log"
	"net/http"
	"time"
)

//定义services名称
const SN = "com.example.api.common"

//运行
func Run() {
	//初始化
	cinit.InitOption(SN, "trace")
	//新建服务
	serviceName := cinit.Config.Service.Name
	serviceVersion := cinit.Config.Service.Version

	service := web.NewService(
		web.Name(serviceName),
		web.Version(serviceVersion),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	)
	log.Info("创建服务:名称:" + serviceName + ",版本:" + serviceVersion)
	// 定义Service动作操作
	service.Init()

	web.AfterStop(func() error {
		log.Info("停止服务")
		//停止配置
		cinit.Close()
		return nil
	})

	//连接服务
	NewUserClient("com.example.srv.user")

	//注册路由
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Origin", "Authorization", "Accept", "Client-Security-Token", "Accept-Encoding"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(Opentracing)

	//metrics
	// Metrics
	if cinit.Config.Metrics.Enable == "yes" {

		/* Pull模式
		e.Use(prometheus.MetricsFunc(
			prometheus.Namespace("common_api"),
		))
		*/

		// Push模式
		m := metrics.NewMetrics()
		e.Use(MetricsFunc(m))
		m.MemStats()
		// InfluxDB
		m.InfluxDBWithTags(
			time.Duration(cinit.Config.Metrics.Duration)*time.Second,
			cinit.Config.Metrics.Url,
			cinit.Config.Metrics.Database,
			cinit.Config.Metrics.UserName,
			cinit.Config.Metrics.Password,
			map[string]string{"service": serviceName},
		)

		// Graphite
		//addr, _ := net.ResolveTCPAddr("tcp", Conf.Metrics.Address)
		//m.Graphite(Conf.Metrics.FreqSec*time.Second, "echo-web.node."+hostname, addr)

	}

	//加验证JWT路由组，版本v1
	g1 := e.Group("/common/v1", JWT)
	g1.POST("/userinfo", userinfo)

	//不验证JWT路由组
	g2 := e.Group("/common/v1")
	g2.POST("/login", login)

	//check
	check := e.Group("/common/check")
	check.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	service.Handle("/", e)

	//启动service
	if err := service.Run(); err != nil {
		log.Fatal("启动服务失败" + err.Error())
	}
}
