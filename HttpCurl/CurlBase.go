package HttpCurl

import(
	"net/http"
	"net/url"
	_"../Common"
	"strings"
	"fmt"
	"io/ioutil"
)

type HttpCurl struct {
	//请求的url
	url string
	//header头参数
	headers map[string]string
	//get参数
	queries map[string]string
	//post参数
	postData map[string]string
}

func (HttpCurl *HttpCurl) SetUrl(url string) {
	HttpCurl.url = url
}

func (HttpCurl *HttpCurl) SetHeaders(headers map[string]string) {
	HttpCurl.headers = headers
}

func (HttpCurl *HttpCurl) SetQueries(queries map[string]string) {
	HttpCurl.queries = queries
}

func (HttpCurl *HttpCurl) SetPostData(postData map[string]string) {
	HttpCurl.postData = postData
}

func (HttpCurl *HttpCurl) getGetUrl() string {
	url := HttpCurl.url
	//get参数处理
	if (HttpCurl.queries != nil) {
		//HttpCurl.queries = Common.MapTrans.MapToString(HttpCurl.queries)
		url = url + "?1=1&"
		queries := make([]string, 0, len(HttpCurl.queries))
		for k, v := range HttpCurl.queries {
			queries = append(queries, k + "=" + v)
		}
		url = url + strings.Join(queries, "&")
	}

	return url
}


func (HttpCurl *HttpCurl) httpCurl(method string) ([]byte, error) {
	client := &http.Client{}
	urlQuery := HttpCurl.getGetUrl()
	fmt.Println(method)

	//添加post参数
	if (method == "POST") {
		data := url.Values{}
		for k, v := range HttpCurl.postData {
			data.Add(k, v)
		}
		u, _ := url.ParseRequestURI(urlQuery)
		u.RawQuery = data.Encode()
		urlQuery = fmt.Sprintf("%v", u)
	}

	//提交请求
	request, err := http.NewRequest(method, urlQuery, nil)
	if (err != nil) {
		panic("can not new request")
	}

	//添加header
	for k, v := range HttpCurl.headers {
		request.Header.Add(k, v)
	}

	response, err := client.Do(request)
	if (err != nil) {
		fmt.Println(err)
		panic("can not get response")
	}

	defer response.Body.Close()	

	str, err := ioutil.ReadAll(response.Body)
	fmt.Printf(string(str))	
	if (err != nil) {
		fmt.Println(response.StatusCode)
		fmt.Printf(string(str))
		panic("can not read response")
	}

	return str, err
}

func (HttpCurl *HttpCurl) GetContentsFromUrl() ([]byte, error) {
	//校验url
	if (HttpCurl.url == "") {
		panic("url is empty")
	}

	var method string
	//Get形式
	if (HttpCurl.postData == nil) {
		method = "GET"
	//Post形式
	} else {
		method = "POST"
	}

	res, _  := HttpCurl.httpCurl(method)

	return res, nil
}

