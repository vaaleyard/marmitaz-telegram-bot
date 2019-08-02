package site

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

// TemperoDeMaeIsOpen returns a bool if the restaurant is open or not
func TemperoDeMaeIsOpen() bool {

	exist := false

	c := colly.NewCollector(
		colly.URLFilters(
			// Visit only urls which belongs to the site
			regexp.MustCompile("https://marmitaz\\.pushsistemas\\.com\\.br/.*"),
		),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		log.Printf("Link found: %s\n", e.Attr("href"))

		if e.Text == "Pedidos Açaí" || e.Attr("href") == "pedidos_acai.php" {
			exist = false
		}
	})

	c.Visit("https://marmitaz.pushsistemas.com.br/")

	return exist
}