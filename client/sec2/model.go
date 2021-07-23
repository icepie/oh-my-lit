package sec2

// HomeParam 主页参数模型
type HomeParam struct {
	Code int64 `json:"code"`
	Data struct {
		Backend      string      `json:"backend"`
		CopyRight    string      `json:"copy_right"`
		HasBeianPerm bool        `json:"hasBeianPerm"`
		HasLibrary   bool        `json:"hasLibrary"`
		IsLocalUser  bool        `json:"isLocalUser"`
		Organization string      `json:"organization"`
		PlacardShow  interface{} `json:"placard_show"`
		Protocols    []string    `json:"protocols"`
		SysPlacard   interface{} `json:"sys_placard"`
		Username     string      `json:"username"`
	} `json:"data"`
	Msg string `json:"msg"`
}
