package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	common "github.com/ljp-lachouchou/go-learn-micro-common"
	"github.com/ljp-lachouchou/go-micro-learn-payment/domain/repository"
	service2 "github.com/ljp-lachouchou/go-micro-learn-payment/domain/service"
	"github.com/ljp-lachouchou/go-micro-learn-payment/handler"
	payment "github.com/ljp-lachouchou/go-micro-learn-payment/proto/payment"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	prometheus2 "github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var QPS = 1000

func main() {
	common.Info("程序启动") // 触发文件创建
	//配置中心
	consulConfig, err := common.GetConsulConfig("localhost",
		8500, "/micro/config")
	if err != nil {
		common.Error(err)
	}
	//注册中心
	cl := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//jaeger链路追踪
	t, io, err := common.NewTracer("go.micro.service.payment", "localhost:6831")
	if err != nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	//mysql设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		common.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)
	//repository.NewPaymentRepository(db).InitTable()
	//监控
	common.PrometheusBoot(9089)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8089"),
		micro.Registry(cl),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(prometheus2.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()
	paymentDataService := service2.NewPaymentDataService(repository.NewPaymentRepository(db))

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), &handler.Payment{PaymentDataService: paymentDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
