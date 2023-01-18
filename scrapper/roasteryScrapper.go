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

	//"strings"
	"time"
)

type roastery_ struct {
	Name      string
	IconURL   string
	ImgURLS   []map[string]string
	Summary   string
	Address   string
	Awards    string
	Instagram string
	WebSite   string
}

func StartRoasteryScrape() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var iconURL map[string]string
	var awards string
	var nodes []*cdp.Node

	var roasterys []roastery_
	var res *runtime.RemoteObject
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.koke.kr/roasters/`),
	)
	errCheck(err)

	for i := 0; i < 10; i++ {
		err := chromedp.Run(ctx,
			chromedp.Evaluate("window.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' });", &res),
			chromedp.Sleep(1*time.Second),
		)
		errCheck(err)
	}

	err = chromedp.Run(ctx,
		chromedp.Nodes(".items_roaster__CDUf7 a", &nodes),
	)
	errCheck(err)

	for _, n := range nodes {
		u := n.AttributeValue("href")

		var name string
		var summary string
		var imgURLS []map[string]string

		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://www.koke.kr`+u),
			chromedp.AttributesAll("div.style_wrapper__acQsm > div img", &imgURLS), //파싱 필요없음
			chromedp.Attributes(".style_wrapper__acQsm > img", &iconURL),
			chromedp.Text(".style_title__aZIt9", &name),
			chromedp.Text("p.style_introduce__iKEQp", &summary),
		)
		var tempIconURL string
		for key, val := range iconURL {
			if key == "src" {
				tempIconURL = val
			}
		}
		searchCtx, searchCancel := context.WithTimeout(ctx, 1*time.Second)
		defer searchCancel()

		err1 := chromedp.Run(searchCtx,
			chromedp.Text("ul.style_awards__M0Wi3", &awards),
			chromedp.Click(".style_more__NDJlO", chromedp.NodeVisible),
			chromedp.Text(".style_awards__M0Wi3", &awards),
		)
		if err1 != nil {
			//fmt.Println("Didn't find Ack")
			//return
		}
		errCheck(err)

		searchCtx2, searchCancel := context.WithTimeout(ctx, 1*time.Second)
		defer searchCancel()

		var addr map[string]string
		err2 := chromedp.Run(searchCtx2,
			chromedp.Attributes("ul.style_links__0JMv1 > li > a", &addr),
		)
		if err2 != nil {
		}
		var addHref string
		for key, val := range addr {
			if key == "href" {
				addHref = val
			}
		}
		searchCtx3, searchCancel := context.WithTimeout(ctx, 1*time.Second)
		defer searchCancel()

		var insta map[string]string
		err3 := chromedp.Run(searchCtx3,
			chromedp.Attributes("ul.style_links__0JMv1 > li+li > a", &insta),
		)
		if err3 != nil {
			//fmt.Println("Didn't find Ack")
			//return
		}
		var instaHref string
		for key, val := range insta {
			if key == "href" {
				instaHref = val
			}
		}

		searchCtx4, searchCancel := context.WithTimeout(ctx, 1*time.Second)
		defer searchCancel()

		var web map[string]string
		err4 := chromedp.Run(searchCtx4,
			chromedp.Attributes("ul.style_links__0JMv1 > li+li+li > a", &web),
		)
		if err4 != nil {
			//fmt.Println("Didn't find Ack")
			//return
		}
		var webHref string
		for key, val := range web {
			if key == "href" {
				webHref = val
			}
		}
		fmt.Println(addHref)
		fmt.Println(instaHref)
		fmt.Println(webHref)
		fmt.Println(name)
		fmt.Println(summary)

		c := roastery_{
			Name:      name,
			IconURL:   tempIconURL,
			ImgURLS:   imgURLS,
			Awards:    awards,
			Summary:   summary,
			Address:   addHref,
			Instagram: instaHref,
			WebSite:   webHref,
		}
		roasterys = append(roasterys, c)
	}
	e, err := json.Marshal(roasterys)
	f, err := os.Create("roasterys.json")

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

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
