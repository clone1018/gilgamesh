package main

import (
	"fmt"
	_ "io/ioutil"
	"log"
	_ "os"
	"runtime"
	_ "time"

	"github.com/clone1018/gilgamesh/loader"
)

var (
	config = &Configuration{}
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := config.load(DefaultConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	for bot, bconfig := range config.Discord_Bot {
		log.Println("Loading bot: ", bot)

		for _, plugin := range bconfig.Plugin {

			log.Println("Loading plugin: ", plugin)

			pluginFolder := fmt.Sprintf("plugins/%s", plugin)

			loader := loader.JsLoader{Dir: pluginFolder}
			err := loader.Load()
			if err != nil {
				log.Fatal(err)
			}
			loader.Run()
		}

	}
}
