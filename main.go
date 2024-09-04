package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/iunary/fakeuseragent"

	"github.com/spf13/cobra"

	"github.com/MrChiz/ppmap/pkg/cmd"

	"github.com/MrChiz/ppmap/pkg/gadgets"
)

// some fancy colour variables here
const (
	Info       = "[\033[33mINFO\033[0m]"
	Vulnerable = "[\033[32mVULN\033[0m]"
	Error      = "[\033[31mERRO\033[0m]"
	Exploit    = "[\033[34mEXPL\033[0m]"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ppmap",
		Short: "A tool 4 finding pp(prototype pollution) bug.",
	}
	rootCmd.AddCommand(cmd.ScanCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	//run tool
	var wg sync.WaitGroup
	wg.Add(2)
	go queryEnum(cmd.Url, "?", &wg)
	go queryEnum(cmd.Url, "#", &wg)
	wg.Wait()
}

func queryEnum(u, quote string, wg *sync.WaitGroup) {
	defer wg.Done()
	payloads := [4]string{
		"constructor[prototype][I1Younes]%dTest4me",
		"__proto__.I1Younes%3dTest4me",
		"constructor.prototype.I1Younes%sdTest4me",
		"__proto__[I1Younes]%3dTest4me",
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		//uncomment the following lines to setup a proxy
		//chromedp.ProxyServer("localhost:8080"),
		//chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent(fakeuseragent.RandomUserAgent()),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(
		ctx,
		//uncomment the next line to see the CDP messages
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	for _, payload := range payloads {
		full_url := u + quote + payload
		// run task list
		var res string
		err := chromedp.Run(ctx,
			chromedp.Navigate(full_url),
			chromedp.Evaluate(`window.I1Younes`, &res),
		)
		if err != nil {
			log.Printf(Error+" %s", full_url)
			continue
		}
		if res != "" {
			log.Printf(Vulnerable+" %s", full_url)
		}
	}
	//now its fingerprinting time
	log.Printf(Info + " Fingerprinting the gadget...")
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
	gadgets.Gad(Exploit, res, u+quote)
}
