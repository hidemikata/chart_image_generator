package main

import (
	"chart_image_generator/chart"
	"chart_image_generator/model"
	"fmt"
)

func main() {
	fmt.Println("main")

	a := model.GetPriceTest()
	fmt.Print(a)
	chart.CreateChart()
}
