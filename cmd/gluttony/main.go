package main

import (
	"encoding/json"
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
		"/usr/bin/google-chrome",
		"--user-data-dir=/home/crom/temp/chrome",
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

func lunch (page *rod.Page) {	
	time.Sleep(time.Millisecond*500)
	//page.MustWaitLoad()
	
	page.MustScreenshot("screenshot/C01.png")
	logger.Debug(page.MustInfo().URL)
	wait := page.MustWaitNavigation()
	page.MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=554971d845ceac19504bbe46")
	wait()
	checkLimit(page)
	page.MustScreenshot("screenshot/C02.png")
	if page.MustHas(".bg-color-blue") {
		page.MustElement(`div[class="fc-event fc-event-hori fc-event-start fc-event-end bg-color-blue"]`).MustClick()
		time.Sleep(time.Millisecond*500)
		page.MustScreenshot("screenshot/C03.png")
		page.MustElement(`a[class="btn btn-info btn-sm"]`).MustClick()
		logger.Info("%s", "점심식사 신청을 완료 하였습니다.")
	} else {
		logger.Warn("%s", "신청할 점심식사가 없습니다.")
	}
	time.Sleep(time.Millisecond*500)
	page.MustScreenshot("screenshot/C04.png")
}

func healthcare (page *rod.Page) {
	page.MustScreenshotFullPage("screenshot/healthcare-01.png")
	logger.Debug(page.MustInfo().URL)
	wait := page.MustWaitNavigation()
	page.MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=555a0f1645cee1e334430183")
	wait()

	page.MustScreenshotFullPage("screenshot/healthcare-02.png")
	elements := page.MustElements(`div[class="fc-event fc-event-hori fc-event-start fc-event-end bg-color-blue"]`)
	if page.MustHas(".bg-color-blue") {
		elements.Last().MustClick()
		time.Sleep(time.Millisecond*500)
		page.MustScreenshotFullPage("screenshot/healthcare-03.png")
		page.MustElement(`a[class="btn btn-info btn-sm"]`).MustClick()
		logger.Info("%s", "헬스케어 신청을 완료 하였습니다.")
	} else {
		logger.Warn("%s", "신청할 헬스케어가 없습니다.")
	}
	time.Sleep(time.Millisecond*500)
	page.MustScreenshotFullPage("screenshot/healthcare-04.png")
}

func initRod() (*rod.Browser, *rod.Page) {
	browser := rod.New().MustConnect()
	browser.DefaultDevice(devices.IPadMini)
	// url := GetWebSocketDebuggerUrl()
	// browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage()
	return browser, page
}

func login (browser *rod.Browser) (*rod.Page) {
	page := browser.MustPage("https://assist9.i-on.net/login")
	page.MustScreenshot("screenshot/login-01.png")
	page.MustElement("input[name=userId]").MustWaitVisible().MustInput(os.Getenv("USERID"))
	page.MustElement("input[name=userPwd]").MustWaitVisible().MustInput(os.Getenv("USERPW"))
	
	page.MustScreenshot("screenshot/login-02.png")
	page.MustElement("input[name=userPwd]").MustType(input.Enter)//.MustWaitInvisible()
	return page
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
	logger.Debug("%s", "Gluttony")
	browser, page := initRod()
	defer browser.MustClose()

	page = login(browser)
	healthcare(page)
	// lunch(page)
}