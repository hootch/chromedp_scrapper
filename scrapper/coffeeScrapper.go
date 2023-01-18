package scrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

type coffees struct {
	CoffeeName     string `json:"CoffeeName"`
	Price          string `json:"Price"`
	Weight         string `json:"Weight"`
	Summary        string
	ImgURL         string   `json:"ImgURL"`
	Roastery       string   `json:"Roastery"`
	Flavor         []string `json:"Flavor"`
	ExtractionType []string `json:"ExtractionType"`
	CookingType    []string `json:"CookingType"`
	Style          []string `json:"style"`
}

func StartCoffeeScrape() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	// navigate to a page
	var example string

	var style string
	var roastery string
	var flavor string
	var coffeeName string
	var priceNweight string
	var imgURL map[string]string
	var extractionType string
	var cookingType string
	var summary string
	var coffees_ []coffees

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.koke.kr/`),
		chromedp.Click(`ul.style_items__Q896m li+li`, chromedp.NodeVisible),
	)
	errCheck(err)

	var res *runtime.RemoteObject
	for i := 0; i < 10; i++ {
		err := chromedp.Run(ctx,
			chromedp.Evaluate("window.scrollTo(0,document.body.scrollHeight);", &res),
			chromedp.Sleep(1*time.Second),
		)
		errCheck(err)
	}

	var nodes []*cdp.Node
	err = chromedp.Run(ctx,
		chromedp.Nodes(".items_coffee__4sr50 a", &nodes),
	)
	errCheck(err)

	for _, n := range nodes {
		u := n.AttributeValue("href")
		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://www.koke.kr`+u),
			chromedp.Text(".style_subscription__XvQjk", &example),
			chromedp.Text(".style_subscription__XvQjk div+div > h3", &coffeeName),
			chromedp.Text(".style_subscription__XvQjk div+div > span", &priceNweight),
			chromedp.Text(".style_subscription__XvQjk > div > span", &flavor),
			chromedp.Text(".style_wrapper___uC2Q > ul", &extractionType),
			chromedp.Text(".style_wrapper__xSR1v > ul", &cookingType),
			chromedp.Attributes(".style_center__J3hY5 > div > img ", &imgURL),
			chromedp.Text(".style_note__dPg4i  ", &summary),
			chromedp.Text(".style_items__PrjWC", &style),
			chromedp.Text(".style_title__aZIt9", &roastery),
		)
		if err != nil {
			log.Fatal(err)
		}

		var tempSlice []string
		var tempSlice2 []string

		tempSlice = strings.Split(priceNweight, "\n")
		tempSlice2 = strings.Split(flavor, "&")

		var tempURL string
		for key, val := range imgURL {
			if key == "src" {
				tempURL = val
			}
		}

		var tempExtrType []string
		tempExtrType = strings.Split(strings.Join(strings.Fields(strings.TrimSpace(extractionType)), " "), "\n")

		var tempCookingType []string
		tempCookingType = strings.Split(strings.Join(strings.Fields(strings.TrimSpace(cookingType)), " "), "\n")

		var tempStyle []string
		tempStyle = strings.Split(strings.Join(strings.Fields(strings.TrimSpace(style)), " "), "\n")

		c := coffees{
			CoffeeName:     coffeeName,
			Price:          strings.Join(strings.Fields(tempSlice[0]), " "),
			Weight:         strings.Join(strings.Fields(tempSlice[1]), " "),
			Flavor:         tempSlice2,
			Roastery:       roastery,
			Style:          tempStyle,
			ImgURL:         tempURL,
			ExtractionType: tempExtrType,
			CookingType:    tempCookingType,
			Summary:        summary,
		}
		coffees_ = append(coffees_, c)
	}
	e, err := json.Marshal(coffees_)
	f, err := os.Create("data.json")

	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(string(e))
	if err2 != nil {
		log.Fatal(err2)
	}
	err = f.Close()
	errCheck(err)
	fmt.Println("야호")

}
