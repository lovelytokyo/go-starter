package main

import (
	"encoding/json"
	//test."./test"
)
import "fmt"

type Sum struct {
	Cost   uint64  `json:"cost"`
	Imp    uint64  `json:"imp"`
	Click  uint64  `json:"click"`
}

type AdGroup struct {
	AdGroupID   string  `json:"ad_group_id"`
	AdGroupName string  `json:"ad_group_name"`
}

type AdGroups struct {
	*Sum `json:"sum"`
	AdGroups []AdGroup `json:"ad_groups"`
}

func main() {
	sum := new(Sum)
	sum.Cost = 3
	sum.Imp = 30
	sum.Click = 10

	// ----- sum
	fmt.Printf("sum: %T (%+v)", sum, sum)
	fmt.Println("")
	fmt.Println("")

	b, err := json.Marshal(sum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", string(b))

	fmt.Println("")
	fmt.Println("=========================")

	// ----- groups
	groups := []AdGroup {
		AdGroup {
			AdGroupID: "1111",
			AdGroupName: "test1",
		},
		AdGroup {
			AdGroupID: "2222",
			AdGroupName: "test2",
		},

	}
	fmt.Printf("groups: %T (%+v)", groups, groups)
	fmt.Println("")
	fmt.Println("")

	c, err := json.Marshal(groups)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", string(c))

	fmt.Println("")
	fmt.Println("=========================")

	// ----- adgroups
	adGroups := AdGroups{
		sum,
		groups,
	}

	fmt.Printf("adGroups: %T (%+v)", adGroups, adGroups)
	fmt.Println("")
	fmt.Println("")

	d, err := json.Marshal(adGroups)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", string(d))



}