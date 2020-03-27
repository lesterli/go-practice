package main

import (
	"github.com/lesterli/go-practice/play-cosmos/config"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "cosmos sync config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

}
