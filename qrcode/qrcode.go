package qrcode

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"github.com/zf1976/vlog"
	"image"
	"image/png"
	"os"
)

type QrCode struct {
	content       string
	img           image.Image
	bytes         *[]byte
	generateImage bool
}

func NewQrCode(content string, generateImage bool) *QrCode {
	qr := &QrCode{
		content:       content,
		generateImage: generateImage,
	}

	err := qr.generateQrCode()
	if err != nil {
		vlog.Errorf("generate qrcode error: %v", err)
		return nil
	}
	return qr
}

//generateQrCode 生成二维码
func (qr *QrCode) generateQrCode() error {
	code, err := qr.generate(qr.content)
	if err != nil {
		return err
	}
	qr.bytes = &code

	buf := bytes.NewBuffer(code)
	img, err := png.Decode(buf)
	if err != nil {
		return err
	}

	qr.img = img

	if qr.generateImage {
		newPng, _ := os.Create("qrcode.png")
		//goland:noinspection GoUnhandledErrorResult
		defer newPng.Close()
		err := png.Encode(newPng, img)
		if err != nil {
			vlog.Errorf("encode qrcode png error: %v", err)
			return err
		}
	}

	return nil
}

func (qr *QrCode) generate(content string) ([]byte, error) {
	var b []byte
	b, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		vlog.Errorf("qrcode generate error: %v", err.Error())
		return b, err
	}
	vlog.Debugf("qrcode content: %v", content)
	return b, nil
}

func (qr QrCode) Get() []byte {
	return *qr.bytes
}
