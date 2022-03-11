package qrcode

import (
	"AliyunPanToken/qrcode/model"
	"AliyunPanToken/util/vjson"
	"errors"
	"github.com/zf1976/vlog"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Api struct {
}

func (api Api) GetQrCodeContent() (*model.QrCodeCK, error) {
	var globalErr error
	get, globalErr := http.Get("https://passport.aliyundrive.com/newlogin/qrcode/generate.do?appName=aliyun_drive&isMobile=true")
	if globalErr != nil {
		vlog.Errorf("get qrcode error: %v", globalErr.Error())
		return nil, globalErr
	}
	body := get.Body
	//goland:noinspection GoUnhandledErrorResult
	defer body.Close()

	bytes, globalErr := ioutil.ReadAll(body)
	q := model.QrCodeResult{}
	globalErr = vjson.ByteArrayConvert(bytes, &q)
	if globalErr != nil {
		vlog.Errorf("convert body error: %v", globalErr)
		return nil, globalErr
	}
	vlog.Debugf("qrcode content result json:\n%v", vjson.PrettifyString(q))
	content := q.Content
	if content.Success {
		result := model.QrCodeCK{
			T:           strconv.FormatInt(content.Data.T, 10),
			CodeContent: content.Data.CodeContent,
			CK:          content.Data.Ck,
		}
		vlog.Debugf("qrcode content and ck for json:\n%v", vjson.PrettifyString(result))
		return &result, nil
	}
	vlog.Debug(content.Data.TitleMsg)
	return nil, errors.New(content.Data.TitleMsg)
}

func QueryQrCode() {}
