package Identity

import(
	"fmt"
	"os"
	"io/ioutil"
)

type Cookie struct {
	cookiePath string
    namePrefix string
}


//创建cookie存放的目录
func (cookie *Cookie) createCookiePath(cookiePath string) {
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
}

//获取cookie存放的目录
func (cookie *Cookie) getCookiePath(filename string) string {
	cookiePath := "./cookies"

	Cookie := &Cookie{}
	Cookie.createCookiePath(cookiePath)

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


func (cookie *Cookie) SetCookie(info string, identityCode string) error {
	Cookie := &Cookie{}
	//获取cookie名称
	cookieFileName,err := Cookie.getCookieFileName(identityCode)
	if (err != nil) {
		return fmt.Errorf("can get cookieFileName")
	}
	//获取cookie存放路径
	cookiePath := Cookie.getCookiePath(cookieFileName)

	//打开文件，获取文件指针
	cookieFile, error := os.Create(cookiePath)
	if (error != nil) {
		return error
	}
	
	defer cookieFile.Close()

	//写入内容
	_, writeErr := cookieFile.WriteString(info)
	
	return writeErr
}

func (cookie *Cookie) GetCookie(identityCode string) string {
	Cookie := &Cookie{}
	//获取cookie名称
	cookieFileName, err := Cookie.getCookieFileName(identityCode)
	if (err != nil) {
		panic("can not get cookie")
	}

	//获取cookie存放路径
	cookiePath := Cookie.getCookiePath(cookieFileName)

	cookieFile, error := os.Open(cookiePath)
	if (error != nil) {
		panic("can not open cookieFile")
	}

	defer cookieFile.Close()  //在该函数即将返回前才执行

	//该方法读取不到中文
	//b1 := make([]byte, 5)
	//_, readErr := cookieFile.Read(b1)

	res, readErr := ioutil.ReadAll(cookieFile)
	if (readErr != nil) {
		panic("can not read cookieFile")
	}

	return string(res)
}