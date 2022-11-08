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
	cookie := "PHPSESSID=9bf0c1fb842a76b8914a0da2b4293f13; _immortal|Arc_nodeId=YFbrRdDoJHPU8htBtRTyfe; widgetOptState={%22state%22:%22UNDECIDED%22%2C%22date%22:%222022-11-08T03:16:13.186Z%22%2C%22dismissedAt%22:null}"
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPost, getURL, strings.NewReader(data.Encode()))
	request.Header.Add("User-Agent", getHeader)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", cookie)
	for {
		resp, _ := client.Do(request)
		if resp.StatusCode == 200 {
			fmt.Println(string(colorGreen), "Success Claim")
		}
		time.Sleep(60 * time.Second)
	}

}
