package main

import (
	"AliyunPanToken/qrcode"
	"AliyunPanToken/qrcode/model"
	"AliyunPanToken/util/vjson"
	"encoding/base64"
	"github.com/zf1976/vlog"
	"github.com/zf1976/vlog/timewriter"
	"io"
	"os"
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
	content, err := api.GetQrCodeContent()
	if err != nil {
		return
	}
	q := qrcode.NewQrCode(content.CodeContent, false)
	q.Output()
	vlog.Info("Please use the mobile client to scan the code to log in.")
	queryQrCodeResult, b := api.QueryQrCode()
	if b {
		bytes, err := base64.StdEncoding.DecodeString(queryQrCodeResult.Content.Data.BizExt)
		if err != nil {
			return
		}
		result := &model.LoginResult{}
		err = vjson.ByteArrayConvert(bytes, result)
		if err != nil {
			return
		}
		vlog.Infof("login result:\n%v", vjson.PrettifyString(result))
		vlog.Infof("refresh_token: %v", result.PdsLoginResult.RefreshToken)
	}
}
