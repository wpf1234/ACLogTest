package models

type RequestBody struct {
	Type int `json:"type"`
	//Date DateTime `json:"date"`
	Page int64 `json:"page"`
	Size int64 `json:"size"`

	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Ip        string `json:"ip"`
	Mac       string `json:"mac"`
	Port      string `json:"port"`
	DstIp     string `json:"dstIp"`
}

type DateTime struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
