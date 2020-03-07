package models

import (
	"github.com/astaxie/beego"
	"time"
)

//行为日志
type UserAction struct {
	RecordId int `orm:"column(record_id)" json:"record_id"` //记录ID
	//DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	//Group string `orm:"column(group)" json:"group"`		 //组名
	User   string `orm:"column(user)"`    //用户名
	HostIp string `orm:"column(host_ip)"` //源地址
	DstIp  string `orm:"column(dst_ip)"`  //目标IP
	//IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	//Site string `orm:"column(site)" json:"site"`//位置
	//TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server   string `orm:"column(serv)"`      //服务
	App      string `orm:"column(app)"`       //网站分类
	SrcPort  int    `orm:"column(src_port)"`  //源端口
	ServPort int    `orm:"column(serv_port)"` //服务端口
	//NetAction string`orm:"column(net_action)" json:"net_action"`//访问控制
	RecordTime string `orm:"column(record_time)" json:"record_time"` //记录时间
	Result     string `orm:"column(result)"`                         //结果(一个 xml 文件)
}

type UserNetAction struct {
	RecordId string `json:"record_id"` //记录ID
	User     string `json:"user"`      //ip 对应的用户
	Server   string `json:"server"`    //服务类型(邮件/访问网站/网络存储/移动终端/游戏)
	App      string `json:"app"`       //应用信息(xx上传/微信/QQ/网盘/网站名/移动端app)
	HostIp   string `json:"host_ip"`   //源IP
	//Mac        string `json:"mac"`         //源MAC
	SrcPort    string `json:"src_port"`    //源端口号
	DstIp      string `json:"dst_ip"`      //目标IP
	ServPort   string `json:"serv_port"`   //目标端口
	RecordTime string `json:"record_time"` //记录时间
	Resource   string `json:"resource"`    //上传内容的相关链接
}

func (u *UserAction) TableName() string {
	t := time.Now().AddDate(0, 0, 0).Format("20060102")
	table := t + beego.AppConfig.String("actTB")
	return table
}
