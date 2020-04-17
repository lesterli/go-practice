package main

import (
	"github.com/lesterli/go-practice/play-cosmos/logger"
	"github.com/lesterli/go-practice/play-cosmos/store"
	"github.com/lesterli/go-practice/play-cosmos/store/document"

	"github.com/lesterli/go-practice/play-cosmos/config"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "cosmos sync config file path.")
)

func main() {
	log := logger.GetLogger("main")
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//start databases service
	log.Info("Databases Service Start...")
	store.Start()

	start := int64(10001)
	end := int64(100000)
	for i := start; i <= end; i++ {
		var blockModel document.Block
		dbBlockHeight, err := blockModel.GetBlock(i)
		if err != nil {
			log.Error("get db height fail", logger.String("err", err.Error()))
		}
		log.Info("get db", logger.Int64("block", dbBlockHeight))
	}

}
