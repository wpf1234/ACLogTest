package database

import (
	"ACTest/kafka"
	"ACTest/models"
	"ACTest/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.DefaultString("dbpwd", "123456")
	dbname := beego.AppConfig.String("dbname")

	maxIdle := 200
	maxConn := 200
	maxLifetime := 100 * time.Second

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("Failed to register data driver: ", err)
		return
	}

	err = orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn, int(maxLifetime.Nanoseconds()/1e6))
	if err != nil {
		beego.Error("Failed to register database: ", err)
		return
	}

	beego.Info("Connect database 1 success!")
}

func InitACDB() {
	/**
	获取存储计算结果的数据库相关信息
	*/
	acDbhost := beego.AppConfig.String("AC_host")
	acDbport := beego.AppConfig.String("AC_port")
	acDbuser := beego.AppConfig.String("AC_user")
	acDbpwd := beego.AppConfig.String("AC_pwd")
	acDbname := beego.AppConfig.String("AC_db")
	acDSN := acDbuser + ":" + acDbpwd + "@tcp(" + acDbhost + ":" + acDbport + ")/" + acDbname + "?charset=utf8"

	maxIdle := 200
	maxConn := 200
	maxLifetime := 100 * time.Second

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("Failed to register data driver: ", err)
		return
	}
	orm.RegisterModel(new(models.ACResult))

	//注册数据库
	err = orm.RegisterDataBase("db1", "mysql", acDSN, maxIdle, maxConn, int(maxLifetime.Nanoseconds()/1e6))
	if err != nil {
		beego.Error("Failed to register database: ", err)
		return
	}

	beego.Info("Connect database 2 success!")
}

func InitNetLog() {
	/**
	获取存储计算结果的数据库相关信息
	*/
	aclogHost := beego.AppConfig.String("log_host")
	aclogPort := beego.AppConfig.String("log_port")
	aclogUser := beego.AppConfig.String("log_user")
	aclogPwd := beego.AppConfig.String("log_pwd")
	aclogDb := beego.AppConfig.String("log_db")
	acDSN := aclogUser + ":" + aclogPwd + "@tcp(" + aclogHost + ":" + aclogPort + ")/" + aclogDb + "?charset=utf8"

	maxIdle := 200
	maxConn := 200
	maxLifetime := 100 * time.Second

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("Failed to register data driver: ", err)
		return
	}
	orm.RegisterModel(new(models.ACResult))

	//注册数据库
	err = orm.RegisterDataBase("db2", "mysql", acDSN, maxIdle, maxConn, int(maxLifetime.Nanoseconds()/1e6))
	if err != nil {
		beego.Error("Failed to register database: ", err)
		return
	}

	beego.Info("Connect database 2 success!")
}

//定时将数据生产到 kafka
func CopyDataToHBase() {
	for {
		var IP []string

		//now := time.Now()

		date := time.Now().AddDate(0, 0, -3).Format("20060102")
		table := date + beego.AppConfig.String("actTB")

		topic := beego.AppConfig.String("kafka_topic") //colFam
		//htable:=beego.AppConfig.String("hbase.table")

		o := orm.NewOrm()
		_, err := o.Raw(fmt.Sprintf("SELECT host_ip FROM %s", table)).QueryRows(&IP)
		if err != nil {
			beego.Error("Get user name failed: ", err)
			return
		}

		fmt.Println("length1: ", len(IP))
		IP = util.Distinct(IP)
		fmt.Println("length2: ", len(IP))
		for n, _ := range IP {
			var userNetAC []string
			//var userNetAC =make([]map[string]string,0)
			var userAC []models.UserAction
			_, err := o.Raw(fmt.Sprintf("SELECT record_id,`user`,host_ip,dst_ip,serv,app,src_port,serv_port,record_time,result FROM %s WHERE host_ip='%s'", table, IP[n])).QueryRows(&userAC)
			if err != nil {
				beego.Error("Get such info error: ", err)
				return
			}

			//方法一:传入 []map[string]string
			//userNetAC = util.HandleAC(userAC)

			//方法二:传入 []string
			userNetAC = util.HandleACTest(userAC)

			//数据存入 kafka
			intIP := strconv.FormatInt(util.InetAtoN(IP[n]), 10)
			intIP = util.ReverseString(intIP)
			//tm:=strconv.FormatInt(time.Now().Unix(),10)
			key := intIP //rowKey

			//方法一
			//kafka.InitKafkaAndSend(topic,key,userNetAC)

			//方法二
			kafka.InitKafkaAndSend2(topic, key, userNetAC)

			//数据存入 HBase,处理速度较慢,优化读入 HBase
			//intIP:=strconv.FormatInt(util.InetAtoN(IP[n]),10)
			//intIP=util.ReverseString(intIP)
			//tm:=strconv.FormatInt(time.Now().Unix(),10)
			//rowKey:=intIP+"_"+tm
			////fmt.Println(rowKey)
			////将 userNetAC 拆分为 column 和 value
			//
			//for cnt,v:=range userNetAC{
			//	fmt.Println("count= ",cnt+1)
			//	err=PutsByRowKey(htable,rowKey,v)
			//		if err!=nil{
			//			beego.Error("Copy data to HBase failed: ",err)
			//			return
			//		}
			//}
		}

		//fmt.Println("Copy data to Kafka successful!!!")

		// 计算下一个零点
		//next := now.Add(time.Hour * 24)
		//next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		//t := time.NewTimer(next.Sub(now))
		//<-t.C
	}
}
