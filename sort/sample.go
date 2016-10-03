package main

import (
	"sort"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Ad struct {
	AdID        string     `json:"ad_id"`
	Title       string     `json:"title"`
	Cost        null.Int   `json:"cost"`
	Imp         null.Int   `json:"imp"`
}

type Ads []Ad

var order map [string]string = map [string]string {
	"column": "imp",
	"sort": "asc",
}

func init() {
}

func main() {
	var ads Ads = []Ad {
		{
			AdID: "394839",
			Title: "タイトル１",
			Cost: null.IntFrom(80),
			Imp: null.IntFrom(20),
		},
		{
			AdID: "194839",
			Title: "タイトル２",
			Cost: null.IntFrom(50),
			Imp: null.IntFrom(10),
		},
		{
			AdID: "994839",
			Title: "タイトル３",
			Cost: null.IntFrom(400),
			Imp: null.IntFrom(1),
		},
	}

	fmt.Printf("before: %+v", ads)
	sort.Sort(ads)
	fmt.Println("")
	fmt.Println("======")
	fmt.Print(ads)

}

func (ads Ads) Len() int {
	return len(ads)
}

func (ads Ads) Swap(i, j int) {
	ads[i], ads[j] = ads[j], ads[i]
}

func (ads Ads) Less(i, j int) bool {
	fmt.Printf("order ::: %+v", order)
	fmt.Println("======")

	switch(order["column"]) {
	case "cost":
		fmt.Println("ORDER cost")
		return ads.LessCost(i, j)
	case "imp":
		fmt.Println("ORDER imp")
		return ads.LessImp(i, j)
	default:
		fmt.Println("ORDER DEFAULT")
		return ads.LessCost(i, j)
	}


}

func (ads Ads) LessCost(i, j int) bool {
	if order["sort"] == "asc" {
		return ads[i].Cost.Int64 < ads[j].Cost.Int64
	} else {
		return ads[i].Cost.Int64 > ads[j].Cost.Int64
	}
}

func (ads Ads) LessImp(i, j int) bool {
	if order["sort"] == "asc" {
		return ads[i].Imp.Int64 < ads[j].Imp.Int64
	} else {
		return ads[i].Imp.Int64 > ads[j].Imp.Int64
	}
}
