package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"njajal/indicator"
	"njajal/model"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

var baseUrl = "https://query1.finance.yahoo.com/v8/finance/chart/{{stock_code}}.JK?period1={{start_date}}&period2={{end_date}}&interval=1d&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US&source=cosaic"

var startDate = time.Now().AddDate(0, 0, -30).Unix()
var endDate = time.Now().Unix()

func main() {
	stocks := readAllEmiten()

	for _, stock := range stocks {
		price := getPrice(stock)
		kLine, dLine := indicator.Stochastic(price.Chart.Result[0].Indicators.Quote[0], 0)
		kLine1, dLine1 := indicator.Stochastic(price.Chart.Result[0].Indicators.Quote[0], 1)
		if kLine > 20 && kLine < 50 && dLine1 > kLine1 && kLine >= dLine {
			fmt.Println(stock, roundToTwo(kLine), roundToTwo(dLine), roundToTwo(kLine1), roundToTwo(dLine1))
		}
	}
}

func getPrice(stockCode string) model.Stock {
	url := strings.ReplaceAll(baseUrl, "{{stock_code}}", stockCode)
	url = strings.ReplaceAll(url, "{{start_date}}", fmt.Sprintf("%d", startDate))
	url = strings.ReplaceAll(url, "{{end_date}}", fmt.Sprintf("%d", endDate))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed formed request %s", stockCode), err.Error())
	}
	req.Header.Add("User-Agent", "PostmanRuntime/7.51.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error response %s", stockCode), err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed parse respon body %s", stockCode), err.Error())
	}

	var price model.Stock
	err = json.Unmarshal(body, &price)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed parse respon body %s", stockCode), err.Error())
	}

	return price
}

func readAllEmiten() []string {
	xlsx, err := excelize.OpenFile("./Daftar Saham  - 20260118.xlsx")
	if err != nil {
		log.Fatal("failed read excel", err.Error())
	}

	sheet1Name := "Sheet1"

	var stocks []string
	for i := 2; i < 957; i++ {
		stock, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i))
		stocks = append(stocks, stock)
	}

	return stocks
}

func roundToTwo(num float64) float64 {
	return math.Round(num*100) / 100
}
