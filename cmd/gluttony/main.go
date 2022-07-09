package main

import (
	"fmt"
	"os"
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

func a9 () {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	tap := browser.MustPage("https://assist9.i-on.net/login")
	
	tap.MustScreenshot("screen/01.png")
	tap.MustElement("input[name=userId]").MustWaitVisible().MustInput("")
	tap.MustElement("input[name=userPwd]").MustWaitVisible().MustInput("")
	
	time.Sleep(time.Millisecond*500)
	tap.MustScreenshot("screen/02.png")
	tap.MustElement("input[name=userPwd]").MustType(input.Enter)
	
	time.Sleep(time.Millisecond*500)
	tap.MustScreenshot("screen/03.png")
	logger.Debug(tap.MustInfo().URL)
	tap.MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=554971d845ceac19504bbe46")
	
	time.Sleep(time.Millisecond*500)
	tap.MustScreenshot("screen/05.png")
	tap.MustElement("div[class=`fc-event fc-event-hori fc-event-start fc-event-end bg-color-blue`]")
	
	// time.Sleep(time.Millisecond*500)
	// browser.MustScreenshot("screen/06.png")
	// res := browser.MustElementR("a", "chromedp").MustParent().MustParent().MustNext().MustText()
	// log.Printf("got: "%s"", strings.TrimSpace(res))
}


func Example_wait_for_request() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://duckduckgo.com/")
	page.MustScreenshot("screen/duckduckgo/01.png")
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
	Example_wait_for_request()
}