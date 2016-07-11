package main

import (
	"fmt"
	"encoding/json"
)

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
	sum    Sum `json:"sum"`
	groups []AdGroup `json:"ad_groups"`
}

func main() {
	sum := Sum {
		Cost: 3,
		Imp: 30,
		Click: 10,
	}

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

	adGroups := AdGroups {
		sum: sum,
		groups: groups,
	}
	fmt.Printf("adGroups: %+v", adGroups)
	fmt.Println("")
	fmt.Println("")

	b, err := json.Marshal(adGroups)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", string(b))
}
