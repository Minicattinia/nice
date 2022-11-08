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
	cookie := "PHPSESSID=5958157c25ce0187d81746fa4ac1e87c; _immortal|Arc_nodeId=RzC88WkaJf2TH7R4L6tZeE; widgetOptState={%22state%22:%22UNDECIDED%22%2C%22date%22:%222022-10-22T07:09:14.153Z%22%2C%22dismissedAt%22:null}; _ga=GA1.2.1160819856.1666422601; _gid=GA1.2.1177212240.1667878780; _gat_gtag_UA_193127892_1=1"
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
