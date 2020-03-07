package util

import (
	"ACTest/models"
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//func HandleAC(colFam string,userAC []models.UserAction) []map[string]map[string][]byte{
//	var ACResult = make([]map[string]map[string][]byte,0)
//	var t xml.Token
//	var err error
//
//	for _,v:=range userAC{
//		var result=make(map[string]map[string][]byte)
//		var urls []string
//		var userNetAC models.UserNetAction
//		input:=strings.NewReader(v.Result)
//		decoder:=xml.NewDecoder(input)
//		for t,err=decoder.Token();err == nil;t,err=decoder.Token(){
//			switch token := t.(type){
//			case xml.CharData:
//
//				content:=string([]byte(token))
//				//正则匹配 Mac 地址
//				reg1:=regexp.MustCompile("[0-9a-fA-F]{2}(-[0-9a-fA-F]{2}){5}")
//				//正则匹配 url 地址
//				reg2:=regexp.MustCompile(`[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?$`)
//				if reg1.MatchString(content){
//					userNetAC.Mac=content
//				}else if reg2.MatchString(content){
//					urls=append(urls,content)
//				}
//			default:
//				//...
//			}//end switch
//		}//end xml
//		urls=Distinct(urls)
//
//		if len(urls) > 1{
//			for j:=0;j<len(urls)-1;j++{
//				if len(urls[j]) > len(urls[j+1]){
//					userNetAC.Resource=urls[j]
//				}else{
//					userNetAC.Resource=urls[j+1]
//				}
//			}
//		}else if len(urls) == 0{
//			userNetAC.Resource=""
//		}else{
//			userNetAC.Resource=urls[0]
//		}
//
//		userNetAC.User=v.User
//		userNetAC.Server=v.Server
//		userNetAC.App=v.App
//		userNetAC.HostIp=v.HostIp
//		userNetAC.SrcPort=strconv.Itoa(v.SrcPort)
//		userNetAC.DstIp=v.DstIp
//		userNetAC.ServPort=strconv.Itoa(v.ServPort)
//
//		//结构体转换为 map
//		key:=reflect.TypeOf(userNetAC)
//		value:=reflect.ValueOf(userNetAC)
//		var m=make(map[string][]byte)
//		for i:=0;i<key.NumField();i++{
//			m[key.Field(i).Name]=[]byte(value.Field(i).String())
//			result[colFam]=m
//			ACResult=append(ACResult,result)
//		}
//
//	}
//
//	return ACResult
//}

//func HandleAC(userAC []models.UserAction) []map[string]string {
//	var ACRes = make([]map[string]string, 0)
//	var t xml.Token
//	var err error
//
//	for _, v := range userAC {
//		var result = make(map[string]string)
//		var urls []string
//		var userNetAC models.UserNetAction
//		input := strings.NewReader(v.Result)
//		decoder := xml.NewDecoder(input)
//		for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
//			switch token := t.(type) {
//			case xml.CharData:
//
//				content := string([]byte(token))
//				//正则匹配 Mac 地址
//				reg1 := regexp.MustCompile("[0-9a-fA-F]{2}(-[0-9a-fA-F]{2}){5}")
//				//正则匹配 url 地址
//				reg2 := regexp.MustCompile(`[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?$`)
//				if reg1.MatchString(content) {
//					userNetAC.Mac = content
//				} else if reg2.MatchString(content) {
//					urls = append(urls, content)
//				}
//			default:
//				//...
//			} //end switch
//		} //end xml
//		urls = Distinct(urls)
//
//		if len(urls) > 1 {
//			for j := 0; j < len(urls)-1; j++ {
//				if len(urls[j]) > len(urls[j+1]) {
//					userNetAC.Resource = urls[j]
//				} else {
//					userNetAC.Resource = urls[j+1]
//				}
//			}
//		} else if len(urls) == 0 {
//			userNetAC.Resource = ""
//		} else {
//			userNetAC.Resource = urls[0]
//		}
//
//		userNetAC.RecordId = strconv.Itoa(v.RecordId)
//		userNetAC.User = v.User
//		userNetAC.Server = v.Server
//		userNetAC.App = v.App
//		userNetAC.HostIp = v.HostIp
//		userNetAC.SrcPort = strconv.Itoa(v.SrcPort)
//		userNetAC.DstIp = v.DstIp
//		userNetAC.ServPort = strconv.Itoa(v.ServPort)
//		userNetAC.RecordTime = v.RecordTime
//
//		//结构体转换为 map
//		key := reflect.TypeOf(userNetAC)
//		value := reflect.ValueOf(userNetAC)
//
//		for i := 0; i < key.NumField(); i++ {
//			result[key.Field(i).Name] = value.Field(i).String()
//
//			ACRes = append(ACRes, result)
//		}
//
//	}
//
//	return ACRes
//}

func HandleACTest(userAC []models.UserAction) []string {
	var ACRes []string
	var t xml.Token
	var err error
	tm := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	for _, v := range userAC {
		var result string
		var urls []string
		var mac string
		var resource string
		input := strings.NewReader(v.Result)
		decoder := xml.NewDecoder(input)
		for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
			switch token := t.(type) {
			case xml.CharData:

				content := string([]byte(token))
				//正则匹配 Mac 地址
				reg1 := regexp.MustCompile("[0-9a-fA-F]{2}(-[0-9a-fA-F]{2}){5}")
				//正则匹配 url 地址
				//((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)
				//[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?$
				reg2 := regexp.MustCompile(`((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)`)

				if reg1.MatchString(content) {
					mac = content
					//userNetAC.Mac=content
				} else if reg2.MatchString(content) {
					urls = append(urls, content)
				}
			default:
				//...
			} //end switch
		} //end xml
		urls = Distinct(urls)

		if len(urls) > 1 {
			for j := 0; j < len(urls)-1; j++ {
				if len(urls[j]) > len(urls[j+1]) {

					resource = urls[j]
				} else {

					resource = urls[j+1]
				}
			}
		} else if len(urls) == 0 {

			resource = ""
		} else {

			resource = urls[0]
		}
		rTime := tm + " " + v.RecordTime
		result = strconv.Itoa(v.RecordId) + "," + v.User + "," + v.HostIp + "," + strconv.Itoa(v.SrcPort) +
			"," + mac + "," + v.DstIp + "," + strconv.Itoa(v.ServPort) + "," + v.Server +
			"," + v.App + "," + rTime + "," + resource

		ACRes = append(ACRes, result)

	}

	return ACRes
}
