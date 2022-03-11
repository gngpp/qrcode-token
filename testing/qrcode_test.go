package testing

import (
	"github.com/mdp/qrterminal/v3"
	"os"
	"testing"
)

func TestQrCode(t *testing.T) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}
