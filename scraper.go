package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/json"
	"os"
)

func main() {
	urls := []string{
		"https://www.npsl.com/team/chattanooga-fc/roster/",
		"https://www.npsl.com/team/detroit-city-fc/roster/",
		"https://www.npsl.com/team/michigan-stars-fc/roster/",
		"https://www.npsl.com/team/milwaukee-torrent/roster/",
		"https://www.npsl.com/team/napa-valley-1839-fc/roster/",
		"https://www.npsl.com/team/new-york-cosmos-b/roster/",
	}
	type Player struct {
		Name	string `json:"name"`
		SquadNumber	string `json:"squadNumber"`
		Position	string `json:"position"`
	}
	type Foe struct {
		Opponent	string `json:"opponent"`
		Players		[]Player `json:"players"`
	}
	foes := []Foe {
		Foe {
			Opponent: "Chattanooga FC",
			Players: []Player{},
		},
		Foe {
			Opponent: "Detroit City FC",
			Players: []Player{},
		},
		Foe {
			Opponent: "Michigan Stars FC",
			Players: []Player{},
		},
		Foe {
			Opponent: "Milwaukee Torrent",
			Players: []Player{},
		},
		Foe {
			Opponent: "Napa Valley 1839 FC",
			Players: []Player{},
		},
		Foe {
			Opponent: "New York Cosmos B",
			Players: []Player{},
		},
	}
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.npsl.com"),
	)
	i := 0
	url := ""

	// On every a element which has href attribute call callback
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		player := Player{}
		player.SquadNumber = e.ChildText(".data-number")
		if player.SquadNumber == "#" {
			return
		}
		player.Name = e.ChildText(".data-name")
		switch position := e.ChildText(".data-position"); position {
			case "Goalkeeper":
				player.Position = "GK"
			case "Midfielder":
				player.Position = "M"
			case "Forward":
				player.Position = "F"
			case "Defender":
				player.Position = "D"
		}
		
		foes[i].Players = append(foes[i].Players, player)
		//fmt.Println(foes[i].players)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for i, url = range urls {
		c.Visit(url)
	}

	f, _ := os.Create("foes.json")
	defer f.Close()
	
	b, err := json.Marshal(foes)
	if err != nil {
		fmt.Println("error:", err)
	}
	f.Write(b)
	f.Sync()
}