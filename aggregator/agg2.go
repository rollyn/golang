package main 

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml"
)


type SitemapIndex struct {
	Locations []string `xml:"url>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}



func main() {
	
	resp, _ := http.Get("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	var n News

	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)
	// for _, Location := range s.Locations {
	// 	fmt.Printf("\n%s",Location)
	// }

	xml.Unmarshal(bytes, &n)
	for _, Title := range n.Titles {
		fmt.Printf("\n%s",Title)
	}	

}