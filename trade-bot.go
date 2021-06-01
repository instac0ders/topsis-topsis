package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/shopspring/decimal"
	"github.com/toorop/go-bittrex"
	"time"
)

var config, err = toml.LoadFile("trade-bot.conf")
var configTree = config.Get("config").(*toml.Tree)
var base_coin = configTree.Get("base_coin").(string)
var market_coin = configTree.Get("market_coin").(string)
v