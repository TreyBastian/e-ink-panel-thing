package main

import (
	// "image"
	"log"

	// "golang.org/x/image/font"
	// "golang.org/x/image/font/basicfont"
	// "golang.org/x/image/math/fixed"
	// "periph.io/x/conn/v3/spi/spireg"
	"github.com/TreyBastian/e-ink-panel-thing/drivers"
	"periph.io/x/devices/v3/epd"

	// "periph.io/x/devices/v3/ssd1306/image1bit"
	"periph.io/x/host/v3"
)

func main() {
	log.Println("We are starting")
	state, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	device := drivers.Driver.Init()
	defer drivers.Driver.Close()

	// b, err := spireg.Open("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer b.Close()
	//
	// log.Println("Creating SPI hat")
	// dev, err := epd.NewSPIHat(b, &EPD4in2V2)
	// if err != nil {
	// 	log.Fatalf("failed to initalize epd %v", err)
	// }
	//
	// // img := image1bit.NewVerticalLSB(dev.Bounds())
	//
	// dev.ClearFrameMemory(0xff)
	// dev.DisplayFrame()

	// f := basicfont.Face7x13
	// drawer := font.Drawer{
	// 	Dst:  img,
	// 	Src:  &image.Uniform{image1bit.On},
	// 	Face: f,
	// 	Dot:  fixed.P(0, img.Bounds().Dy()-1-f.Descent),
	// }
	//
	// drawer.DrawString("Hello World")
	//
	// log.Println("drawing")
	// if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
	// 	log.Fatal(err)
	// }
	// dev.DisplayFrame()
}
