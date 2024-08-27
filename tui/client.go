package tui

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"
)

type Client struct {
	Hosts []string // 服务器地址s
	host  string   // 检查后的服务器地址
}

func (c *Client) SetProxy(lang string) {
	defer c.setHost()
	proxy := httplib.BeegoHTTPSettings{}.Proxy
	proxyText := ""
	if os.Getenv("http_proxy") != "" {
		proxy = func(request *http.Request) (*url.URL, error) {
			return url.Parse(os.Getenv("http_proxy"))
		}
		proxyText = os.Getenv("http_proxy") + " " + Trr.Tr("经由") + " http_proxy " + Trr.Tr("代理访问")
	}
	if os.Getenv("https_proxy") != "" {
		proxy = func(request *http.Request) (*url.URL, error) {
			return url.Parse(os.Getenv("https_proxy"))
		}
		proxyText = os.Getenv("https_proxy") + " " + Trr.Tr("经由") + " https_proxy " + Trr.Tr("代理访问")
	}
	if os.Getenv("all_proxy") != "" {
		proxy = func(request *http.Request) (*url.URL, error) {
			return url.Parse(os.Getenv("all_proxy"))
		}
		proxyText = os.Getenv("all_proxy") + " " + Trr.Tr("经由") + " all_proxy " + Trr.Tr("代理访问")
	}
	httplib.SetDefaultSetting(httplib.BeegoHTTPSettings{
		Proxy:            proxy,
		ReadWriteTimeout: 30 * time.Second,
		ConnectTimeout:   30 * time.Second,
		Gzip:             true,
		DumpBody:         true,
		UserAgent:        fmt.Sprintf(`{"lang":"%s","GOOS":"%s","ARCH":"%s","version":%d,"deviceID":"%s"}`, lang, runtime.GOOS, runtime.GOARCH, version, deviceID),
	})
	if len(proxyText) > 0 {
		fmt.Printf(yellow, proxyText)
	}
}

func (c *Client) setHost() {
	for _, v := range c.Hosts {
		_, err := httplib.Get(v).SetTimeout(4*time.Second, 4*time.Second).String()
		if err == nil {
			c.host = v
			return
		}
	}
	return
}

func (c *Client) GetAD() (ad string) {
	res, err := httplib.Get(c.host + "/ad").String()
	if err != nil {
		return
	}
	return res
}

func (c *Client) GetPayUrl() (payUrl, orderID string) {
	res, err := httplib.Get(c.host + "/payUrl").String()
	if err != nil {
		fmt.Println(err)
		return
	}
	payUrl = gjson.Get(res, "payUrl").String()
	orderID = gjson.Get(res, "orderID").String()
	return
}
func (c *Client) PayCheck(orderID, deviceID string) (isPay bool) {
	res, err := httplib.Get(c.host + "/payCheck?orderID=" + orderID + "&deviceID=" + deviceID).String()
	if err != nil {
		fmt.Println(err)
		return
	}
	isPay = gjson.Get(res, "isPay").Bool()
	return
}

func (c *Client) GetMyInfo(deviceID string) (sCount, sPayCount, isPay, ticket, exp string) {
	body, _ := json.Marshal(map[string]string{
		"device":  deviceID,
		"sDevice": getPromotion(),
	})
	res, err := httplib.Post(host + "/my").Body(body).String()
	if err != nil {
		return
	}
	sCount = gjson.Get(res, "sCount").String()
	sPayCount = gjson.Get(res, "sPayCount").String()
	isPay = gjson.Get(res, "isPay").String()
	ticket = gjson.Get(res, "ticket").String()
	exp = gjson.Get(res, "exp").String()
	return
}

func (c *Client) CheckVersion(version string) (upUrl string) {
	res, err := httplib.Get(host + "/version?version=" + version + "&plat=" + runtime.GOOS + "_" + runtime.GOARCH).String()
	if err != nil {
		return ""
	}
	upUrl = gjson.Get(res, "url").String()
	return
}

func (c *Client) GetLic() (isOk bool, result string) {
	req := httplib.Get(host + "/getLic?device=" + getMacMD5())
	res, err := req.String()
	if err != nil {
		isOk = false
		result = err.Error()
		return
	}
	code := gjson.Get(res, "code").Int()
	msg := gjson.Get(res, "msg").String()
	result = msg
	if code != 0 {
		isOk = false
		return
	}
	isOk = true
	return
}
