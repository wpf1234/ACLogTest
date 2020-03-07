package main

import (
	"ACTest/database"
	_ "ACTest/routers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"log/ACTest.log","maxdays":1}`)
	beego.SetLogFuncCall(true)
	//beego.BeeLogger.DelLogger("console") //删除console日志输出
	//将日志输出到ES 中
	//ES 配置 level
	//logs.SetLogger(logs.AdapterEs, `{"dsn":"http://loclhost:9200/","level":1}`)

}

func main() {
	database.InitMysql()
	//database.InitACDB()
	database.InitNetLog()
	database.ConnHBase()
	//go kafka.KafkaConsumer(beego.AppConfig.String("kafka_topic"))
	//go util.StartTimer(database.CopyDataToHBase)
	//go kafka.GetACResFromKafka()

	//orm.Debug = true

	beego.Run()
}
