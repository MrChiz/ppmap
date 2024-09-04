package runner

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/iunary/fakeuseragent"

	"github.com/MrChiz/ppmap/pkg/cmd"
	"github.com/MrChiz/ppmap/pkg/gadgets"
)

func queryEnum(u, quote string, wt bool, wg *sync.WaitGroup) {
	defer wg.Done()
	var out []string
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(fakeuseragent.RandomUserAgent()),
	)
	//Set proxy setver
	if cmd.Proxy != "" {
		opts = append(opts, chromedp.ProxyServer(cmd.Proxy))
		opts = append(opts, chromedp.Flag("ignore-certificate-errors", true))
	}
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(
		ctx,
		//uncomment the next line to see the CDP messages
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	for _, payload := range Payload() {
		full_url := u + quote + payload
		// run task list
		var res string
		err := chromedp.Run(ctx,
			chromedp.Navigate(full_url),
			chromedp.Evaluate(`window.I1Younes`, &res),
		)
		if err != nil {
			log.Printf("[%s] %s", red("Error"), full_url)
			continue
		}
		if res != "" {
			fmt.Printf("[%s] %s", green("Vulnerable"), full_url)
			out = append(out, full_url)
		}
	}
	//now its fingerprinting time
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		//change the value 5 to a higher one if your internet connection is slow
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(gadgets.Fingerprint, &res),
	)
	if err != nil {
		log.Fatal(err)
	}
	exp := gadgets.Gad(cyan("Exploit"), res, u+quote)
	if len(exp) >= 1 {
		out = append(out, out...)
	}
	//Save
	if len(out) >= 1 && wt {
		for _, i := range out {
			Save(cmd.Output, i)
		}
	}
}
