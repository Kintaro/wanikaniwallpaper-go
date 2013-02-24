package main

import "flag"

var (
	outputFile	= flag.String("output", "out.png", "path to ouput file")
	orderFilename = flag.String("orderfile", "data/order", "path to order file")
	key = flag.String("key", "", "API key")
)

func main() {
	flag.Parse()

	kanjiList := GetKanjiForApiKey(*key)
	order := NewOrder(*orderFilename)
	order.Update(kanjiList)
		
	renderer := NewRenderer()

	Draw(order, renderer, *width, *height)

	renderer.SaveImage(*outputFile)
}
