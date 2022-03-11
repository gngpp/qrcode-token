package qrcode

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
	"github.com/zf1976/vlog"
	"image"
	"image/color"
	"image/png"
	"os"
)

//goland:noinspection ALL
const (
	QR_CODE_SIZE        = 256
	SHRINK_QR_CODE_SIZE = 35
	MARGIN              = 29
	MULTIPLE            = 6
)

type QrCode struct {
	content         string
	img             image.Image
	Path            string
	kind            int
	genImg          bool
	points          [QR_CODE_SIZE][QR_CODE_SIZE]int
	tmpShrinkPoints [QR_CODE_SIZE][SHRINK_QR_CODE_SIZE]int
	shrinkPoints    [SHRINK_QR_CODE_SIZE][SHRINK_QR_CODE_SIZE]int
}

//NewQrCode 返回二维码结构
func NewQrCode(content string, genImg bool) *QrCode {
	qr := &QrCode{
		content: content,
		genImg:  genImg,
		kind:    0,
	}
	err := qr.generateQrCode()
	if err != nil {
		vlog.Errorf("generate qrcode error: %v", err)
		return nil
	}
	return qr
}

// generateQrCode 生成二维码矩阵
//goland:noinspection GoUnhandledErrorResult
func (qr *QrCode) generateQrCode() error {
	var err error
	var img image.Image
	if qr.kind == 0 {
		code, err := qrcode.Encode(qr.content, qrcode.Medium, QR_CODE_SIZE)
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(code)
		img, err = png.Decode(buf)
	} else if qr.kind == 1 {
		file, err := os.Open(qr.Path)
		if err != nil {
			return err
		}
		defer file.Close()
		dstImag, _, err := image.Decode(file)
		img = imaging.Resize(dstImag, QR_CODE_SIZE, QR_CODE_SIZE, imaging.Lanczos)
	}
	if err != nil {
		return err
	}
	qr.img = img

	if qr.genImg {
		newPng, _ := os.Create("qrcode.png")
		defer newPng.Close()
		png.Encode(newPng, img)
	}

	return nil
}

// 二维码图片二值化 0－1
//goland:noinspection SpellCheckingInspection
func (qr *QrCode) binarization() {
	gray := image.NewGray(image.Rect(0, 0, QR_CODE_SIZE, QR_CODE_SIZE))
	for x := 0; x < QR_CODE_SIZE; x++ {
		for y := 0; y < QR_CODE_SIZE; y++ {
			r32, g32, b32, _ := qr.img.At(x, y).RGBA()
			r, g, b := int(r32>>8), int(g32>>8), int(b32>>8)
			if (r+g+b)/3 > 180 {
				qr.points[y][x] = 0
				gray.Set(x, y, color.Gray{Y: uint8(255)})
			} else {
				qr.points[y][x] = 1
				gray.Set(x, y, color.Gray{Y: uint8(0)})
			}
		}
	}
}

//shrink 缩小二值化数组
func (qr *QrCode) shrink() {
	for x := 0; x < QR_CODE_SIZE; x++ {
		cal := 1
		for y := MARGIN + 1; y < QR_CODE_SIZE-MARGIN; y += MULTIPLE {
			qr.tmpShrinkPoints[x][cal] = qr.points[x][y]
			cal++
		}
	}

	for y := 1; y < SHRINK_QR_CODE_SIZE-1; y++ {
		row := 1
		for x := MARGIN + 1; x < QR_CODE_SIZE-MARGIN; x += MULTIPLE {
			qr.shrinkPoints[row][y] = qr.tmpShrinkPoints[x][y]
			row++
		}
	}
}

//Output 控制台输出二维码
func (qr *QrCode) Output() {
	qr.binarization()
	qr.shrink()
	for x := 0; x < SHRINK_QR_CODE_SIZE; x++ {
		for y := 0; y < SHRINK_QR_CODE_SIZE; y++ {
			if qr.shrinkPoints[x][y] == 1 {
				fmt.Print("\033[40;40m  \033[0m")
				//randColor()
			} else {
				fmt.Print("\033[47;30m  \033[0m")
			}
		}
		fmt.Print("\n")
	}
}
