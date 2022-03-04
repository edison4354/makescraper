package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

type valTeam struct {
	Name string
	Rank int
}

func main() {
	var valRanking = []valTeam{}

	var allTeams []string

	// Instantiate default collector
	c := colly.NewCollector()

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a.rank-item-team.fc-flex", func(e *colly.HTMLElement) {
		teams := strings.TrimSpace(e.Attr("data-sort-value"))
		allTeams = append(allTeams, teams)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.vlr.gg/rankings/north-america")

	for i := 0; i < 5; i++ {
		team := valTeam{Name: allTeams[i], Rank: i + 1}
		valRanking = append(valRanking, team)
	}

	valTeamJSON, _ := json.Marshal(valRanking)
	err := ioutil.WriteFile("output.json", valTeamJSON, 0644)
	if err != nil {
		panic(err)
	}
}
