package main

import (
	"flag"
	"fmt"
	"go-simple-parse-ini/parseconf"
)

func main() {

	conf := flag.String("config path", "src/go-simple-parse-ini/example_conf.ini", "usage ini type file")

	flag.Parse()

	parseRst := parseconf.InitialConf(conf)

	if parseRst != nil {
		fmt.Println(parseRst.Error())
	}

}
