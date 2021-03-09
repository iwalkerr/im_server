package main

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	qrcodeReader "github.com/tuotoo/qrcode"
)

func TestMain(t *testing.T) {
	// qrcode.WriteFile("http://www.flysnow.org/", qrcode.High, 256, "./blog_qrcode.png")
	//Boombuler
	dest2 := "qrcode2.png"
	content := "https://www.cnblogs.com/fanbi"
	size := 200
	CreateQRCodeByBoombuler(content, qr.M, size, dest2)

	// fmt.Println(ReadQRCode("qrcode2.png"))

}

func ReadQRCode(filename string) (content string) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcodeReader.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return qrmatrix.Content
}

func CreateQRCodeByBoombuler(content string, quality qr.ErrorCorrectionLevel, size int, dest string) (err error) {
	qrCode, err := qr.Encode(content, quality, qr.Auto)
	if err != nil {
		return
	}

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return
	}

	// create the output file
	file, err := os.Create(dest)
	if err != nil {
		return
	}

	defer file.Close()
	// encode the barcode as png
	err = png.Encode(file, qrCode)
	if err != nil {
		return
	}

	return
}
