package main

import (
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

func main () {
	page := rod.New().MustConnect().MustPage("https://assist9.i-on.net/login")
	page.MustScreenshot("screen/01.png")
	page.MustElement("input[name=userId]").MustWaitVisible().MustInput("")
	page.MustElement("input[name=userPwd]").MustWaitVisible().MustInput("")
	page.MustScreenshot("screen/02.png")
	page.MustElement("input[name=userPwd]").MustType(input.Enter)
	page.MustScreenshot("screen/03.png")
	time.Sleep(time.Millisecond*500)
	logger.Debug(page.MustInfo().URL)
	page.MustScreenshot("screen/04.png")
	time.Sleep(time.Millisecond*500)
	page.MustNavigate("https://assist9.i-on.net/rb/main#booking/calendar?resourceId=554971d845ceac19504bbe46")
	time.Sleep(time.Millisecond*500)
	page.MustScreenshot("screen/05.png")
	page.MustElementX("//*[contains(@class, 'fc-event fc-event-hori fc-event-start fc-event-end bg-color-red')]").Click()
	time.Sleep(time.Millisecond*500)
	page.MustScreenshot("screen/06.png")
	// res := page.MustElementR("a", "chromedp").MustParent().MustParent().MustNext().MustText()
	// log.Printf("got: "%s"", strings.TrimSpace(res))

}