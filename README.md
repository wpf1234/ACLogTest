# 用户上网行为接口
----
## 获取用户上网行为
* 接口名：/action
* 描述：获取信息
* 请求类型：POST
* 请求URL：http://10.2.14.64:7777/ac/action
* 请求参数：

| **请求参数** | **方式** | **类型** | **必须** | **备注** |
| :----: | :----: | :----: | :----: | :----: |
| type | body | int | Y | 访问网站或APP类型(1.邮件、2.网站、3.网盘、4.微信、5.QQ、6.移动端) |
| page | body | int | Y | 页码(从0开始，表示第一页) |
| size | body | int | Y | 每页显示的数量 |
| startTime | body | string | N | 开始时间(用于条件查询) |
| endTime | body | string | N | 结束时间（用于条件查询） |
| ip | body | string | N | ip地址（发件人ip、源ip，用于条件查询） |
| mac | body | string | N | MAC地址（用于条件查询） |
| port | body | string | N | 端口（发件人端口、源ip端口） |
| dstIp | body | string | N | 目标ip（条件查询） |


* 返回参数：

| **返回参数** | **方式** | **类型** | **必须** | **备注** |
| :----: | :----: | :----: | :----: | :----: |
| code | body | int | Y | 状态码（404/500/200） |
| data | body | struct | Y | 返回的数据 |
| message | body | string | Y | 描述信息（获取成功、错误信息） |

* 请求实例：

**Host**
>10.2.14.54

**Header**(对接单点登录后使用)

>
Authorization : eyJhbGciOiJIUzI1NiIsInR5cCI6Imp3dCJ9.eyJlbWFpbCI6IjE4MDQzNSIsImV4cCI6MTU1NzcyMDk4OCwiaWF0IjoxNTU3NzEwMTg4fQ.mWC54-V9Q6HAQBqOYyhbnEzJxfkM_zbfJe4yhTiRvKM

**URL**

>
http://10.2.14.64:7777/ac/action

**Body**
无开始时间和结束时间

```json
{
	"type":1,
	"page":0,
	"size":10,
	"startTime":"",
	"endTime":"",
	"ip":"",
	"mac":"",
	"port":"",
	"dstIp":""
}
```

有开始时间和结束时间及其他查询条件
```json
{
	"type":1,
	"page":0,
	"size":10,
	"startTime":"2019-05-13 00:00:00",
	"endTime":"2019-05-13 10:00:00",
	"ip":"172.16.136.205",
	"mac":"",
	"port":"",
	"dstIp":""
}
```

**Response**
success

```json
{
    "code": 200,
    "data": {
        "list": [
            {
                "id": 26,
                "time": "2019-05-13 08:16:29",
                "ip": "172.16.70.50",
                "port": "7273",
                "mac": "e4-c7-22-04-2e-c1",
                "dstIp": "180.97.33.12:8829" 
            },
            {
                "id": 42,
                "time": "2019-05-13 08:16:49",
                "ip": "172.16.70.50",
                "port": "49425",
                "mac": "d8-67-d9-7b-44-41",
                "dstIp": "180.149.145.241:443",
            }
        ],
        "total": 2
    },
    "message": "获取数据成功!"
}
```
failed
```json
{
	"code" : 500,
	"data" : "",
  	"message": "获取信息失败!/数据解析错误!"
}
```

----
