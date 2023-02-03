package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly"
)

func main() {
	file, err := os.Open("url.txt")
	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		parseRun(fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}

func parseRun(url string) {
	c := colly.NewCollector()

	c.OnHTML("body", func(h *colly.HTMLElement) {
		data := h.DOM.Find("h1").Text()

		bodyArticle, err := h.DOM.Find("div.article-formatted-body > div").Html()
		if err != nil {
			log.Println(err)
		}
		SaveArticles(data, bodyArticle)
	})

	c.Visit(url)
}

func SaveArticles(title, body string) {
	file, err := os.Create("output/" + title + ".md")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	tag := "\ntags: #Habr #Articles \n\n---\n\n"
	file.WriteString(tag)

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(body)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(markdown)

	fmt.Println("Done.")

}
