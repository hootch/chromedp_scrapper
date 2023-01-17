package tempMain

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
)

type roastery_ struct {
}

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var imgURLS map[string]string

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.koke.kr/roasters/`),
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
		chromedp.Nodes(".items_roaster__CDUf7 a", &nodes),
	)
	errCheck(err)

	for _, n := range nodes {
		u := n.AttributeValue("href")

		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://www.koke.kr`+u),
			chromedp.Attributes(".style_wallImages__vfQb3", &imgURLS),
		)
		errCheck(err)
		fmt.Println(imgURLS)
		var tempURL string
		for key, val := range imgURLS {
			if key == "src" {
				tempURL = val
			}
		}
		fmt.Println(tempURL)

	}

}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
