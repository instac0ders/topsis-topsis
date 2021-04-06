package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/shopspring/decimal"
	"github.com/toorop/go-bittrex"
	"time"
)

var config, err = toml.LoadFile("t