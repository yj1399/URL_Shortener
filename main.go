
package main


import (
	"flag"
	"fmt"
	"os"

	"github.com/yj1399/shortme/conf"
	"github.com/yj1399/shortme/short"
	"github.com/yj1399/shortme/web"
)

func main() {
	cfgFile := flag.String("c", "config.conf", "configuration file")
	version := flag.Bool("v", false, "Version")

	flag.Parse()

	if *version {
		fmt.Println(conf.Version)
		os.Exit(0)
	}

	// parse config
	conf.MustParseConfig(*cfgFile)

	// short service
	short.Start()

	// api
	web.Start()
}
