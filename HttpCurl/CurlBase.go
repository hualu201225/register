package HttpCurl

import(
	"net/http"
	_"../Common"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type HttpCurl struct {
	//请求的url
	url string
	//header头参数
	headers map[string]string
	//get参数
	queries map[string]string
	//post参数
	postData map[string]interface{}
}

func (HttpCurl *HttpCurl) SetUrl(url string) {
	HttpCurl.url = url
}

func (HttpCurl *HttpCurl) SetHeaders(headers map[string]string) {
	HttpCurl.headers = headers

}

func (HttpCurl *HttpCurl) httpGet() (map[string]interface{}, error) {
	client := &http.Client{}
	url := HttpCurl.url

	//get参数处理
	if (HttpCurl.queries != nil) {
		//HttpCurl.queries = Common.MapTrans.MapToString(HttpCurl.queries)
		url = url + "?1=1"
		queries := make([]string, 0, len(HttpCurl.queries))
		for k, v := range HttpCurl.queries {
			queries = append(queries, k + "=" + v)
		}
		url = url + strings.Join(queries, "&")

	}

	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	if (err != nil) {
		panic("can not new request")
	}

	//添加header
	for k, v := range HttpCurl.headers {
		request.Header.Add(k, v)
	}

	response, _ := client.Do(request)
	defer response.Body.Close()

	str, _ := ioutil.ReadAll(response.Body)
	fmt.Printf(string(str))

	res := make(map[string]interface{})
	error := json.Unmarshal(str, res)

	return res, error
}

func (HttpCurl *HttpCurl) httpPost() (map[string]interface{}, error) {
	res := make(map[string]interface{})
	return res, nil
}

func (HttpCurl *HttpCurl) GetContentsFromUrl() (map[string]interface{}, error) {
	//校验url
	if (HttpCurl.url == "") {
		panic("url is empty")
	}

	res := make(map[string]interface{})
	//Get形式
	if (HttpCurl.postData == nil) {
		res, _  = HttpCurl.httpGet()
	} else {
		res, _  = HttpCurl.httpPost()
	}

	return res, nil
}

