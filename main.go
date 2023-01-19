package main

import (
	"chromedp_scrapper/scrapper"
	"fmt"
	"github.com/labstack/echo"
)

func handleHome(ctx echo.Context) error {
	return ctx.File("home.html")
}

func handleCoffeeScrape(ctx echo.Context) error {
	fmt.Println("야호")
	scrapper.StartCoffeeScrape()
	return ctx.Attachment("coffee.json", "coffee.json")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape/coffee", handleCoffeeScrape)
	//e.POST("/scrape/roastery", handleRoasteryScrape)
	e.Logger.Fatal(e.Start(":4000"))
}
