package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("parameter not found.\n")
		return
	}

	// TODO: 引数がURLになっているかのチェック

	// driver準備
	agoutiDriver := agouti.ChromeDriver(agouti.Browser("chrome"))
	defer agoutiDriver.Stop()
	if err := agoutiDriver.Start(); err != nil {
		fmt.Printf("Failed to Start. %s\n", err)
		return
	}
	page, err := agoutiDriver.NewPage()
	if err != nil {
		fmt.Printf("Failed NewPage. %s\n", err)
		return
	}

	day := time.Now()
	const layout2 = "20060102_150405"
	fileSuffix := day.Format(layout2)

	// 引数のURLを開く
	page.Navigate(os.Args[1])
	// スクリーンショット
	page.Screenshot("Screenshot_" + fileSuffix + "_.png")
}
