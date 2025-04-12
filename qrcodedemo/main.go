package main

import (
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	qr, err := qrcode.New("https://blog.csdn.net/qq_42828912?spm=3001.5343&type=blog", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.Black
		qr.WriteFile(256, "./golang_qrcode.png")
	}
}
