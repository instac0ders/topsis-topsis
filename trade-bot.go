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
var api_version = configTree.Get("api_version").(string)
var base_url = configTree.Get("base_url").(string)
var api_key = configTree.Get("api_key").(string)
var api_secret = configTree.Get("api_secret").(string)
var last_sell = configTree.Get("last_sell").(float6