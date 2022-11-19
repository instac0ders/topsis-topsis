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
var last_sell = configTree.Get("last_sell").(float64)
var target_sell = configTree.Get("target_sell").(float64)

// Inititalize the Bittrex client using credentials found in loaded config file
var bittrex_client = bittrex.New(api_key, api_secret)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func get_balance() string {
	base_coin_balance, _ := bittrex_client.GetBalance(base_coin)
	market_coin_balance, _ := bittrex_client.GetBalance(market_coin)
	return fmt.Sprintf("[Balances: %v %v / %v %v]",
		base_coin_balance.Available,
		base_coin,
		market_coin_balance.Available,
		market_coin,
	)
}

func last_price(t time.Time) {
	// Using decimal package for accuracy
	ticker, _ := bittrex_client.GetTicker(base_coin + "-" + market_coin)
	last_sell := decimal.NewFromFloat(last_sell)
	last_bid := ticker.Bid
	net_diff := fmt.Sprint(((last_bid.Sub(last_sell)).Div(last_sell)).Mul(decimal.NewFromFloat(100)))

	fmt.Printf("[+] Last Bid: %v || Last Ask: %v || Last market sell: %v [ My last sell: %v Net diff: %v ] %v\n",
		ticker.Bid,
		ticker.Ask,
		ticker.Last,
		last_sell,
		net_diff,
		