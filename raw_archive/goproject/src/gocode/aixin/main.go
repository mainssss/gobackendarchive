package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

// 申明画板的颜色组
var palette = []color.Color{color.White, color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}}

func main() {
	const (
		nframes = 50  // GIF的帧数
		delay   = 10  // 每帧间的时间间隔
		size    = 400 // 图片大小
	)

	a := 0.0
	anim := gif.GIF{LoopCount: nframes} // GIF文件对象

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, size+1, size+1)
		img := image.NewPaletted(rect, palette) // 新建一个画板，指定宽度、高度和调色板只要色彩

		for x := -2.0; x < 2.0; x += 0.0001 {
			f1 := math.Pow(math.Abs(x), 2.0/3)
			f2 := math.E / 4 * math.Sqrt(math.Pi-math.Pow(x, 2.0)) * math.Sin(math.Pi*a*x)
			if math.IsNaN(f2) {
				f2 = 0
			}
			y := -(f1 + f2)
			img.SetColorIndex(int(x*size/4)+200, int(y*size/4)+250, 2)
		}
		a++

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	var filename = "test.gif"
	if len(os.Args) > 1 {
		filename = os.Args[1] + ".gif"
	}
	file, _ := os.Create(filename)
	defer file.Close()

	gif.EncodeAll(file, &anim)
}
