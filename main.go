package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/gngpp/vlog"
	"github.com/gngpp/vlog/timewriter"
	"io"
	"os"
	"qrcode-token/qrcode"
	"qrcode-token/qrcode/model"
	"qrcode-token/util/vjson"
)

func main() {
	var isMobile bool

	flag.BoolVar(&isMobile, "mobile", false, "get the mobile phone token by default")
	flag.Parse()
	fmt.Println(isMobile)
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	// global settings
	vlog.SetSyncOutput(true)
	vlog.SetOutput(io.MultiWriter(os.Stdout, timeWriter))

	api := qrcode.NewApi(isMobile)
	content, err := api.GetGeneratorQrCodeContent()
	if err != nil {
		return
	}
	// new QrCode
	q := qrcode.NewQrCode(content.CodeContent, false)
	// print QrCode
	q.Print()
	fmt.Println("Please use the mobile client to scan the code to log in.")
	// get login result
	qrCodeResult, b := api.GetQueryQrCodeResult()
	if b {
		bytes, err := base64.StdEncoding.DecodeString(qrCodeResult.Content.Data.BizExt)
		if err != nil {
			return
		}
		result := &model.LoginResult{}
		err = vjson.ByteArrayConvert(bytes, result)
		if err != nil {
			return
		}
		vlog.Infof("refresh_token: %v", result.PdsLoginResult.RefreshToken)
	}
}
