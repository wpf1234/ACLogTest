package models

type TokenResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"message"`
}

type TokenInfo struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data TokenUserInfo `json:"data"`
}

type TokenUserInfo struct {
	ApiAuth     []string `json:"apiauth"`
	Avatar      string   `json:"avatar"`
	AvatarName  string   `json:"avatarName"`
	DisplayName string   `json:"displayName"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	RealName    string   `json:"realName"`
	UiAuth      []string `json:"uiauth"`
}
