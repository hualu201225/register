package HttpCurl

import(
	"net/http"
	"net/url"
	_"../Common"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http/cookiejar"
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
	//是否需要保存cookie
	needCookies bool
	//存放cookie
	curCookies []*http.Cookie

	curCookieJar *cookiejar.Jar
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

func (HttpCurl *HttpCurl) SetNeedCookie(isNeedCookie bool) {
	HttpCurl.needCookies = isNeedCookie
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

func (HttpCurl *HttpCurl) initCookieJar() {
	HttpCurl.curCookies = nil
	HttpCurl.curCookieJar, _ = cookiejar.New(nil)
}

func (HttpCurl *HttpCurl) printCurCookies() {
    var cookieNum int = len(HttpCurl.curCookies);
    fmt.Printf("cookieNum=%d\r\n", cookieNum)
    for i := 0; i < cookieNum; i++ {
        var curCk *http.Cookie = HttpCurl.curCookies[i];
        fmt.Printf("curCk.Raw=%s\r\n", curCk.Value)
    }
}


func (HttpCurl *HttpCurl) httpCurl(method string) ([]byte, error) {
	HttpCurl.initCookieJar()
	client := &http.Client{
		Jar : HttpCurl.curCookieJar,
	}
	urlQuery := HttpCurl.getGetUrl()
	fmt.Println(method)
	fmt.Println(urlQuery)
	//添加post参数
	var urlPost string
	if (method == "POST") {
		data := url.Values{}
		for k, v := range HttpCurl.postData {
			data.Add(k, v)
		}
		//u, _ := url.ParseRequestURI(urlQuery)
		urlPost = data.Encode()
	}
	fmt.Printf(urlPost)

	//提交请求
	request, err := http.NewRequest(method, urlQuery, strings.NewReader(urlPost))
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
	// fmt.Printf(string(str))	
	// fmt.Println(response.StatusCode)
	if (err != nil) {
		fmt.Printf(string(str))
		panic("can not read response")
	}

	http.HandleFunc("/", HttpCurl.set)
	http.HandleFunc("/read", HttpCurl.read)

	HttpCurl.curCookies = HttpCurl.curCookieJar.Cookies(request.URL)
	HttpCurl.printCurCookies()
	return str, err
}

func (HttpCurl *HttpCurl) set(w http.ResponseWriter, req *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:  "my-cookie",
        Value: "some value",
    })
    //fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func (HttpCurl *HttpCurl) read(w http.ResponseWriter, req *http.Request) {

    c, err := req.Cookie("my-cookie")
    if err != nil {
        http.Error(w, http.StatusText(400), http.StatusBadRequest)
        return
    }

    fmt.Fprintln(w, "YOUR COOKIE:", c)
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

