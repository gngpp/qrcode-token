package main

import (
	"encoding/base64"
	"fmt"
	"github.com/zf1976/vlog"
	"github.com/zf1976/vlog/timewriter"
	"io"
	"os"
	"qrcode-token/qrcode"
	"qrcode-token/qrcode/model"
	"qrcode-token/util/vjson"
)

func main() {
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	// global settings
	vlog.SetSyncOutput(true)
	vlog.SetOutput(io.MultiWriter(os.Stdout, timeWriter))
	api := &qrcode.Api{}
	content, err := api.GetGeneratorQrCodeContent()
	if err != nil {
		return
	}
	q := qrcode.NewQrCode(content.CodeContent, false)
	q.Print()
	fmt.Println("Please use the mobile client to scan the code to log in.")
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
