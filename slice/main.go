package main

import "fmt"

type AdGroup struct {
	AdGroupID   string  `json:"ad_group_id"`
	AdGroupName string  `json:"ad_group_name"`
}

func sliceToMapByAdGroupID(b []AdGroup) map[string]AdGroup {
	r := map[string]AdGroup{}

	if len(b) < 1 {
		return r
	}

	for _, v := range b {
		r[v.AdGroupID] = v
	}
	return r
}


func main() {
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


	remap := sliceToMapByAdGroupID(groups)
	fmt.Printf("%+v", remap)
}
