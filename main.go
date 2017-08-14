package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// BittrexCurrency market
type BittrexCurrency struct {
	Volume         float64 `json:"BaseVolume"`
	Bid            float64 `json:"Bid"`
	Ask            float64 `json:"Ask"`
	Last           float64 `json:"Last"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	MarketName     string  `json:"MarketName"`
	OpenBuyOrders  int     `json:"OpenBuyOrders"`
	OpenSellOrders int     `json:"OpenSellOrders"`
}

// BittrexCurrencies Pair
type BittrexCurrencies struct {
	Result []BittrexCurrency
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	log.Println("Server is listening at port 8080")

	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8080", nil)

}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Ajout d'un utilisateur au socket ")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	go echo(conn)
}

func echo(conn *websocket.Conn) {
	url := "https://bittrex.com/api/v1.1/public/getmarketsummaries"

	ticker := time.NewTicker(time.Second * time.Duration(5))

	for range ticker.C {

		var record BittrexCurrencies

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error get currency : ", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &record)

		conn.WriteJSON(record)
	}
}
