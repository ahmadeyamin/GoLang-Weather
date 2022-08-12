package main

import (
	"regexp"

	"github.com/electricbubble/go-toast"
	"github.com/gocolly/colly"
)

var temperature string
var nextHourRainPrediction string

func main() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		// fmt.Println("Visited", r.Request.URL.String())
		temperature = e.ChildText("span.CurrentConditions--tempValue--3a50n")
		// nextHourRainPrediction = strconv.Itoa()

		exp := e.DOM.Find("span.Column--precip--2ck8J").Eq(4).Text()
		expResult := regexp.MustCompile("\\d+%").FindAllString(exp, -1)[0]
		expText := regexp.MustCompile("\\d+%").ReplaceAllString(exp, "")

		nextHourRainPrediction = expResult
		// number to string

		_ = toast.Push(
			"Now temperature is "+temperature+" degrees, with a "+nextHourRainPrediction+" "+expText,
			toast.WithTitle("Weather Alert"),
			toast.WithAudio(toast.Mail),
			toast.WithDuration(toast.NotificationDuration(toast.Long)),
		)
	})

	c.Visit("https://weather.com/weather/today/l/22.685402078923453,90.63322120653807?par=google&temp=c&unit=c&cc=us&partner=none&output=json")
}

// https://weather.com/weather/today/l/22.685402078923453,90.63322120653807?par=google&temp=c&unit=c&cc=us&partner=none&output=json
