package Identity

import(
	_"fmt"
	"../Common"
)

type Cookie struct {
	cookiePath string
    namePrefix string
    CacheFile *Common.CacheFile
}


//获取cookie存放的目录
func (cookie *Cookie) getCookiePath() string {
	cookiePath := "./Storage/Cache/Cookie"
	return cookiePath
}

func (cookie *Cookie) getCookieFileName(identityCode string) (string) {
	fileNamePrefix := "12580_cookie_"
	if (identityCode == "") {
		panic("identityCode is empty")
	}
	cookieFileName := fileNamePrefix + identityCode + ".txt"

	return cookieFileName
}


func (cookie *Cookie) SetCookie(info string, identityCode string) error {
	//获取cookie名称
	cookieFileName := cookie.getCookieFileName(identityCode)
	
	//获取cookie存放路径
	cookiePath := cookie.getCookiePath()

	//设置cookie
    cookie.CacheFile = &Common.CacheFile{}
	cookie.CacheFile.SetFilePath(cookiePath)
	cookie.CacheFile.SetFileName(cookieFileName)
	writeErr := cookie.CacheFile.SetContentToFile(info)
	return writeErr
}

func (cookie *Cookie) GetCookie(identityCode string) string {
	//获取cookie名称
	cookieFileName := cookie.getCookieFileName(identityCode)
	
	//获取cookie存放路径
	cookiePath := cookie.getCookiePath()

	cookie.CacheFile = &Common.CacheFile{}
	cookie.CacheFile.SetFilePath(cookiePath)
	cookie.CacheFile.SetFileName(cookieFileName)
	res := cookie.CacheFile.GetContentFromFile()

	return string(res)
}