package chart

import (
	"chart_image_generator/model"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const chart_num = 100
const chart_x_pix = 9

func CreateChart() {
	fmt.Println("createChart")

	width := chart_num * chart_x_pix
	height := 1000

	canvas_myBound := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: width, Y: height}}

	canvas := image.NewRGBA(canvas_myBound)
	//初期化
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			canvas.SetRGBA(i, j, color.RGBA{255, 255, 255, 255})
		}
	}

	data := model.GetPriceTest()
	data_min := 9999999.0
	data_max := 0.0
	for _, v := range data {
		if data_min > v.Low {
			data_min = v.Low
		}
		if data_max < v.High {
			data_max = v.High
		}
	}
	price_diff := data_max - data_min
	one_price_on_pix := float64(height) / price_diff

	set_i := 0
	k := -1
	//(値段-data_min) * one_price_on_pix = heightに設定する値
	for j := 0; j <= width; j++ {
		if j%chart_x_pix == 0 {
			k = k + 1
		}
		if k >= len(data) {
			continue
		}
		for i := 0; i <= height; i++ {
			//枠線
			set_i = height - i

			//上下枠線
			if int(((data[k].Close-data_min)*one_price_on_pix)) == i || int(((data[k].Open-data_min)*one_price_on_pix)) == i {
				canvas.SetRGBA(j, set_i, color.RGBA{0, 0, 0, 255})
			}
			//横枠線
			//陰線
			if data[k].Close <= data[k].Open {
				//枠
				if int(((data[k].Close-data_min)*one_price_on_pix)) < i && int(((data[k].Open-data_min)*one_price_on_pix)) > i {
					canvas.SetRGBA(j, set_i, color.RGBA{0, 0, 0, 255})
				}
				//髭
				if j%chart_x_pix == chart_x_pix/2 && int(((data[k].Low-data_min)*one_price_on_pix)) <= i && int(((data[k].High-data_min)*one_price_on_pix)) >= i {
					canvas.SetRGBA(j, set_i, color.RGBA{0, 0, 0, 255})
				}
				//陽線
			} else if data[k].Close > data[k].Open {
				//枠
				if (j%chart_x_pix == 0 || j%chart_x_pix == chart_x_pix-1) && int(((data[k].Open-data_min)*one_price_on_pix)) < i && int(((data[k].Close-data_min)*one_price_on_pix)) > i {
					canvas.SetRGBA(j, set_i, color.RGBA{0, 0, 0, 255})
				}

				//髭
				if j%chart_x_pix == chart_x_pix/2 && int(((data[k].Low-data_min)*one_price_on_pix)) <= i && int(((data[k].High-data_min)*one_price_on_pix)) >= i &&
					!(int(((data[k].Open-data_min)*one_price_on_pix)) <= i && int(((data[k].Close-data_min)*one_price_on_pix)) >= i) {
					canvas.SetRGBA(j, set_i, color.RGBA{0, 0, 0, 255})
				}
			}

		}
	}

	savefile, err := os.Create("/Users/ajikatashuuyoshimi/Desktop/go/chart_image_generator/test.png")
	if err != nil {
		fmt.Println("保存するためのファイルが作成できませんでした。")
		os.Exit(1)
	}
	defer savefile.Close()
	// PNG形式で保存する
	png.Encode(savefile, canvas)
}
