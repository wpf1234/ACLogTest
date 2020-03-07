package models

import "time"

//域名流量日志
type DomainFlux struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	DstIp string `orm:"column(dst_ip)" json:"dst_ip"`//目标IP
	DstIpBin []byte `orm:"column(dst_ip_bin)" json:"dst_ip_bin"` //
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	LineNo int `orm:"column(line_no)" json:"line_no"`//线路编号
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
	UpFlux int `orm:"column(up_flux)" json:"up_flux"` //上行流量
	DownFlux int `orm:"column(down_flux)" json:"down_flux"` //下行流量
	Result string`orm:"column(result)"`//结果(一个 xml 文件)
}

type Flux struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	DstIp string `orm:"column(dst_ip)" json:"dst_ip"`//目标IP
	DstIpBin []byte `orm:"column(dst_ip_bin)" json:"dst_ip_bin"` //
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	LineNo int `orm:"column(line_no)" json:"line_no"`//线路编号
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
	UpFlux int `orm:"column(up_flux)" json:"up_flux"` //上行流量
	DownFlux int `orm:"column(down_flux)" json:"down_flux"` //下行流量
	Result string`orm:"column(result)"`//结果(一个 xml 文件)
}

//网络应用流量
type WebAppFlux struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	DstIp string `orm:"column(dst_ip)" json:"dst_ip"`//目标IP
	DstIpBin []byte `orm:"column(dst_ip_bin)" json:"dst_ip_bin"` //
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	LineNo int `orm:"column(line_no)" json:"line_no"`//线路编号
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
	UpFlux int `orm:"column(up_flux)" json:"up_flux"` //上行流量
	DownFlux int `orm:"column(down_flux)" json:"down_flux"` //下行流量
	Result string`orm:"column(result)"`//结果(一个 xml 文件)
}

//组别流量
type GroupFlux struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	LineNo int `orm:"column(line_no)" json:"line_no"`//线路编号
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
	UpFlux int `orm:"column(up_flux)" json:"up_flux"` //上行流量
	DownFlux int `orm:"column(down_flux)" json:"down_flux"` //下行流量
	Result string`orm:"column(result)"`//结果(一个 xml 文件)
}

type GroupFluxWebApp struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	LineNo int `orm:"column(line_no)" json:"line_no"`//线路编号
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
	UpFlux int `orm:"column(up_flux)" json:"up_flux"` //上行流量
	DownFlux int `orm:"column(down_flux)" json:"down_flux"` //下行流量
	Result string`orm:"column(result)"`//结果(一个 xml 文件)
}

type HttpTypeFlux struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`   //记录ID
	DevId int`orm:"column(dev_id)" json:"dev_id"`          //设备ID
	Group string `orm:"column(group)" json:"group"`		 //组名
	User string `orm:"column(user)" json:"user"`		 //用户名
	HostIp string `orm:"column(host_ip)"json:"host_ip"`    //源地址
	HostIpBin []byte `orm:"column(host_ip_bin)" json:"host_ip_bin"`//
	DstIp string `orm:"column(dst_ip)" json:"dst_ip"`//目标IP
	DstIpBin []byte `orm:"column(dst_ip_bin)" json:"dst_ip_bin"` //
	IpVersion string `orm:"column(ip_version)" json:"ip_version"`//IP版本
	HttpType string `orm:"column(http_type)" json:"http_type"`  //http请求类型
	Site string `orm:"column(site)" json:"site"`//位置
	TmType string `orm:"column(tm_type)" json:"tm_type"`//终端详情
	Server string `orm:"column(serv)" json:"server"`//服务
	App  string `orm:"column(app)" json:"app"`//网站分类
	LineNo int `orm:"column(line_no)" json:"line_no"`//线路编号
	Ntime int`orm:"column(ntime)" json:"ntime"`//
	RecordTime time.Time`orm:"column(record_time)"`//记录时间
	UpFlux int `orm:"column(up_flux)" json:"up_flux"` //上行流量
	DownFlux int `orm:"column(down_flux)" json:"down_flux"` //下行流量
	Result string`orm:"column(result)"`//结果(一个 xml 文件)
}