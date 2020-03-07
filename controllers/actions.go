package controllers

import (
	"ACTest/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
)

type GetActionsController struct {
	baseController
}

//1.网页、2.邮件、3.网盘、4.微信、5.QQ、6.移动端
var TypeIndex = map[int]string{
	1: "邮件",
	2: "网站",
	3: "网盘",
	4: "qq",
	5: "微信",
	6: "移动端",
}

//获取用户某段时间内上网的行为
// 需要显示访问网站的排名,进行的操作,资源文件

//func (this *GetActionsController) GetActions() {
//	//从前端获取请求，然后根据type 去访问相应的信息
//	//注意需要分页显示，注意时间
//	//默认的话 按照时间的升序，访问量的降序显示
//	var request models.RequestBody
//	var response models.Response
//	var acRes []models.Action
//	var id []int
//	var mTimes []string
//	var rTimes []string
//	var ip []string
//	var mac []string
//	var dstIp []string
//	var res string
//
//	total := 0
//	o := orm.NewOrm()
//
//	o.Using("db1")
//	body := this.Ctx.Input.RequestBody
//
//	err := json.Unmarshal(body, &request)
//	if err != nil {
//		beego.Error("数据解析失败!", err)
//		this.Data["json"] = models.Result{Code: 500, Data: "", Message: "数据解析失败!"}
//		this.ServeJSON()
//		return
//	}
//
//	t := time.Now().Format("2006-01-02")
//	//没有发送起始时间和结束时间，使用默认方法查询
//	//前端传输的是 2019-05-27 10:00:00
//	//go 默认为 UTC 时间，需要转换为 CTS 时间
//	if request.StartTime == "" && request.EndTime == "" {
//		//fmt.Println(typeIndex[request.Type])
//
//		err = o.Raw("select count(*) from user_action  "+
//			"where (`server` like ? or app like ?) "+
//			"and m_time = ? and ip like ? and ip like ? "+
//			"and mac like ? and dst_ip like ? ",
//			"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
//			t, "%"+request.Ip+"%", "%"+request.Port+"%",
//			"%"+request.Mac+"%", "%"+request.DstIp+"%").QueryRow(&total)
//		if err != nil {
//			beego.Error("获取总数错误: ", err)
//			this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取总数失败!"}
//			this.ServeJSON()
//			return
//		}
//		sql := fmt.Sprintf("select id,m_time,r_time,ip,mac,dst_ip from user_action "+
//			"where (`server` like '%s' or app like '%s')"+
//			"and m_time like '%s' and ip like '%s' and ip like '%s' "+
//			"and mac like '%s' "+
//			"and dst_ip like '%s' "+
//			"order by cnt,r_time desc limit %d,%d",
//			"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
//			"%"+t+"%", "%"+request.Ip+"%", "%"+request.Port+"%", "%"+request.Mac+"%",
//			"%"+request.DstIp+"%", request.Page*request.Size, request.Size)
//		_, err = o.Raw(sql).QueryRows(&id, &mTimes, &rTimes, &ip, &mac, &dstIp)
//		if err != nil {
//			beego.Error("获取信息失败: ", err)
//			this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取数据失败!"}
//			this.ServeJSON()
//			return
//		}
//
//		for n, _ := range mTimes {
//			var ac models.Action
//			ac.Time = rTimes[n]
//			ac.Id = id[n]
//			ac.Ip = strings.Split(ip[n], ":")[0]
//			ac.Port = strings.Split(ip[n], ":")[1]
//			ac.Mac = mac[n]
//			ac.DstIp = strings.Split(dstIp[n], ":")[0]
//			o.Raw("select resource from user_action where id=?", ac.Id).QueryRow(&res)
//			acRes = append(acRes, ac)
//		}
//
//		response.List = acRes
//		response.Total = total
//
//		this.Data["json"] = models.Result{Code: 200, Data: response, Message: "获取数据成功!"}
//
//	} else {
//		//var start, end time.Time
//		//s := strings.Split(request.StartTime, "(")
//		//
//		//loc, _ := time.LoadLocation("Asia/Chongqing")
//		//
//		//start, _ = time.ParseInLocation("Mon Jan 02 2006 15:04:05 GMT+0800 ", s[0], loc)
//		//
//		//if request.EndTime != "" {
//		//	e := strings.Split(request.EndTime, "(")
//		//	end, _ = time.ParseInLocation("Mon Jan 02 2006 15:04:05 GMT+0800 ", e[0], loc)
//		//} else {
//		//	end = time.Now()
//		//}
//		//
//		//fmt.Println(start, end)
//
//		err = o.Raw("select count(*) from user_action  "+
//			"where (`server` like ? or app like ?) "+
//			"and r_time between ? and ? "+
//			"and ip like ? and ip like ? "+
//			"and mac like ? and dst_ip like ? "+
//			"order by cnt desc",
//			"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
//			request.StartTime, request.EndTime,
//			"%"+request.Ip+"%", "%"+request.Port+"%",
//			"%"+request.Mac+"%", "%"+request.DstIp+"%").QueryRow(&total)
//		if err != nil {
//			beego.Error("获取总数错误: ", err)
//			this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取总数失败!"}
//			this.ServeJSON()
//			return
//		}
//
//		sql := fmt.Sprintf("select id,m_time,r_time,ip,mac,dst_ip from user_action "+
//			"where (`server` like '%s' or app like '%s') "+
//			"and r_time between '%s' and '%s' "+
//			"and ip like '%s' and ip like '%s' "+
//			"and mac like '%s' "+
//			"and dst_ip like '%s' "+
//			"order by cnt,r_time desc limit %d,%d",
//			"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
//			request.StartTime, request.EndTime,
//			"%"+request.Ip+"%", "%"+request.Port+"%",
//			"%"+request.Mac+"%",
//			"%"+request.DstIp+"%", request.Page*request.Size, request.Size)
//
//		_, err = o.Raw(sql).QueryRows(&id, &mTimes, &rTimes, &ip, &mac, &dstIp)
//		if err != nil {
//			beego.Error("获取信息失败: ", err)
//			this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取数据失败!"}
//			this.ServeJSON()
//			return
//		}
//		for n, _ := range mTimes {
//			var ac models.Action
//			ac.Time = rTimes[n]
//			ac.Id = id[n]
//			ac.Ip = strings.Split(ip[n], ":")[0]
//			ac.Port = strings.Split(ip[n], ":")[1]
//			ac.Mac = mac[n]
//			ac.DstIp = strings.Split(dstIp[n], ":")[0]
//			o.Raw("select resource from user_action where id=?", ac.Id).QueryRow(&res)
//
//			acRes = append(acRes, ac)
//
//		}
//		response.List = acRes
//		response.Total = total
//
//		this.Data["json"] = models.Result{Code: 200, Data: response, Message: "获取数据成功!"}
//	}
//
//	//beego.Info("Res: ", res)
//
//	this.ServeJSON()
//}
func (this *GetActionsController) GetActions() {
	//从前端获取请求，然后根据type 去访问相应的信息
	//注意需要分页显示，注意时间
	//默认的话 按照时间的升序，访问量的降序显示
	var request models.RequestBody
	var response models.Response
	var acRes []models.Action
	var total1 int
	o := orm.NewOrm()

	o.Using("db2")
	body := this.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &request)
	if err != nil {
		beego.Error("数据解析失败!", err)
		this.Data["json"] = models.Result{Code: 500, Data: "", Message: "数据解析失败!"}
		this.ServeJSON()
		return
	}

	//t := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	//没有发送起始时间和结束时间，使用默认方法查询
	//前端传输的是 2019-05-27 10:00:00
	//go 默认为 UTC 时间，需要转换为 CTS 时间
	if request.StartTime == "" && request.EndTime == "" {
		//fmt.Println(typeIndex[request.Type])

		var id []int
		var rTimes []string
		var ip []string
		var port []string
		var dstIp []string
		var dstPort []string
		var res []string
		err = o.Raw("select count(*) from net_log  "+
			"where (serv like ? or app like ?) "+
			"and host_ip like ? "+
			"and src_port like ? "+
			"and dst_ip like ? ",
			"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
			"%"+request.Ip+"%", "%"+request.Port+"%",
			"%"+request.DstIp+"%").QueryRow(&total1)
		//err = o.Raw("select count(*) from two_net_log  "+
		//	"where (serv like ? or app like ?) "+
		//	"and record_time like ? and host_ip like ? "+
		//	"and src_port like ? "+
		//	"and dst_ip like ? ",
		//	"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
		//	"%"+t+"%", "%"+request.Ip+"%", "%"+request.Port+"%",
		//	"%"+request.DstIp+"%").QueryRow(&total2)
		if err != nil {
			beego.Error("获取总数错误: ", err)
			this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取总数失败!"}
			this.ServeJSON()
			return
		}
		sql := fmt.Sprintf("select record_id,record_time,host_ip,"+
			"src_port,dst_ip,serv_port,result from net_log "+
			"where (serv like '%s' or app like '%s') "+
			"and host_ip like '%s' and src_port like '%s' "+
			"and dst_ip like '%s' "+
			"order by record_time desc limit %d,%d",
			"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
			"%"+request.Ip+"%", "%"+request.Port+"%",
			"%"+request.DstIp+"%",
			request.Page*request.Size, request.Size)
		_, err = o.Raw(sql).QueryRows(&id, &rTimes, &ip, &port, &dstIp, &dstPort, &res)
		if err != nil {
			beego.Error("获取信息失败: ", err)
			this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取数据失败!"}
			this.ServeJSON()
			return
		}

		for n, _ := range rTimes {
			var ac models.Action
			ac.Time = rTimes[n]
			ac.Id = id[n]
			ac.Ip = ip[n]
			ac.Port = port[n]
			ac.DstIp = dstIp[n]
			ac.DstPort = dstPort[n]
			input := strings.NewReader(res[n])
			decoder := xml.NewDecoder(input)
			for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
				switch token := t.(type) {
				case xml.CharData:

					content := string([]byte(token))
					//正则匹配 Mac 地址
					reg1 := regexp.MustCompile("[0-9a-fA-F]{2}(-[0-9a-fA-F]{2}){5}")
					//正则匹配 url 地址
					//((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)
					//[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?$
					//reg2 := regexp.MustCompile(`((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)`)

					if reg1.MatchString(content) {
						ac.Mac = content
						//userNetAC.Mac=content
					}
				//	else if reg2.MatchString(content) {
				//	urls = append(urls, content)
				//}
				default:
					//...
				} //end switch
			} //end xml
			acRes = append(acRes, ac)
		}

		//{
		//	var id []int
		//	var rTimes []string
		//	var ip []string
		//	var port []string
		//	var dstIp []string
		//	var dstPort []string
		//	//var res []string
		//	err = o.Raw("select count(*) from two_net_log  "+
		//		"where (serv like ? or app like ?) "+
		//		"and host_ip like ? "+
		//		"and src_port like ? "+
		//		"and dst_ip like ? ",
		//		"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
		//		"%"+request.Ip+"%", "%"+request.Port+"%",
		//		"%"+request.DstIp+"%").QueryRow(&total2)
		//	if err != nil {
		//		beego.Error("获取总数错误: ", err)
		//		this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取总数失败!"}
		//		this.ServeJSON()
		//		return
		//	}
		//	sql := fmt.Sprintf("select record_id,record_time,host_ip,"+
		//		"src_port,dst_ip,serv_port from two_net_log "+
		//		"where (serv like '%s' or app like '%s') "+
		//		"and host_ip like '%s' and src_port like '%s' "+
		//		"and dst_ip like '%s' "+
		//		"order by record_time desc limit %d,%d",
		//		"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
		//		"%"+request.Ip+"%", "%"+request.Port+"%",
		//		"%"+request.DstIp+"%", request.Page*request.Size, request.Size)
		//	_, err = o.Raw(sql).QueryRows(&id, &rTimes, &ip, &port, &dstIp, &dstPort)
		//	if err != nil {
		//		beego.Error("获取信息失败: ", err)
		//		this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取数据失败!"}
		//		this.ServeJSON()
		//		return
		//	}
		//
		//	for n, _ := range rTimes {
		//		var ac models.Action
		//		ac.Time = rTimes[n]
		//		ac.Id = id[n]
		//		ac.Ip = ip[n]
		//		ac.Port = port[n]
		//		ac.DstIp = dstIp[n]
		//		ac.DstPort = dstPort[n]
		//
		//		acRes = append(acRes, ac)
		//	}
		//}

		response.List = acRes
		response.Total = total1

		this.Data["json"] = models.Result{Code: 200, Data: response, Message: "获取数据成功!"}

	} else {
		//var start, end time.Time
		//s := strings.Split(request.StartTime, "(")
		//
		//loc, _ := time.LoadLocation("Asia/Chongqing")
		//
		//start, _ = time.ParseInLocation("Mon Jan 02 2006 15:04:05 GMT+0800 ", s[0], loc)
		//
		//if request.EndTime != "" {
		//	e := strings.Split(request.EndTime, "(")
		//	end, _ = time.ParseInLocation("Mon Jan 02 2006 15:04:05 GMT+0800 ", e[0], loc)
		//} else {
		//	end = time.Now()
		//}
		//
		//fmt.Println(start, end)
		{
			var id []int
			var rTimes []string
			var ip []string
			var port []string
			var dstIp []string
			var dstPort []string
			var res []string

			err = o.Raw("select count(*) from net_log  "+
				"where (serv like ? or app like ?) "+
				"and record_time between ? and ? "+
				"and host_ip like ? "+
				"and src_port like ? "+
				"and dst_ip like ? ",
				"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
				request.StartTime, request.EndTime,
				"%"+request.Ip+"%", "%"+request.Port+"%",
				"%"+request.DstIp+"%").QueryRow(&total1)
			if err != nil {
				beego.Error("获取总数错误: ", err)
				this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取总数失败!"}
				this.ServeJSON()
				return
			}

			sql := fmt.Sprintf("select record_id,record_time,host_ip,"+
				"src_port,dst_ip,serv_port,result from net_log "+
				"where (serv like '%s' or app like '%s') "+
				"and record_time between '%s' and '%s' "+
				"and host_ip like '%s' and src_port like '%s' "+
				"and dst_ip like '%s' "+
				"order by record_time desc limit %d,%d",
				"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
				request.StartTime, request.EndTime,
				"%"+request.Ip+"%", "%"+request.Port+"%",
				"%"+request.DstIp+"%", request.Page*request.Size, request.Size)

			_, err = o.Raw(sql).QueryRows(&id, &rTimes, &ip, &port, &dstIp, &dstPort, &res)
			if err != nil {
				beego.Error("获取信息失败: ", err)
				this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取数据失败!"}
				this.ServeJSON()
				return
			}
			for n, _ := range rTimes {
				var ac models.Action
				ac.Time = rTimes[n]
				ac.Id = id[n]
				ac.Ip = ip[n]
				ac.Port = port[n]
				ac.DstIp = dstIp[n]
				ac.DstPort = dstPort[n]
				input := strings.NewReader(res[n])
				decoder := xml.NewDecoder(input)
				for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
					switch token := t.(type) {
					case xml.CharData:

						content := string([]byte(token))
						//正则匹配 Mac 地址
						reg1 := regexp.MustCompile("[0-9a-fA-F]{2}(-[0-9a-fA-F]{2}){5}")
						//正则匹配 url 地址
						//((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)
						//[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?$
						//reg2 := regexp.MustCompile(`((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)`)

						if reg1.MatchString(content) {
							ac.Mac = content
							//userNetAC.Mac=content
						}
						//	else if reg2.MatchString(content) {
						//	urls = append(urls, content)
						//}
					default:
						//...
					} //end switch
				} //end xml
				acRes = append(acRes, ac)
			}
		}

		//{
		//	var id []int
		//	var rTimes []string
		//	var ip []string
		//	var port []string
		//	var dstIp []string
		//	var dstPort []string
		//	//var res []string
		//
		//	err = o.Raw("select count(*) from two_net_log  "+
		//		"where (serv like ? or app like ?) "+
		//		"and record_time between ? and ? "+
		//		"and host_ip like ? "+
		//		"and src_port like ? "+
		//		"and dst_ip like ? ",
		//		"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
		//		request.StartTime, request.EndTime,
		//		"%"+request.Ip+"%", "%"+request.Port+"%",
		//		"%"+request.DstIp+"%").QueryRow(&total2)
		//	if err != nil {
		//		beego.Error("获取总数错误: ", err)
		//		this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取总数失败!"}
		//		this.ServeJSON()
		//		return
		//	}
		//
		//	sql := fmt.Sprintf("select record_id,record_time,host_ip,"+
		//		"src_port,dst_ip,serv_port from two_net_log "+
		//		"where (serv like '%s' or app like '%s') "+
		//		"and record_time between '%s' and '%s' "+
		//		"and host_ip like '%s' and src_port like '%s' "+
		//		"and dst_ip like '%s' "+
		//		"order by record_time desc limit %d,%d",
		//		"%"+TypeIndex[request.Type]+"%", "%"+TypeIndex[request.Type]+"%",
		//		request.StartTime, request.EndTime,
		//		"%"+request.Ip+"%", "%"+request.Port+"%",
		//		"%"+request.DstIp+"%", request.Page*request.Size, request.Size)
		//
		//	_, err = o.Raw(sql).QueryRows(&id, &rTimes, &ip, &port, &dstIp, &dstPort)
		//	if err != nil {
		//		beego.Error("获取信息失败: ", err)
		//		this.Data["json"] = models.Result{Code: 500, Data: "", Message: "获取数据失败!"}
		//		this.ServeJSON()
		//		return
		//	}
		//	for n, _ := range rTimes {
		//		var ac models.Action
		//		ac.Time = rTimes[n]
		//		ac.Id = id[n]
		//		ac.Ip = ip[n]
		//		ac.Port = port[n]
		//		ac.DstIp = dstIp[n]
		//		ac.DstPort = dstPort[n]
		//		acRes = append(acRes, ac)
		//	}
		//}

		response.List = acRes
		response.Total = total1

		this.Data["json"] = models.Result{Code: 200, Data: response, Message: "获取数据成功!"}
	}

	//beego.Info("Res: ", res)

	this.ServeJSON()
}
