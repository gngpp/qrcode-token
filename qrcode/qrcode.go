package qrcode

import (
	"bytes"
	"github.com/mdp/qrterminal/v3"
	"github.com/skip2/go-qrcode"
	"github.com/zf1976/vlog"
	"image"
	"image/png"
	"io"
	"os"
)

const Size = 256

type QrCode struct {
	content string
	img     image.Image
	genImg  bool
}

func NewQrCode(content string, genImg bool) *QrCode {
	qr := &QrCode{
		content: content,
		genImg:  genImg,
	}
	err := qr.generateQrCode()
	if err != nil {
		vlog.Errorf("generate qrcode error: %v", err)
		return nil
	}
	return qr
}

//goland:noinspection GoUnhandledErrorResult
func (qr *QrCode) generateQrCode() error {
	var err error
	var img image.Image
	code, err := qrcode.Encode(qr.content, qrcode.Medium, Size)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(code)
	img, err = png.Decode(buf)
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

func (qr *QrCode) Get() ([]byte, error) {
	encode, err := qrcode.Encode(qr.content, qrcode.Medium, Size)
	if err != nil {
		return nil, err
	}
	return encode, nil
}

func (qr *QrCode) Output(out ...io.Writer) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}

	if len(out) == 0 {
		qrterminal.GenerateWithConfig(qr.content, config)
	} else if len(out) > 0 {
		multiWriter := io.MultiWriter(out...)
		config.Writer = multiWriter
		qrterminal.GenerateWithConfig(qr.content, config)
	}
}
