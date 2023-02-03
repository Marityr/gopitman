package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
	file, err := os.Create(title + ".md")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	tag := "\ntags: #Go \n\n---\n\n"
	file.WriteString(tag)

	replacer := strings.NewReplacer("<pre><code class=\"go\">", "\n```go\n", "</code></pre>", "\n```\n")
	out := replacer.Replace(body)

	replacer = strings.NewReplacer("<br/>", "", "<ul>", "", "</ul>", "\n", "<li>", "- ", "</li>", "\n")
	out = replacer.Replace(out)

	mdHeader := "\n#"
	for i := 1; i <= 6; i++ {
		header := "<h" + strconv.Itoa(i) + ">"
		replacer = strings.NewReplacer(header, mdHeader+" ")
		out = replacer.Replace(out)
		header = "</h" + strconv.Itoa(i) + ">"
		replacer = strings.NewReplacer(header, "\n\n")
		out = replacer.Replace(out)
		mdHeader += "#"
	}

	file.WriteString(out)

	fmt.Println("Done.")

}
