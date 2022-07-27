package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/deuxksy/zzizily-gluttony-go/internal/configuration"
	"github.com/deuxksy/zzizily-gluttony-go/internal/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/input"
	"github.com/spf13/viper"
)

var yyMMddHHmm = time.Now().Local().Format("0601021504")

func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
}

func setRuntimeConfig(profile string) {
	viper.AddConfigPath("configs")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	viper.Set("Verbose", true)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&configuration.RuntimeConf)
	if err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Warn("Config file changed: %s", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			logger.Error(err.Error())
			return
		}
		err = viper.Unmarshal(&configuration.RuntimeConf)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	})
	viper.WatchConfig()
}

func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	// logger.Debug(profile)
	return profile
}

type ChromeVersion struct {
	Browser string `json:"browser"`
	ProtocolVersion string `json:"Protocol-Version"`
	UserAgent string `json:"User-Agent"`
	V8Version string `json:"V8-Version"`
	WebkitVersion string `json:"WebKit-Version"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

func GetWebSocketDebuggerUrl () string {
	lsCmd := exec.Command(
		"/Applications/Google\\ Chrome.app/Contents/MacOS/Google\\ Chrome",
		"--user-data-dir=/Users/crong/TEMP/chrome",
		// "C:/Program Files/Google/Chrome/Application/chrome.exe", 
		// "--user-data-dir=D:/TEMP/chrome", 
		"--remote-debugging-port=12222", 
		"--enable-logging", 
		"--v=1",
	)
	lsCmd.Stdout = os.Stdout
	lsCmd.Start()
	resp, err := http.Get("http://localhost:12222/json/version")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	textBytes := []byte(string(body))
	chromeVersion := ChromeVersion {}
	if err := json.Unmarshal(textBytes, &chromeVersion); err != nil {
		panic(err)
	}
	return chromeVersion.WebSocketDebuggerUrl
}

func checkLimit (page *rod.Page) {
	time.Sleep(time.Millisecond*1000)
	logger.Debug(page.MustInfo().URL)
	page.MustScreenshotFullPage("screenshot/CL01.png")
	logger.Debug("%d", len(page.MustElementsX(`//td[contains(@data-date, "2022-07") and class=".bg-color-blueLight"]`)))
	// for element := range elements {
	// 	logger.Debug("%s", element)
	// 	// MustElement(`div[class="fc-event fc-event-hori fc-event-start fc-event-end bg-color-blueLight"]`).a)
	// }
	os.Exit(1)
}

func initRod() (*rod.Browser) {
	browser := rod.New().MustConnect()
	// url := GetWebSocketDebuggerUrl()
	// browser := rod.New().ControlURL(url).MustConnect()
	browser.DefaultDevice(devices.IPadMini)
	return browser
}

func login (browser *rod.Browser) (*rod.Page) {
	page := browser.MustPage("https://assist9.i-on.net/login")
	logger.Debug(page.MustInfo().URL)
	page.MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "01-login", "01"))
	page.MustWaitLoad().MustElement("input[name=userId]").MustWaitVisible().MustInput(os.Getenv("USERID"))
	page.MustElement("input[name=userPwd]").MustWaitVisible().MustInput(os.Getenv("USERPW"))
	
	page.MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "01-login", "02"))
	page.MustElement("input[name=userPwd]").MustType(input.Enter)//.MustWaitInvisible()
	return page
}

func healthcare (page *rod.Page) {
	page.MustWaitLoad().MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=555a0f1645cee1e334430183")
	logger.Debug(page.MustInfo().URL)
	page.MustWaitLoad().MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "02-healthcare", "01"))

	if page.MustWaitLoad().MustHas(".bg-color-blue") {
		page.MustWaitLoad().MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "02-healthcare", "02"))
		elements := page.MustElements(`div[class="fc-event fc-event-hori fc-event-start fc-event-end bg-color-blue"]`)
		elements.Last().MustClick()
		page.MustWaitLoad().MustElement(`a[class="btn btn-info btn-sm"]`).MustClick()
		logger.Info("%s", "Complate Healthcare")
	} else {
		logger.Warn("%s", "Not Found HealthCare")
	}
	page.MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "02-healthcare", "03"))
}

func lunch (page *rod.Page) {
	page.MustWaitLoad().MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=554971d845ceac19504bbe46")
	logger.Debug(page.MustInfo().URL)
	page.MustWaitLoad().MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "03-lunch", "01"))

	if page.MustWaitLoad().MustHas(".bg-color-blue") {
		page.MustWaitLoad().MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "03-lunch", "02"))
		elements := page.MustElements(`div[class="fc-event fc-event-hori fc-event-start fc-event-end bg-color-blue"]`)
		elements.Last().MustClick()
		page.MustWaitLoad().MustElement(`a[class="btn btn-info btn-sm"]`).MustClick()
		logger.Info("%s", "Complate Lunch")
		} else {
			logger.Warn("%s", "Not Found Lunch")
		}
		page.MustScreenshotFullPage(fmt.Sprintf("screenshot/%s/%s-%s.png", yyMMddHHmm, "03-lunch", "03"))
	}

func GetLimit(browser *rod.Browser) (int, int) {
	logger.Debug("%s", browser.MustVersion().UserAgent)
	cookies, err := browser.GetCookies()
	if err != nil {
		panic(err)
	}
	for _, cookie := range cookies {
		logger.Debug("%s: %s=%s", cookie.Domain, cookie.Name, cookie.Value)
	}
	return 1, 1
}

func GetUserAgent (page *rod.Page) string {
	if (page == nil) {
		logger.Error("%s", "Page is NULL")
		os.Exit(1)
	}
	// page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
	// 	UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
	// 	AcceptLanguage: "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7",
	// 	Platform: "Windows NT",
	// 	UserAgentMetadata: &proto.EmulationUserAgentMetadata{
	// 	},
	// })
	return page.MustEval(`()=>window.navigator.userAgent`).String()
}

func main () {
	logger.Info("%s", "Gluttony")
	browser := initRod()
	defer browser.MustClose()
	
	page := login(browser)
	page.Eval(`window.alert = () => {}`)
	healthcare(page)
	lunch(page)
}