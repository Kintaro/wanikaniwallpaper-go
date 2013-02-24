package main

import "fmt"

func WastedSpace(width int, height int, num int, ratio float64) (int, float64) {
	if num == 0 {
		return height, 1.0
	}

	if num < width {
		width = num
	}

	if num % width == 0 {
		height = num / width
	} else {
		height = (num / width) + 1
	}

	contentRatio := float64(width) / float64(height)

	var wasted float64
	var area float64
	if contentRatio > ratio {
		wasted = (1 / ratio - 1 / contentRatio) * ratio
		cwidth := 1.0 / float64(width)
		area = cwidth * cwidth * ratio
	} else {
		wasted = (ratio - contentRatio) / ratio
		cheight := 1.0 / float64(height)
		area = (cheight * cheight) / ratio
	}

	missing := float64(width * height - num)
	wasted = wasted + missing * area

	return height, wasted
}

func FindBest(num int, ratio float64) (int, int, float64) {
	bestWidth := -1
	bestHeight := -1
	bestWasted := -1.0

	for i := 1; i <= num; i++ {
		currentWidth := i
		currentHeight := 0
		currentHeight, currentWasted := WastedSpace(currentWidth, currentHeight, num, ratio)

		if bestWasted < 0 || currentWasted < bestWasted {
			bestWidth = currentWidth
			bestHeight = currentHeight
			bestWasted = currentWasted
		}
	}

	return bestWidth, bestHeight, bestWasted
}

func Draw(order *Order, renderer *Renderer, width int, height int) {
	ratio := float64(width) / float64(height)
	w, h, _ := FindBest(order.Size(), ratio)

	fontsize := height / h
	renderer.SetFontSize(fontsize)

	fmt.Println("Fontsize is ", fontsize)

	for i := 0; i < order.Size(); i++ {
		x := (i % w) * fontsize
		y := (i / w) * fontsize

		renderer.DrawKanji(order.KanjiForPosition(i), x, y + h)
	}
}
