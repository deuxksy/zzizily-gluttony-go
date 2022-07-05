package main

import (
	"log"
	"os"
	"strings"

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
	logger.Debug(profile)
	return profile
}

func main () {
	page := rod.New().MustConnect().MustPage("https://github.com/search")
	page.MustScreenshot("screen/search.png")
	page.MustElement(`input[name=q]`).MustWaitVisible().MustInput("chromedp").MustType(input.Enter)
	page.MustScreenshot("screen/input.png")
	res := page.MustElementR("a", "chromedp").MustParent().MustParent().MustNext().MustText()
	log.Printf("got: `%s`", strings.TrimSpace(res))
}