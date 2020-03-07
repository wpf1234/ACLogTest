package kafka

import (
	"ACTest/models"
	"ACTest/util"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gree/common/ghbase"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

//版本一: data []map[string]string
func InitKafkaAndSend(topic, key string, data []map[string]string) {
	hosts := beego.AppConfig.String("kafka_broker")
	//kafkaThreadNum,_:=beego.AppConfig.Int("kafka_threadNum")
	Address := strings.Split(hosts, ",")
	if len(Address) == 0 {
		beego.Error("kafka主机地址配置错误!")
		return
	}
	//设置配置
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_8_2_0

	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(Address, config)
	if e != nil {
		beego.Error("创建生产者失败: ", e)
		return
	}
	defer producer.AsyncClose()

	//发送的消息,主题,key
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Timestamp: time.Now(),
	}

	for _, v := range data {
		//fmt.Println("Message is: ", v)
		//将字符串转化为字节数组
		myVal, _ := json.Marshal(v)
		msg.Value = sarama.StringEncoder(myVal)
		//msg.Value = sarama.StringEncoder(v)

		//使用通道发送
		producer.Input() <- msg

		//循环判断哪个通道发送过来数据.
		select {
		//suc :=
		case <-producer.Successes():
			//fmt.Println("offset: ", suc.Offset, "partitions: ", suc.Partition)
			//fmt.Println("Send to Kafka success!!!")
		case fail := <-producer.Errors():
			beego.Error("err: ", fail.Err)

		}
	}

}

//方法二:数据结构为 string 类型  (data []string)
func InitKafkaAndSend2(topic, key string, data []string) {
	hosts := beego.AppConfig.String("kafka_broker")
	//kafkaThreadNum,_:=beego.AppConfig.Int("kafka_threadNum")
	Address := strings.Split(hosts, ",")
	if len(Address) == 0 {
		beego.Error("kafka主机地址配置错误!")
		return
	}
	//设置配置
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_8_2_0

	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(Address, config)
	if e != nil {
		beego.Error("创建生产者失败: ", e)
		return
	}
	defer producer.AsyncClose()

	//发送的消息,主题,key
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		Timestamp: time.Now(),
	}

	for _, v := range data {
		//fmt.Println("Message is: ", v)
		//将字符串转化为字节数组
		//myVal,_:=json.Marshal(v)
		//msg.Value = sarama.StringEncoder(myVal)
		msg.Value = sarama.StringEncoder(v)

		//使用通道发送
		producer.Input() <- msg

		//循环判断哪个通道发送过来数据.
		select {
		//suc :=
		case <-producer.Successes():
			//fmt.Println("offset: ", suc.Offset, "partitions: ", suc.Partition)
			//fmt.Println("Send to Kafka success!!!")
		case fail := <-producer.Errors():
			beego.Error("err: ", fail.Err)

		}
	}

}

func KafkaConsumer(topic string) {
	//获取 HBase 表的名称
	hTable := beego.AppConfig.String("hbase.table")
	//获取所有 broker 的集合
	hosts := beego.AppConfig.String("kafka_broker")
	//kafkaThreadNum,_:=beego.AppConfig.Int("kafka_threadNum")
	Address := strings.Split(hosts, ",")
	if len(Address) == 0 {
		beego.Error("kafka主机地址配置错误!")
		return
	}
	//配置
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_8_2_0
	//新建一个消费者
	consumer, e := sarama.NewConsumer(Address, config)
	if e != nil {
		beego.Error("Error get consumer: ", e)
		return
	}
	defer func() error {
		if err := consumer.Close(); err != nil {
			beego.Error(err.Error())
			return err
		}
		return nil
	}()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		beego.Error("Error get partition consumer", err)
		return
	}

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	//定义一个计数器
	count := 0

	//循环等待接受消息.
ConsumerLoop:
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			//fmt.Println(" timestrap: ", msg.Timestamp.Format("2006-01-02 15:04"), " value: ", string(msg.Value))
			var colVal = make(map[string][]byte)
			colVal["ACInfo"] = msg.Value
			colFam := topic
			tm := strconv.FormatInt(time.Now().Unix(), 10)
			rowKey := string(msg.Key) + "_" + tm

			_, err = ghbase.Puts(hTable, rowKey, colFam, colVal)
			if err != nil {
				beego.Error("Puts into HBase failed: ", err)
				return
			}
			count++
			//fmt.Println("Count: ", count, " offset: ", msg.Offset, " Put into HBase success!")
		case err := <-partitionConsumer.Errors():
			beego.Error("Get info from kafka failed", err.Err)
			return
		case <-signals:
			break ConsumerLoop
		}
	}

}

//将spark处理完的数据从 kafka 中读出，消费进入 MySQL 数据库
func GetACResFromKafka() {
	//获取计算结果的topic
	topic := beego.AppConfig.String("kafka_resTopic")
	//获取 broker 的集合
	hosts := beego.AppConfig.String("kafka_broker")
	Address := strings.Split(hosts, ",")
	if len(Address) == 0 {
		beego.Error("Kafka集群地址配置错误!")
		return
	}

	//配置 kafka
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_8_2_0
	//新建一个消费者
	consumer, e := sarama.NewConsumer(Address, config)
	if e != nil {
		beego.Error("Error get consumer: ", e)
		return
	}
	defer func() error {
		if err := consumer.Close(); err != nil {
			beego.Error(err.Error())
			return err
		}
		return nil
	}()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		beego.Error("Error get partition consumer", err)
		return
	}

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	//循环等待接收消息
ConsumerLoop:
	for {
		select {
		//接收消息通道和错误消息的内容
		case msg := <-partitionConsumer.Messages():

			//fmt.Println("Value: ", string(msg.Value), " timestamp: ", msg.Timestamp.Format("2006-01-02 15:04:05"))
			//util.ColorPrintln(string(msg.Value), 12)
			var acRes models.ACResult
			var check models.ACResult
			o := orm.NewOrm()
			//注册多个数据库后，default 别名的数据库不需要使用Using
			//其他数据库在使用前需要使用Using("别名"),才能正确操作数据库
			o.Using("db1")
			//dr := o.Driver()
			//fmt.Println("1111111111111: ", dr.Name())
			data := string(msg.Value)
			datas := strings.Split(data, "\t")
			cnt, _ := strconv.Atoi(datas[4])
			t := time.Now().Format("2006-01-02")

			//s1 := strings.Split(datas[5], "[")
			//s2 := strings.Split(s1[1], "]")
			//s3 := s2[0]
			//s4 := strings.Split(s3, ",")
			//s4 = util.Distinct(s4)

			//a1 := strings.Split(datas[6], "[")
			//a2 := strings.Split(a1[1], "]")
			//a3 := a2[0]
			//a4 := strings.Split(a3, ",")
			//a4 = util.Distinct(a4)

			r1 := strings.Split(datas[7], "[")
			r2 := strings.Split(r1[1], "]")
			r3 := r2[0]
			//r4 := strings.Split(r3, ",")
			//r4 = util.Distinct(r4)

			rc1 := strings.Split(datas[8], "[")
			rc2 := strings.Split(rc1[1], "]")
			rc3 := rc2[0]
			rc4 := strings.Split(rc3, ",")
			rc4 = util.Distinct(rc4)
			//fmt.Println("888888888: ", rc3, rc4[len(rc4)-1])

			acRes.User = datas[0]
			acRes.Ip = datas[1]
			acRes.Mac = datas[2]
			acRes.DstIp = datas[3]
			acRes.Cnt = cnt
			acRes.Server = datas[5]
			acRes.App = datas[6]
			acRes.Resource = r3
			acRes.RecordTime = rc4[len(rc4)-1]
			//acRes.Resource = strings.Join(r4, ",")
			//acRes.RecordTime = strings.Join(rc4, ",")
			acRes.MTime = t
			// time.Now() 得到的当前时间的时区跟电脑的当前时区一样
			acRes.CTime = time.Now()
			//acRes.MTime = time.Now().AddDate(0, 0, 0).Format("2006-01")
			//acRes.CTime = time.Now().AddDate(0, 0, 0)

			create, id, err := o.ReadOrCreate(&acRes, "User", "DstIp", "MTime")
			if err != nil {
				beego.Error("Create failed: ", err)
				return
			}
			if create {
				//fmt.Println("*****************INSERT*******************")
				beego.Info("New insert an object and id: ", id, acRes.CTime)
				acRes.Id = id
			} else {
				//fmt.Println("=================UPDATE===================")
				//beego.Warn("Get an object ,id is: ", id)
				check.Id = id
				err = o.Read(&check, "Id")
				if err != nil {
					beego.Error("Read failed: ", err)
					return
				}

				acRes.Cnt = cnt + check.Cnt
				//acRes.RecordTime = rc3
				acRes.Server = datas[5]
				acRes.App = datas[6]
				acRes.Resource = r3 + "," + check.Resource
				acRes.RecordTime = rc4[len(rc4)-1]
				acRes.CTime = time.Now()

				if check.Cnt != acRes.Cnt && check.RecordTime != acRes.RecordTime {
					_, err = o.Update(&acRes, "Cnt", "Server", "App", "Resource", "RecordTime", "cTime")
					if err != nil {
						beego.Error("Update AC result failed: ", err)
						return
					}

					beego.Info("Update success,count is: ", acRes.Cnt, acRes.CTime)
				}
			}

		case err := <-partitionConsumer.Errors():
			beego.Error("从kafka中获取消息失败: ", err.Err)
			return
		case <-signals:
			break ConsumerLoop
		}
	}
}
