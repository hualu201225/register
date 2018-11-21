package Identity

import(
	"fmt"
	"os"
)

type Cookie struct {
	cookiePath string
    namePrefix string
}


//创建cookie存放的目录
func (cookie *Cookie) createCookiePath(cookiePath string, filename string) {
	//检测目录是否存在
	_, err := os.Stat(cookiePath)
	//不存在则建立该目录
	if (err != nil || os.IsNotExist(err)) {
		errMk := os.Mkdir(cookiePath, os.ModePerm)
		if (errMk != nil) {
			fmt.Printf("mkdir failed!")
		} else {
			fmt.Printf("mkdir success")
		}
	}

	//创建文件
	os.Create(cookiePath + "/" + filename)
}

//获取cookie存放的目录
func (cookie *Cookie) getCookiePath(filename string, path string) string {
	cookiePath := "./cookies"
	if (path != "") {
		cookiePath = path
	}

	Cookie := &Cookie{}
	Cookie.createCookiePath(cookiePath, filename)

	return cookiePath + "/" + filename
}

func (cookie *Cookie) getCookieFileName(identityCode string) (string,error) {
	fileNamePrefix := "12580_cookie_"
	if (identityCode == "") {
		return "", fmt.Errorf("error") 
	}
	cookieFileName := fileNamePrefix + identityCode + ".txt"

	return cookieFileName, nil
}


func (cookie *Cookie) SetCookie(info , identityCode string) error {
	Cookie := &Cookie{}
	//获取cookie名称
	cookieFileName,err := Cookie.getCookieFileName(identityCode)
	if (err != nil) {
		return fmt.Errorf("can get cookieFileName")
	}
	//获取cookie存放路径
	cookiePath := Cookie.getCookiePath(cookieFileName, "")
	
	//打开文件，获取文件指针
	cookieFile, error := os.Open(cookiePath)
	if (error != nil) {
		return error
	}
	
	//写入内容
	writeStr, writeErr := cookieFile.WriteString("fadfadsfdsa" + info)
	fmt.Printf("wrote %d bytes", writeStr)
	cookieFile.Close()
	return writeErr
}