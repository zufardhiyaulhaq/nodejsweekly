package scrappers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

type NodeWeekly struct{}

func (s *NodeWeekly) GetWeekly() []communityv1alpha1.ArticleSpec {
	var articles []communityv1alpha1.ArticleSpec

	// scrapper logic here
	// your scrapper must populate articles

	response, err := http.Get("https://nodeweekly.com/issues/latest?layout=bare")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("div#content").Each(func(i int, div *goquery.Selection) {
		div.Find("span.mainlink").Each(func(i int, span *goquery.Selection) {
			var article communityv1alpha1.ArticleSpec
			article.Title = span.Find("a").Text()
			article.Type = "News"
			article.Url, _ = span.Find("a").Attr("href")

			articles = append(articles, article)
		})
	})

	return articles
}

func (s *NodeWeekly) GetWeeklyName() string {
	var weeklyName string

	// scrapper logic here
	// your scrapper must populate weeklyName

	response, err := http.Get("https://nodeweekly.com/issues/latest?layout=bare")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	table := document.Find("div#content").ChildrenFiltered("table:first-of-type").First()
	td := table.Find("td:first-of-type").First()
	text := td.Find("p").Text()
	regex, _ := regexp.Compile(`\d+`)

	weeklyName = "Node Weekly " + regex.FindString(text)
	return weeklyName
}
