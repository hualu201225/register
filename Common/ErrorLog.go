package Common

type ErrorLog struct {
	LogPath string
	LogName string
}

func (ErrorLog *ErrorLog) init() {
	ErrorLog.LogPath = "./Storage/Cache/Logs/"
}

func (ErrorLog *ErrorLog) SetError(errors string) {
	ErrorLog.init()
	CacheFile := &CacheFile{}
	CacheFile.SetFilePath(ErrorLog.LogPath)
	CacheFile.SetFileName(ErrorLog.LogName)
	CacheFile.SetContentToFile(errors)
}