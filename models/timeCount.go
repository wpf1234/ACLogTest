package models

import "time"

//域名时长
type DomainTimeCount struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	Second int `orm:"column(second)" json:"second"`
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
}


type TimeCount struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	Second int `orm:"column(second)" json:"second"`
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
}

type TimeCountWebApp struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	Second int `orm:"column(second)" json:"second"`
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
}