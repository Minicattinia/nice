package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {

	colorGreen := "\033[32m"
	getURL := "https://cryptofy.club/home.php?LTC=1&DOGE=1&DGB=1&TRX=1&USDT=1&BCH=1&DASH=1&FEY=1&ZEC=1&SOL=1&redirect=1"
	data := url.Values{}
	getHeader := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPost, getURL, strings.NewReader(data.Encode()))
	request.Header.Add("User-Agent", getHeader)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", "PHPSESSID=d8230d64aba3447e0157de1bd077fe43; _immortal|Arc_nodeId=JvxN5ac7gBV9THPxnDxbWj; widgetOptState={%22state%22:%22UNDECIDED%22%2C%22date%22:%222022-11-04T08:16:20.606Z%22%2C%22dismissedAt%22:null}")
	for {
		resp, _ := client.Do(request)
		if resp.StatusCode == 200 {
			fmt.Println(string(colorGreen), "Success Claim")
		}
		time.Sleep(60 * time.Second)
	}

}
