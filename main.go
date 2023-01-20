package main

import (
	"chromedp_scrapper/scrapper"
)

func main() {
	scrapper.StartAsyncScrapeRoastery()
}

//
//package main
//
//import (
//	"chromedp_scrapper/scrapper"
//	"github.com/labstack/echo"
//)
//
//func handleHome(ctx echo.Context) error {
//	return ctx.File("home.html")
//}
//
//func handleCoffeeScrape(ctx echo.Context) error {
//	scrapper.StartCoffeeScrape()
//	return ctx.Attachment("coffee.json", "coffees.json")
//}
//
//func handleRoasteryScrape(ctx echo.Context) error {
//	scrapper.StartAsyncScrapeRoastery()
//	return ctx.Attachment("roasterys.json", "roasterys.json")
//}
//
//func main() {
//	e := echo.New()
//	e.GET("/", handleHome)
//	e.POST("/scrape/coffee", handleCoffeeScrape)
//	e.POST("/scrape/roastery", handleRoasteryScrape)
//	e.Logger.Fatal(e.Start(":4000"))
//}
