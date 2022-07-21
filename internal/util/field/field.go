package field

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetInputField(body string, old map[string]string) map[string]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var field map[string]string

	if old != nil {
		field = old
	} else {
		field = make(map[string]string)
	}
	// Find the review items
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		if name == "" {
			return
		}
		field[name] = value
	})
	return field
}
