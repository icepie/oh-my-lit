package jw

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func md5Encode(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// chkpwd 将用户密码进行处理
func chkpwd(username string, password string) string {

	/* JavaScript
	function chkpwd(obj)
	{
		var schoolcode="11070";
		var yhm=document.all.txt_sdsdfdsfryuiighgdf.value;
		if(obj.value!="")
		{
			if(document.all.Sel_Type.value=="ADM")
				yhm=yhm.toUpperCase();
			var s=md5(yhm+md5(obj.value).substring(0,30).toUpperCase()+schoolcode).substring(0,30).toUpperCase();
			document.all.sdfdfdhgwerewt.value=s;
		}
		else
		{
			document.all.sdfdfdhgwerewt.value=obj.value;
		}
	}
	*/

	return strings.ToUpper(md5Encode(username + strings.ToUpper(md5Encode(password)[0:30]) + SchoolCode)[0:30])
}
