package tempMain

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

func main() {
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

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.koke.kr/`),
		//chromedp.Click(`a.style_text__qlqJK`, chromedp.NodeVisible),
		//chromedp.Text(`a.style_text__qlqJK`, &example),
		//chromedp.Nodes("a", &nodes),
		//chromedp.Text([]cdp.NodeID{75}, &example, chromedp.ByNodeID),
		chromedp.Click(`ul.style_items__Q896m li+li`, chromedp.NodeVisible),
		//chromedp.Text(`.items_coffee__4sr50 a`, &example),
		//chromedp.SendKeys(`div.style_listWrapper__BVuv_`, kb.ArrowDown),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.Sleep(1*time.Second),

		//chromedp.ActionFunc(func(context.Context) error {
		//	log.Printf("waiting 3s for box to become visible")
		//	return nil
		//}),
		//
		//chromedp.ActionFunc(func(context.Context) error {
		//	log.Printf(">>>>>>>>>>>>>>>>>>>> BOX1 IS VISIBLE")
		//	return nil
		//}),
		//chromedp.WaitVisible(`#box2`),
		//chromedp.ActionFunc(func(context.Context) error {
		//	log.Printf(">>>>>>>>>>>>>>>>>>>> BOX2 IS VISIBLE")
		//	return nil
		//}),
		chromedp.Nodes(".items_coffee__4sr50 a", &nodes),
		//chromedp.Text(`ul.style_items__Q896m`, &example),
		//chromedp.WaitVisible(`body`),
		//chromedp.Text(`.items_coffees__jhp3v > .items_coffee__4sr50> a > .item_wrapper__BdFuT > .item_characters__AJbwz`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range nodes {
		u := n.AttributeValue("href")
		//fmt.Println(u)
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

		log.Println(summary)
		//log.Println(priceNweight)
		//log.Println(flavor)
		//log.Println(style)
		//log.Println(roastery)

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
		log.Println(c.CoffeeName)
		//log.Println(c.Price)
		//log.Println(c.Weight)
		//log.Println(c.Flavor)
		//log.Println(c.Roastery)
		//log.Println(c.Style)
		//log.Println(c.ImgURL)
		//log.Println(c.ExtractionType)
		//log.Println(c.CookingType)
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
	f.Close()
	fmt.Println("야호")

}

//
//type coffees struct {
//	flavor     []string
//	price      string
//	coffeeName string
//	weight     string
//}
