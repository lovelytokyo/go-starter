package main

import (
	"github.com/Unknwon/paginater"
	"fmt"
)

func generateLinks(p *paginater.Paginater) string {
	fmt.Printf("current: %d\n",p.Current())
	return fmt.Sprintf("current page: %d", p.Current())
}

func main() {
	// Arguments:
	// - Total number of rows
	// - Number of rows in one page
	// - Current page number
	// - Number of page links
	//totalRowNum := 45
	//numberOfOnePage := 10
	//currentPage := 5
	//numberOfPageLinks := 3
	totalRowNum := 4
	numberOfOnePage := 1
	currentPage := 3
	numberOfPageLinks := 0

	p := paginater.New(totalRowNum, numberOfOnePage, currentPage, numberOfPageLinks)

	fmt.Printf("generateLinks: %+v\n", generateLinks(p))

	fmt.Printf("page: %+v\n", p)
	fmt.Printf("current: %d\n",p.Current())
	fmt.Printf("hasNext: %d\n",p.HasNext())
	fmt.Printf("hasPrevios: %d\n",p.HasPrevious())
	fmt.Printf("IsFirst: %d\n",p.IsFirst())
	fmt.Printf("IsLast: %d\n",p.IsLast())
	fmt.Printf("Next: %d\n",p.Next())
	fmt.Printf("Pages: %+v\n",p.Pages())
	fmt.Printf("Previous: %d\n",p.Previous())
	fmt.Printf("Total: %d\n",p.Total())
	fmt.Printf("TotalPages: %d\n",p.TotalPages())



}