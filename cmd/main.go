package main

import (
	"encoding/xml"
	"fmt" //
	"io/ioutil"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"` // root element
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"` 
}

type Channel struct {
	Generator   string  `xml:"generator"`
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	Language    string  `xml:"language"`
	Copyright   string  `xml:"copyright"`
	Items       []Item  `xml:"item"` // exchanges
}

type Item struct {
	Title       string  `xml:"title"`       
	PubDate     string  `xml:"pubDate"`     
	Description string  `xml:"description"` 
	Quant       int     `xml:"quant"`       
	Index       string  `xml:"index"`       
	Change      float64 `xml:"change"`      
	Link        string  `xml:"link"`        
}

func main() {
	resp, err := http.Get("https://nationalbank.kz/rss/rates_all.xml")
	if err != nil {
		fmt.Println("response error:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("reading error:", err)
		return
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println("parsing error:", err)
		return
	}

	fmt.Println("title: ", rss.XMLName)
	for _, i:= range rss.Channel.Items {
		fmt.Printf("exhange name: %s, value: %s \n", i.Title, i.Description)
	}

}
