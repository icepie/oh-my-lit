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

// CurrentMember 当前用户返回结构
type CurrentMemberRte struct {
	Attributes interface{} `json:"attributes"`
	Count      interface{} `json:"count"`
	Msg        string      `json:"msg"`
	Obj        struct {
		LastLoginTime           string      `json:"lastLoginTime"`
		MemberAcademicNumber    string      `json:"memberAcademicNumber"`
		MemberCreateTime        int64       `json:"memberCreateTime"`
		MemberID                string      `json:"memberId"`
		MemberIDNumber          string      `json:"memberIdNumber"`
		MemberImage             interface{} `json:"memberImage"`
		MemberMailbox           interface{} `json:"memberMailbox"`
		MemberNickname          string      `json:"memberNickname"`
		MemberOtherBirthday     interface{} `json:"memberOtherBirthday"`
		MemberOtherClass        interface{} `json:"memberOtherClass"`
		MemberOtherDepartment   interface{} `json:"memberOtherDepartment"`
		MemberOtherGrade        interface{} `json:"memberOtherGrade"`
		MemberOtherMajor        interface{} `json:"memberOtherMajor"`
		MemberOtherNation       interface{} `json:"memberOtherNation"`
		MemberOtherNative       interface{} `json:"memberOtherNative"`
		MemberOtherSchoolNumber interface{} `json:"memberOtherSchoolNumber"`
		MemberPhone             interface{} `json:"memberPhone"`
		MemberPwd               interface{} `json:"memberPwd"`
		MemberSex               int64       `json:"memberSex"`
		MemberSign              interface{} `json:"memberSign"`
		MemberState             int64       `json:"memberState"`
		MemberUsername          string      `json:"memberUsername"`
		RoleCodeList            []string    `json:"roleCodeList"`
		RoleList                []struct {
			RoleCode    string      `json:"roleCode"`
			RoleComment interface{} `json:"roleComment"`
			RoleName    string      `json:"roleName"`
			RoleState   interface{} `json:"roleState"`
		} `json:"roleList"`
	} `json:"obj"`
	Success bool `json:"success"`
}

// GetStudentRte 查询学生返回结构
type GetStudentRte struct {
	Attributes interface{} `json:"attributes"`
	Count      interface{} `json:"count"`
	Msg        string      `json:"msg"`
	Obj        struct {
		StudentAdmissionTime    string      `json:"studentAdmissionTime"`
		StudentAdress           interface{} `json:"studentAdress"`
		StudentBirthday         string      `json:"studentBirthday"`
		StudentCategory         string      `json:"studentCategory"`
		StudentClassCode        string      `json:"studentClassCode"`
		StudentClassName        string      `json:"studentClassName"`
		StudentEductionalSystme string      `json:"studentEductionalSystme"`
		StudentFacultiesCode    string      `json:"studentFacultiesCode"`
		StudentFacultiesName    string      `json:"studentFacultiesName"`
		StudentGrade            string      `json:"studentGrade"`
		StudentID               string      `json:"studentId"`
		StudentIDNumber         string      `json:"studentIdNumber"`
		StudentMajor            string      `json:"studentMajor"`
		StudentMajorName        string      `json:"studentMajorName"`
		StudentName             string      `json:"studentName"`
		StudentNation           interface{} `json:"studentNation"`
		StudentPhone            string      `json:"studentPhone"`
		StudentPoliticalStatus  interface{} `json:"studentPoliticalStatus"`
		StudentRegisterState    string      `json:"studentRegisterState"`
		StudentSex              string      `json:"studentSex"`
	} `json:"obj"`
	Success bool `json:"success"`
}
