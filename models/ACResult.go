package models

import "time"

type ACResult struct {
	Id         int64     `orm:"pk;auto;column(id)"`
	User       string    `orm:"column(user)"`
	Ip         string    `orm:"column(ip)"`
	Mac        string    `orm:"column(mac)"`
	DstIp      string    `orm:"column(dst_ip)"`
	Cnt        int       `orm:"column(cnt)"`
	Server     string    `orm:"column(server)"`
	App        string    `orm:"column(app)"`
	Resource   string    `orm:"column(resource)"`
	RecordTime string    `orm:"column(r_time)"`
	MTime      string    `orm:"column(m_time)"`
	CTime      time.Time `orm:"column(c_time)"`
}

func (a *ACResult) TableName() string {
	return "user_action"
}

type Action struct {
	Id      int    `json:"id"`
	Time    string `json:"time"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Mac     string `json:"mac"`
	DstIp   string `json:"dstIp"`
	DstPort string `json:"dstPort"`
	//Result  string `json:"result"`
}

type Desc struct {
	Time     []string `json:"time"`
	Ip       string   `json:"ip"`
	Port     string   `json:"port"`
	Mac      string   `json:"mac"`
	DstIp    string   `json:"dstIp"`
	Data     string   `json:"data"`     //具体内容
	Resource string   `json:"resource"` //传输的资源
}

type Response struct {
	List  []Action `json:"list"`
	Total int      `json:"total"`
}
