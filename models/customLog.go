package models

import "time"

//用户日志
type CustomLog struct {
	RecordId int `orm:"column(record_id)" json:"record_id"`
	DevId int `orm:"column(dev_id)" json:"dev_id"`
	Identify string `orm:"column(identify)" json:"identify"`
	RecordTime time.Time `orm:"column(record_time)" json:"record_time"`
	Result string `orm:"column(result)" json:"result"`
}
