package cinit

import (
	"github.com/jinzhu/configor"
	"log"
)

const (
	//上下文
	TRACE_CONTEXT = "trace_ctx"     //trace
	REQ_PARAM     = "req_param"     //请求参数绑定
	JWT_NAME      = "Authorization" //JWT请求头名称
	JWT_MSG       = "JWT-MSG"       //JWT自定义的消息
)

//公共配置
var Config = struct {
	Service struct {
		Name     string `default:"com.example.srv.test"` //服务名称
		Version  string `default:"v1.0"`                 //服务版本号
		RateTime int    `default:"1024"`                 //限制请求
	}
	//tracing
	Trace struct { //链路跟踪
		Address string `default:"127.0.0.1:6831"`
	}
	//log config
	Log struct { //日志
		Path         string `default:"tmp"` //日志保存路径
		IsStdOut     string `default:"yes"` //是否输出日志到标准输出 yes:输出 no:不输出
		MaxAge       int    `default:"7"`   //日志最大的保存时间，单位天
		RotationTime int    `default:"1"`   //日志分割的时间，单位天
		MaxSize      int    `default:"100"` //日志分割的尺寸，单位MB
	}
	//mysql config
	Mysql struct {
		DbName   string `default:"test"`      //数据库名称
		Addr     string `default:"127.0.0.1"` //地址
		User     string `default:"root"`
		Password string `default:"root"`
		Port     int    `default:"3306"` //required:"true" env:"DB_PROT"
		IdleConn int    `default:"5"`    //空闲连接
		MaxConn  int    `default:"20"`   //最大连接
	}
	//mongo config
	Mongo struct {
		Hosts     string `default:"127.0.0.1:27017"` //数据库地址，可以多个，用逗号分割
		DbName    string `default:"test"`            //数据库名称
		User      string `default:"root"`
		Password  string `default:"root"`
		PoolLimit int    `default:"4096"` //连接池限制
	}
}{}

//初始化配置文件
//配置加载顺序1.是否设置了变量conf，设置了第一个加载，如果文件不存在，加载默认配置文件
//如果设置了环境变量 CONFIGOR_ENV = test等，那么加载config_test.yml的配置文件
//最后加载环境变量,是否设置环境变量前缀,如果设置了CONFIGOR_ENV_PREFIX=WEB,设置环境变量为WEB_DB_NAME=root,否则为DB_NAME=root
func configInit(sn string) {

	//config := flag.String("conf", "conf/config.yml", "you configuer file")
	//flag.Parse()
	//err := configor.Load(&Config, *config)
	configor.Load(&Config, "config.yml")
	Config.Service.Name = sn //使用传入的名称
	log.Printf("config: %+v\n", Config)
}

//保存需要关闭的选项
var closeArgs []string

//初始化选项
//log:日志(必须) trace:链路跟踪 mysql:mysql数据库 mongo:MongoDB
func InitOption(sn string, args ...string) {
	//保存需要关闭的参数
	closeArgs = args
	//1.初始化配置参数
	configInit(sn)
	//2.初始化日志
	logInit()
	//3.其他服务
	for _, o := range args {
		switch o {
		case "trace":
			traceInit()
		case "mysql":
			//原始
			mysqlInit()
		case "mongo":
			//todo
		}
	}
}

//关闭打开的服务
func Close() {
	for _, o := range closeArgs {
		switch o {
		case "trace":
			//关闭链路跟踪
			tracerClose()
		case "mysql":
			//关闭mysql
			//原始
			mysqlClose()
		case "mongo":
			//TODO
		}
	}

}
