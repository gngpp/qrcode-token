package qrcode

import (
	"fmt"
	"github.com/zf1976/vlog"
	"testing"
)

func Test(t *testing.T) {
	vlog.SetLevel(vlog.Level.DEBUG)
	qr := Api{}
	content, err := qr.GetQrCodeContent()
	if err != nil {
		return
	}
	println(content)
	q := NewQrCode(content.CodeContent, true)
	q.Output()
	fmt.Println()
}
