package qrcode

import (
	"testing"
)

func Test(t *testing.T) {
	qr := NewQrCode("https://passport.aliyundrive.com/qrcodeCheck.htm?lgToken=10254f1b2c45b306ce7606b6893257913_0000000&_from=havana", true)
	//qr.Debug()
	qr.Output()
}
