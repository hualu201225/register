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
	//响应的cookie字符串
	responseCookie string

	needUrlParse bool
}

func (HttpCurl *HttpCurl) SetUrl(url string) {
	HttpCurl.url = url
}

func (HttpCurl *HttpCurl) SetNeedUrlparse(isneed bool) {
	HttpCurl.needUrlParse = isneed
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
	urlNormal := HttpCurl.url

	var urlStr string
	//get参数处理
	if (HttpCurl.queries != nil) {
		//HttpCurl.queries = Common.MapTrans.MapToString(HttpCurl.queries)
		urlStr = "?1=1&"
		queries := make([]string, 0, len(HttpCurl.queries))
		for k, v := range HttpCurl.queries {
			queries = append(queries, k + "=" + v)
		}
		urlStr = urlStr + strings.Join(queries, "&")
	}
	
	if (HttpCurl.needUrlParse == true) {
		urlParse, _ := url.Parse(urlStr)
		urlStr = "?" + urlParse.Query().Encode()
	}
	
	urlStr = urlNormal + urlStr
	return urlStr
}

func (HttpCurl *HttpCurl) saveCookies(response *http.Response) {
	cookies := response.Cookies()

	fmt.Printf("response Cookies :%v", cookies)
	var cookieStr string
    for _, cookie := range cookies {
    	//cookieStr = cookieStr + "；" + cookie    
    	// fmt.Sprintf("%s", cookie)
        cookieStr = cookieStr + fmt.Sprintf("%s;", cookie)
    }	
    
    HttpCurl.responseCookie = cookieStr
}

func (HttpCurl *HttpCurl) GetCookies() string {
	return HttpCurl.responseCookie
}

func (HttpCurl *HttpCurl) transferPostData(method string) string {
	var urlPost string
	if (method == "POST") {
		data := url.Values{}
		for k, v := range HttpCurl.postData {
			data.Add(k, v)
		}
		urlPost = data.Encode()
	}
	fmt.Println(urlPost)
	return urlPost
}


func (HttpCurl *HttpCurl) httpCurl(method string) ([]byte, error) {

	client := &http.Client{}

	//获取url
	urlQuery := HttpCurl.getGetUrl()
	fmt.Println(method)
	fmt.Println(urlQuery)

	//添加post参数
	urlPost := HttpCurl.transferPostData(method)

	//初始化请求
	request, err := http.NewRequest(method, urlQuery, strings.NewReader(urlPost))
	if (err != nil) {
		panic("can not new request")
	}

	//添加header
	for k, v := range HttpCurl.headers {
		request.Header.Add(k, v)
	}

	//发送请求
	response, err := client.Do(request)
	if (err != nil) {
		panic(err)
	}

	defer response.Body.Close()

	//cookie处理
	HttpCurl.saveCookies(response)

	str, err := ioutil.ReadAll(response.Body)

	// fmt.Printf(string(str))	
	fmt.Println(response.StatusCode)
	if (err != nil) {
		//fmt.Printf(string(str))
		panic(err)
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

