package common

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//定义统一的返回格式
type ReturnMsg struct {
	Code    int         //状态码，成功为0，其他均为错误
	Message interface{} //成功跟错误的消息
	Data    interface{} //data数据
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

var Domain string
var ch chan string

func init() {
	Domain = "https://6000.vip"
}

func SpeedTest(DomainList string) {
	tmpArr := strings.Split(DomainList, ",")
	//站点筛选
	ch = make(chan string, len(tmpArr))
	for _, v := range tmpArr {
		go selectDomain(v)
	}
	fmt.Println("Domain", Domain)
	for v := range ch {
		if v != "null" {
			Domain = v
			return
		}
	}
}

//选择域名
func selectDomain(url string) {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("%v", err)
			for i := 0; i < 10; i++ {
				funcName, file, line, ok := runtime.Caller(i)
				if ok {
					fmt.Println("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
				}
			}
			ch <- "null"

		}
	}()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jar, _ := cookiejar.New(nil)
	//url := "http://www.dsdv00"+strconv.Itoa(i) + ".com"
	client := &http.Client{
		Jar:       jar,
		Transport: tr,
	}

	resp, err := client.Get(url)
	if err != nil {
		ch <- "null"
		return
	}
	defer resp.Body.Close()
	ch <- url
}

//const Domain string = "http://lar.boya6.net"
type Util struct {
	Url         string
	Jar         *cookiejar.Jar
	Params      map[string]string // a=b&c=d
	ParamsBytes []byte
	Before      bool
	//重新覆盖cookie
	RefreshCookie bool
	//添加cookie
	AddCookie bool
	Header    map[string]string
	WebSite   string
}

/**
post请求
传入地址，传入参数
type 传入POST,GET.使用json发送post数据
*/
func (c *Util) MethodJson(Type string) string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cli := &http.Client{
		Transport: tr,
		Jar:       c.Jar,
	}
	cli.Timeout = 10 * time.Second
	//fmt.Println("c.ParamsBytes",c.ParamsBytes)
	//进行第一次请求获取cookie,并且获取验证码
	//fmt.Println("c.Url",c.Url)
	req, reqErr := http.NewRequest(Type, c.Url, bytes.NewBuffer(c.ParamsBytes))
	if reqErr != nil {
		fmt.Println("http error %v", reqErr)
	}
	//req.TLS =
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36")

	if c.Header != nil {
		for k, v := range c.Header {
			req.Header.Set(k, v)
		}
	}
	resP, reqErr := cli.Do(req)
	//fmt.Println("req",req)
	if reqErr != nil {
		fmt.Println("cli do error %v", reqErr.Error())
	} else {
		//fmt.Println("resP",resP.Request.Header)
	}

	body, bodyErr := ioutil.ReadAll(resP.Body)

	if bodyErr != nil {
		fmt.Println("http error %v", reqErr.Error())
	}
	defer resP.Body.Close()

	return string(body[:])
}

/**
post请求
传入地址，传入参数
type 传入POST,GET
*/
func (c *Util) MethodRedirect(Type string) string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//建立表单
	var form http.Request
	form.ParseForm()
	for k, v := range c.Params {
		form.Form.Add(k, v)
	}
	bodystr := strings.TrimSpace(form.Form.Encode())

	cli := &http.Client{Transport: tr}
	jar, _ := cookiejar.New(nil)
	cli.Jar = jar
	cli.Timeout = 10 * time.Second
	cli.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		fmt.Println(req.URL, "req", req.Response)
		fmt.Println("jar.cookie", jar.Cookies(req.URL))
		return nil
	}

	//进行第一次请求获取cookie,并且获取验证码
	req, reqErr := http.NewRequest(Type, c.Url, strings.NewReader(bodystr))
	if reqErr != nil {
		fmt.Println("http error %v", reqErr)
	}
	//req.TLS =
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36")

	resP, reqErr := cli.Do(req)
	if reqErr != nil {
		fmt.Println("cli do error %v", reqErr.Error())
	}

	body, bodyErr := ioutil.ReadAll(resP.Body)

	if bodyErr != nil {
		fmt.Println("http error %v", reqErr.Error())
	}
	defer resP.Body.Close()
	return string(body[:])
}

/**
post请求
传入地址，传入参数
type 传入POST,GET
*/
func (c *Util) MethodPayLoad(Type string) string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//建立表单
	var form http.Request
	form.ParseForm()
	for k, v := range c.Params {
		form.Form.Add(k, v)
	}
	bodystr := strings.TrimSpace(form.Form.Encode())
	cli := &http.Client{
		Transport: tr,
		Jar:       c.Jar,
	}
	cli.Timeout = 30 * time.Second
	//fmt.Println("req",Type,Domain + c.Url,strings.NewReader(bodystr),Type)

	//进行第一次请求获取cookie,并且获取验证码
	req, reqErr := http.NewRequest(Type, c.Url, strings.NewReader(bodystr))
	//req, reqErr := http.NewRequest(Type, "http://www.163.com", strings.NewReader(bodystr))
	//fmt.Println("req",req.Response,reqErr)
	defer req.Body.Close()
	if reqErr != nil {
		fmt.Println("http error %v", reqErr)
	}
	//req.TLS =
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//req.Header.Add("Pragma", "no-cache")
	//req.Header.Add("Origin", "http://65b.91cid.com")
	//添加cookie

	resP, reqErr := cli.Do(req)
	if reqErr != nil {
		fmt.Println("cli do error %v", reqErr.Error())
	}

	body, bodyErr := ioutil.ReadAll(resP.Body)

	if bodyErr != nil {
		fmt.Println("http error %v", reqErr.Error())
	}
	defer resP.Body.Close()

	return string(body[:])
}

/**
post请求
传入地址，传入参数
type 传入POST,GET
*/
func (c *Util) Method(Type string) string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//建立表单
	var form http.Request
	form.ParseForm()
	for k, v := range c.Params {
		form.Form.Add(k, v)
	}
	bodystr := strings.TrimSpace(form.Form.Encode())
	cli := &http.Client{
		Transport: tr,
		Jar:       c.Jar,
	}
	cli.Timeout = 30 * time.Second

	//进行第一次请求获取cookie,并且获取验证码
	req, reqErr := http.NewRequest(Type, c.WebSite+c.Url, strings.NewReader(bodystr))
	fmt.Println("c.WebSite+c.Url", c.WebSite+c.Url)
	//req, reqErr := http.NewRequest(Type, "http://www.163.com", strings.NewReader(bodystr))
	//fmt.Println("url",c.Url)
	defer req.Body.Close()
	if reqErr != nil {
		fmt.Println("http error %v", reqErr)
	}
	//req.TLS =
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36")
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//req.Header.Add("Pragma", "no-cache")
	//req.Header.Add("Origin", "http://65b.91cid.com")
	//添加cookie

	resP, reqErr := cli.Do(req)
	if reqErr != nil {
		fmt.Println("cli do error %v", reqErr.Error())
	}

	body, bodyErr := ioutil.ReadAll(resP.Body)

	if bodyErr != nil {
		fmt.Println("http error %v", reqErr.Error())
	}
	defer resP.Body.Close()

	return string(body[:])
}

/**
号码生成
*/
func AlterThree(mode bool) string {
	//生成map
	var initMap []string
	for i := 0; i < 1000; i++ {
		iString := strconv.Itoa(i)
		if len(iString) == 1 {
			iString = "00" + iString
		} else if len(iString) == 2 {
			iString = "0" + iString
		}
		initMap = append(initMap, iString)
	}
	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(100)
	fmt.Println("rand", s)
	//生成一个700-800的随机数
	//确定要删除的个数
	deleteNum := 1000 - (800 - s)
	fmt.Println("deleteNum", deleteNum)
	for i := 0; i < deleteNum; i++ {
		mapLen := len(initMap)
		//生成随机数
		index := rand.Intn(mapLen - 1)
		initMap = append(initMap[:index], initMap[(index+1):]...)
	}
	//fmt.Println("initMap",len(initMap),initMap)

	if mode == true {
		if initMap[len(initMap)-1] != "999" {
			initMap = append(initMap, "999")
		}
	} else {
		initMap = append(initMap[:len(initMap)-1], initMap[len(initMap):]...)
	}
	return strings.Join(initMap, "#")
}
