package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.detcityfc.com"),
	)
	type playerData struct {
		Name string
		Url string
	}
	players := []playerData{}
	// On every a element which has href attribute call callback
	c.OnHTML(".textBlockElement", func(e *colly.HTMLElement) {
		playerName := e.ChildText("h3")
		uri := e.ChildAttr("img", "src")
		if(strings.Contains(playerName, "ROSTER")) {
			return
		}
		player := playerData{Name: playerName, Url: uri}
		players = append(players, player)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.detcityfc.com/teamroster")
	for _, data := range players {
		resp, err := http.Get(data.Url)
		if(err != nil) {
			fmt.Println("ERROR RETRIEVING ", data.Name)
			break;
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if(err != nil) {
			fmt.Println("ERROR RETRIEVING ", data.Name)
			break;
		}
		filename:= fmt.Sprintf("%s.jpg", data.Name)
		f, _ := os.Create(filename)
		defer f.Close()
		f.Write(body)
		f.Sync()
	}
}