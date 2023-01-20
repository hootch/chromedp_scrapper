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

	//"strings"
	"time"
)

type roastery_ struct {
	Name    string
	IconURL string
	ImgURLS string
	//ImgURLS   []map[string]string
	Summary   string
	Address   string
	Awards    string
	Instagram string
	WebSite   string
}

// 동기식
func StartRoasteryScrape() {
	startTime := time.Now()
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

		err1 := chromedp.Run(searchCtx,
			chromedp.Text("ul.style_awards__M0Wi3", &awards),
			chromedp.Click(".style_more__NDJlO", chromedp.NodeVisible),
			chromedp.Text(".style_awards__M0Wi3", &awards),
		)
		if err1 != nil {
		}
		errCheck(err)

		searchCtx2, searchCancel := context.WithTimeout(ctx, 1*time.Second)

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

		var insta map[string]string
		err3 := chromedp.Run(searchCtx3,
			chromedp.Attributes("ul.style_links__0JMv1 > li+li > a", &insta),
		)
		if err3 != nil {
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
		}
		var webHref string
		for key, val := range web {
			if key == "href" {
				webHref = val
			}
		}

		c := roastery_{
			Name:      name,
			IconURL:   tempIconURL,
			ImgURLS:   fmt.Sprint(imgURLS),
			Awards:    awards,
			Summary:   summary,
			Address:   addHref,
			Instagram: instaHref,
			WebSite:   webHref,
		}
		roasterys = append(roasterys, c)
	}

	e, err := json.Marshal(roasterys)
	errCheck(err)

	f, err := os.Create("roasterys.json")
	errCheck(err)

	_, err2 := f.WriteString(string(e))
	errCheck(err2)

	err = f.Close()
	errCheck(err)
	endTime := time.Now()
	fmt.Println("Operating time: ", endTime.Sub(startTime))
}

// 비동기식
func StartAsyncScrapeRoastery() {
	startTime := time.Now()
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	var nodes []*cdp.Node
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
	cancel()

	ch := make(chan roastery_)
	for _, n := range nodes {
		go startMainScrapeForRoastery(n.AttributeValue("href"), ch)
	}
	var roasterys []roastery_
	for i := 0; i < len(nodes); i++ {
		temp := <-ch
		//fmt.Println(temp)
		roasterys = append(roasterys, temp)

	}
	e, err := json.Marshal(roasterys)
	errCheck(err)

	f, err := os.Create("roasterys.json")
	errCheck(err)

	_, err2 := f.WriteString(string(e))
	errCheck(err2)

	err = f.Close()
	errCheck(err)
	endTime := time.Now()
	fmt.Println("Operating time: ", endTime.Sub(startTime))

}

func startMainScrapeForRoastery(url string, ch chan<- roastery_) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	var name string
	var summary string
	var imgURLS []map[string]string
	var iconURL map[string]string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.koke.kr`+url),
		chromedp.AttributesAll("div.style_wrapper__acQsm > div img", &imgURLS), //파싱 필요없음
		chromedp.Attributes(".style_wrapper__acQsm > img", &iconURL),
		chromedp.Text(".style_title__aZIt9", &name),
		chromedp.Text("p.style_introduce__iKEQp", &summary),
	)
	errCheck(err)
	var tempIconURL string
	for key, val := range iconURL {
		if key == "src" {
			tempIconURL = val
		}
	}

	ch2 := make(chan string)
	ch3 := make(chan map[string]string)
	go searchRoastery2(ctx, ch2)
	go searchRoastery3(ctx, ch3)

	awards := <-ch2
	addrNInstaNWeb := <-ch3
	Address := addrNInstaNWeb["위치"]
	Instagram := addrNInstaNWeb["Instagram"]
	WebSite := addrNInstaNWeb["WebSite"]

	ch <- roastery_{
		Name:      name,
		IconURL:   tempIconURL,
		ImgURLS:   fmt.Sprint(imgURLS),
		Summary:   summary,
		Awards:    awards,
		Address:   Address,
		Instagram: Instagram,
		WebSite:   WebSite,
	}

	//ch <- roastery_{
	//	Name:      name,
	//	IconURL:   tempIconURL,
	//	ImgURLS:   fmt.Sprint(imgURLS),
	//	Awards:    <-ch2,
	//	Summary:   summary,
	//	Address:   <-ch3,
	//	Instagram: <-ch4,
	//	WebSite:   <-ch5,
	//}
}

func searchRoastery2(ctx context.Context, ch2 chan<- string) {
	searchCtx, searchCancel := context.WithTimeout(ctx, 5*time.Second)
	var awards string
	defer searchCancel()
	err := chromedp.Run(searchCtx,
		chromedp.ActionFunc(func(ctx3 context.Context) error {
			chromedp.Text("ul.style_awards__M0Wi3", &awards).Do(ctx3)
			return nil
		}),
		chromedp.ActionFunc(func(ctx3 context.Context) error {
			chromedp.Click(".style_more__NDJlO", chromedp.NodeVisible).Do(ctx3)
			chromedp.Text(".style_awards__M0Wi3", &awards).Do(ctx3)
			return nil
		}),
	)
	errCheck(err)
	ch2 <- awards
}

func searchRoastery3(ctx context.Context, ch3 chan<- map[string]string) {
	searchCtx2, searchCancel := context.WithTimeout(ctx, 5*time.Second)
	defer searchCancel()
	var addr string
	err := chromedp.Run(searchCtx2,
		chromedp.ActionFunc(func(ctx3 context.Context) error {
			err := chromedp.Text("ul.style_links__0JMv1", &addr).Do(ctx3)
			errCheck(err)
			return nil
		}),
	)
	errCheck(err)
	elements := strings.Split(addr, "\n")
	elementMap := map[string]string{
		"위치":        "",
		"Instagram": "",
		"Website":   "",
	}
	for i := 0; i < len(elements); i += 2 {
		elementMap[elements[i]] = elements[i+1]
	}
	ch3 <- elementMap
}

//func scrapeRoastery1(ctx context.Context, url string, ch1 chan roastery_) {
//	var name string
//	var summary string
//	var imgURLS []map[string]string
//	var iconURL map[string]string
//
//	err := chromedp.Run(ctx,
//		chromedp.Navigate(`https://www.koke.kr`+url),
//		chromedp.AttributesAll("div.style_wrapper__acQsm > div img", &imgURLS), //파싱 필요없음
//		chromedp.Attributes(".style_wrapper__acQsm > img", &iconURL),
//		chromedp.Text(".style_title__aZIt9", &name),
//		chromedp.Text("p.style_introduce__iKEQp", &summary),
//	)
//	errCheck(err)
//	var tempIconURL string
//	for key, val := range iconURL {
//		if key == "src" {
//			tempIconURL = val
//		}
//	}
//	ch2 := make(chan string)
//	go searchRoastery2(ctx, ch2)
//	ch3 := make(chan string)
//	go searchRoastery3(ctx, ch3)
//	ch4 := make(chan string)
//	go searchRoastery4(ctx, ch4)
//	ch5 := make(chan string)
//	go searchRoastery5(ctx, ch5)
//
//	ch <- roastery_{
//		Name:      name,
//		IconURL:   tempIconURL,
//		ImgURLS:   fmt.Sprint(imgURLS),
//		Awards:    <-ch2,
//		Summary:   summary,
//		Address:   <-ch3,
//		Instagram: <-ch4,
//		WebSite:   <-ch5,
//	}
//
//}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
