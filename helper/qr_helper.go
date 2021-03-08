package helper

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	qrcodeReader "github.com/tuotoo/qrcode"
)

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

// 读取二维码内容
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
