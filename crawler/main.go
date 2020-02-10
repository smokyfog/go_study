package main

import (
	"fmt"
	"regexp"
	"yq.com/go_study/crawler/fetcher"
)

func main() {
	fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//fmt.Printf("%s\n", all)
	printCityList(all)
}



func printCityList(contents []byte)  {
	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {

		fmt.Printf("City: %s, URL: %s\n", m[1], m[0])

	}
	fmt.Printf("Matches found: %d\n", len(matches))
}