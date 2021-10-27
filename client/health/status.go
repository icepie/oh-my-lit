package health

import (
	"encoding/json"
	"errors"
	"fmt"
)

// GetHealthyReportCount 获取上报统计
// identity 身份 (student, teacher, other) 如果为空字符串则为全部 , orgIDs 部门列表, createTime 时间 2021-10-27
func (u *HealthUser) GetHealthyReportCount(identity string, orgIDs []uint, createTime string) (healthyReportCount HealthyReportCount, err error) {

	//
	orgIDsStr := ""

	for i, orgID := range orgIDsStr {
		orgIDsStr += fmt.Sprint(orgID)
		if i != len(orgIDsStr)-1 {
			orgIDsStr += ","
		}
	}

	resp, err := u.Client.R().
		SetQueryParams(map[string]string{
			"teamId":          fmt.Sprint(u.UserInfo.TeamID),
			"userId":          fmt.Sprint(u.UserInfo.UserID),
			"identity":        fmt.Sprint(IdentityNameList[identity]),
			"organizationIds": orgIDsStr,
			"createTime":      createTime,
		}).SetResult(Result{}).
		Get(HealthyReportCountUrl)

	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	if r.Code != 200 {
		err = errors.New(r.Msg)
		return
	}

	byteData, _ := json.Marshal(r.Data)
	err = json.Unmarshal(byteData, &healthyReportCount)
	if err != nil {
		return
	}

	return
}
