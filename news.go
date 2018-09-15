package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

type SitemapIndex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

type Article struct {
  Title string
  Keywords string
  Location string
}

func main() {
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var SiteMap SitemapIndex
  var News News
  xml.Unmarshal(bytes, &SiteMap)

  for _, location := range SiteMap.Locations {
    resp, _ := http.Get(location)
    bytes, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    xml.Unmarshal(bytes, &News)
  }

  var Art Article
  for i,_ := range News.Titles {
    Art = Article{strings.TrimSpace(News.Titles[i]), strings.TrimSpace(News.Keywords[i]), News.Locations[i]}
    fmt.Printf("Title: %s \n Keywords: %s \n URL: %s \n\n\n",
      Art.Title, Art.Keywords, Art.Location)
  }

}
