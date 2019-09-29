package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/json"
	"os"
	"strings"
)

func main() {
	urls := []string{
		"https://www.uwssoccer.com/roster/show/4813346?subseason=590879", //ann arbor
		"https://www.uwssoccer.com/roster/show/4813292?subseason=590879", //detroit sun
		"https://www.uwssoccer.com/roster/show/4813296?subseason=590879", //grand rapids fc
		"https://www.uwssoccer.com/roster/show/4813297?subseason=590879", //indiana union
		"https://www.uwssoccer.com/roster/show/4813298?subseason=590879", //lansing united
		"https://www.uwssoccer.com/roster/show/4813299?subseason=590879", //michigan legends
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
	foes := make(map[string]*Foe)
	foes["AFC Ann Arbor"] = &Foe {
		Opponent: "AFC Ann Arbor",
		Players: []Player{},
	}
	foes["Detroit Sun"] = &Foe {
		Opponent: "Detroit Sun",
		Players: []Player{},
	}
	foes["Grand Rapids FC"] = &Foe {
		Opponent: "Grand Rapids FC",
		Players: []Player{},
	}
	foes["Indiana Union"] = &Foe {
		Opponent: "Indiana Union",
		Players: []Player{},
	}
	foes["Lansing United"] = &Foe {
		Opponent: "Lansing United",
		Players: []Player{},
	}
	foes["Michigan Legends"] = &Foe {
		Opponent: "Michigan Legends",
		Players: []Player{},
	}

	// Instantiate default collector
	rosterCollector := colly.NewCollector(
		colly.AllowedDomains("www.uwssoccer.com"),
	)

	playerCollector := colly.NewCollector(
		colly.AllowedDomains("www.uwssoccer.com"),
	)

	// On every a element which has href attribute call callback
	rosterCollector.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach(".name > a[href]", func(_ int, el *colly.HTMLElement) { 
			playerCollector.Visit(el.Attr("href"))
		})
	})

	playerCollector.OnHTML("body", func(e *colly.HTMLElement) {
		playerData:= e.ChildText(".playerName")
		rosterName:= e.ChildText("h2")
		dataPieces:= strings.Split(playerData, "\n")
		player:= Player{}
		player.Name = dataPieces[0];
		if(len(dataPieces) > 1) {
			positionInfo := strings.Split(dataPieces[1], "·")
			player.SquadNumber = strings.Trim(positionInfo[0], " #")
			if(len(positionInfo) > 1) {
				player.Position = positionInfo[1]
			} else {
				player.Position = ""
			}
		} else {
			player.SquadNumber = "0"
			player.Position = "";
		}
		foes[rosterName].Players = append(foes[rosterName].Players, player)
	})

	// Before making a request print "Visiting ..."
	rosterCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting roster", r.URL.String())
	})

	playerCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting player", r.URL.String())
	})

	for _, url := range urls {
		rosterCollector.Visit(url)
	}

	f, _ := os.Create("uws-foes.json")
	defer f.Close()

	foesList := make([]Foe, len(foes))
	idx := 0
    for  _, value := range foes {
       foesList[idx] = *value
       idx++
    }
	b, err := json.Marshal(foesList)
	if err != nil {
		fmt.Println("error:", err)
	}
	f.Write(b)
	f.Sync()
}