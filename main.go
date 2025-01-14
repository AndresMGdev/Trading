package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Trade struct {
	Symbol    string `json:"symbol"`
	TradeID   string `json:"trade_id"`
	Timestamp int64  `json:"timestamp"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Price     string `json:"price"`
}

type Message struct {
	Topic  string  `json:"topic"`
	Symbol string  `json:"symbol"`
	Data   []Trade `json:"data"`
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	symbols = []string{"BTC_USDT", "ETH_USDT", "BNB_USDT", "ADA_USDT", "XRP_USDT", "DOGE_USDT", "LINK_USDT"}

	lastPrice = 50000.0
)

func generateMockTrade(timestamp int64) Trade {
	priceChange := rnd.Float64() *20 - 10
	lastPrice += priceChange

	return Trade{
		Symbol: symbols[rnd.Intn(len(symbols))],
		TradeID: strconv.Itoa(rnd.Intn(1000000)),
		Price: strconv.FormatFloat(lastPrice, 'f', 2, 64),
		Size: strconv.FormatFloat(0.01 + rnd.Float64() * 0.1, 'f', 4, 64),
		Side: []string{"BUY", "SELL"}[rnd.Intn(2)],
		Timestamp: timestamp,
	}
}

func tradeStreamHandler(c *gin.Context) {

}

func main() {

}
