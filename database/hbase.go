package database

import (
	"context"
	"errors"
	"fmt"
	"gree/common/ghbase"
	"io"

	"github.com/astaxie/beego"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var client gohbase.Client

func ConnHBase(){
	hbaseUrl:=beego.AppConfig.String("hbase.host")
	user:=beego.AppConfig.String("hbase.user")
	option:=gohbase.EffectiveUser(user)
	ghbase.InitHbase(hbaseUrl,option)
	//client = gohbase.NewClient(hbaseUrl, option)

}

//向表中添加数据
func PutsByRowKey(table,rowKey string,values map[string]map[string][]byte) (err error){
	//values := map[string]map[string][]byte{colFam: colVal}     //values 结构:{列族:map[列名]值}
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values)
	if err != nil {
		beego.Error("hrpc.NewPutStr: %s", err.Error())
	}
	_, err = client.Put(putRequest)
	if err != nil {
		beego.Error("hbase clients: %s", err.Error())
	}
	return
}

//获取数据
func Gets(table, rowKey string) (*hrpc.Result, error) {

	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowKey)
	if err != nil {
		beego.Error("hrpc.NewGetStr: %s", err.Error())
	}
	res, err := client.Get(getRequest)
	if err != nil {
		beego.Error("hbase clients: %s", err.Error())
	}
	defer func() {
		if errs := recover(); errs != nil {
			switch fmt.Sprintf("%v", errs) {
			case "runtime error: index out of range":
				err = errors.New("no such rowKey or qualifier exception")
			case "runtime error: invalid memory address or nil pointer dereference":
				err = errors.New("no such colFamily exception")
			default:
				err = fmt.Errorf("%v", errs)
			}
		}
	}()
	return res, nil
}

//查看数据
func Scan(table, startRow, stopRow string) (rsp []*hrpc.Result, err error) {
	var (
		scanRequest *hrpc.Scan
		res         *hrpc.Result
	)
	scanRequest, err = hrpc.NewScanRangeStr(context.Background(), table, startRow, stopRow)
	if err != nil {
		beego.Error("hrpc.NewScanStr: %s", err.Error())
	}
	scanner := client.Scan(scanRequest)
	for {
		res, err = scanner.Next()
		if err == io.EOF || res == nil {
			break
		}
		if err != nil {
			beego.Error("hrpc.Scan: %s", err.Error())
		}
		rsp = append(rsp, res)
	}
	return rsp, err
}

//使用过滤器筛选

//查看rowkey是否存在
func IsExistRowkey(table, rowKey string) bool {
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowKey)
	if err != nil {
		beego.Error("hrpc.NewGetStr: %s", err.Error())
	}
	res, err := client.Get(getRequest)
	if err != nil {
		beego.Error("get from hbase: %s", err.Error())
	}
	if len(res.Cells) > 0 {
		return true
	} else {
		return false
	}
}

//删除数据
func DeleteByRowkey(table, rowkey string, value map[string]map[string][]byte) (err error) {
	deleteRequest, err := hrpc.NewDelStr(context.Background(), table, rowkey, value)
	if err != nil {
		beego.Error("hrpc.NewDelStrRef: %s", err.Error())
	}
	//fmt.Println("deleteRequest:", deleteRequest)
	res, err := client.Delete(deleteRequest)

	fmt.Println(res)
	if err != nil {
		beego.Error("hrpc.Scan: %s", err.Error())
	}
	return
}