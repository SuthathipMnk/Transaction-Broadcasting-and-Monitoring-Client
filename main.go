package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Payload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type ServerPostResponse struct {
	TxHash string `json:"tx_hash"`
}

type ServerGetResponse struct {
	TxStatus string `json:"tx_status"`
}

func main() {
	payload := &Payload{Symbol: "ETH", Price: 4500, Timestamp: 1678912345}

	serverPostResponse := getTxHash(*payload)
	serverGetResponse := getTxStatus(*serverPostResponse)

	fmt.Println("tx_status:", serverGetResponse.TxStatus)
}

func getTxHash(payload Payload) *ServerPostResponse {
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	posturl := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"

	responseBody := bytes.NewBuffer(b)

	resp, err := http.Post(posturl, "application/json", responseBody)

	if err != nil {
		fmt.Printf("An Error Occured %v \n", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	serverPostResponse := &ServerPostResponse{}
	json.NewDecoder(resp.Body).Decode(serverPostResponse)

	return serverPostResponse
}

func getTxStatus(serverPostResponse ServerPostResponse) *ServerGetResponse {
	path := fmt.Sprintf("https://mock-node-wgqbnxruha-as.a.run.app/check/%s", serverPostResponse.TxHash)
	resp2, err2 := http.Get(path)

	if err2 != nil {
		fmt.Printf("An Error Occured %v \n", err2)
		os.Exit(0)
	}

	defer resp2.Body.Close()

	serverGetResponse := &ServerGetResponse{}
	json.NewDecoder(resp2.Body).Decode(serverGetResponse)

	return serverGetResponse
}
