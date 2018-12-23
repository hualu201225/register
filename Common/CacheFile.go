package Common

import(
	"fmt"
	"os"
	"io/ioutil"
)

type CacheFile struct {
	filePath string
	fileName string
    namePrefix string
}

func (CacheFile *CacheFile) SetFilePath(filePath string) {
	CacheFile.filePath = filePath
}

func (CacheFile *CacheFile) SetFileName(fileName string) {
	CacheFile.fileName = fileName
}

//创建File存放的目录
func (CacheFile *CacheFile) createFilePath(FilePath string) {
	//检测目录是否存在
	_, err := os.Stat(FilePath)
	//不存在则建立该目录
	if (err != nil || os.IsNotExist(err)) {
		errMk := os.Mkdir(FilePath, os.ModePerm)
		if (errMk != nil) {
			fmt.Printf("mkdir failed!")
		} else {
			fmt.Printf("mkdir success")
		}
	}
}

//获取File存放的目录
func (CacheFile *CacheFile) getFileFullPath() string {
	if (len(CacheFile.filePath) == 0) {
		CacheFile.filePath = "./Storage/Files"
	}

	//若该目录不存在则创建
	CacheFile.createFilePath(CacheFile.filePath)

	fullPath := CacheFile.filePath + "/" + CacheFile.fileName
	return fullPath
}

func (CacheFile *CacheFile) getFileFileName() (string) {
	return CacheFile.fileName
}

func (CacheFile *CacheFile) SetContentToFile(info string) error {
	//获取File全路径
	FilePath := CacheFile.getFileFullPath()

	fmt.Printf(FilePath)

	//打开文件，获取文件指针
	FileFile, error := os.Create(FilePath)
	if (error != nil) {
		return error
	}

	defer FileFile.Close()

	//写入内容
	_, writeErr := FileFile.WriteString(info)
	
	return writeErr
}

func (CacheFile *CacheFile) GetContentFromFile() string {
	//获取File全路径
	FilePath := CacheFile.getFileFullPath()

	FileFile, error := os.Open(FilePath)
	if (error != nil) {
		panic("can not open FileFile")
	}

	defer FileFile.Close()  //在该函数即将返回前才执行

	//该方法读取不到中文
	//b1 := make([]byte, 5)
	//_, readErr := FileFile.Read(b1)

	res, readErr := ioutil.ReadAll(FileFile)
	if (readErr != nil) {
		panic("can not read FileFile")
	}

	return string(res)
}