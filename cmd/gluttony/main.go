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

func initChrome () string {
	lsCmd := exec.Command(
		"C:/Program Files/Google/Chrome/Application/chrome.exe", 
		"--user-data-dir=D:/TEMP/chrome", 
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

func a9 () {
	url := initChrome()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	tap := browser.MustPage("https://assist9.i-on.net/login")
	
	tap.MustScreenshot("screenshot/01.png")
	tap.MustElement("input[name=userId]").MustWaitVisible().MustInput(os.Getenv("USERNAME"))
	tap.MustElement("input[name=userPwd]").MustWaitVisible().MustInput(os.Getenv("USERPASSWORD"))
	
	time.Sleep(time.Millisecond*500)
	tap.MustScreenshot("screenshot/02.png")
	tap.MustElement("input[name=userPwd]").MustType(input.Enter)
	
	time.Sleep(time.Millisecond*500)
	tap.MustScreenshot("screenshot/03.png")
	logger.Debug(tap.MustInfo().URL)
	tap.MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=554971d845ceac19504bbe46")
	
	time.Sleep(time.Millisecond*500)
	tap.MustScreenshot("screenshot/05.png")
	// tap.MustElement("div[class=`fc-event fc-event-hori fc-event-start fc-event-end bg-color-blue`]")
	
	// time.Sleep(time.Millisecond*500)
	// browser.MustScreenshot("screenshot/06.png")
	// res := browser.MustElementR("a", "chromedp").MustParent().MustParent().MustNext().MustText()
	// log.Printf("got: "%s"", strings.TrimSpace(res))
}


func Example_wait_for_request() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://duckduckgo.com/")
	page.MustScreenshot("screenshot/duckduckgo/01.png")
	// Start to analyze request events
	wait := page.MustWaitRequestIdle()

	// This will trigger the search ajax request
	page.MustElement("#search_form_input_homepage").MustClick().MustInput("lisp")

	// Wait until there's no active requests
	wait()

	// We want to make sure that after waiting, there are some autocomplete
	// suggestions available.
	fmt.Println(len(page.MustElements(".search__autocomplete .acp")) > 0)

	// Output: true
}

func main () {
	a9()
}



