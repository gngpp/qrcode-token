package qrcode

import (
	"AliyunPanToken/qrcode/model"
	"AliyunPanToken/util/vhttp"
	"AliyunPanToken/util/vjson"
	"errors"
	"github.com/zf1976/vlog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const (
	CONFIRMED = "CONFIRMED"
	EXPIRED   = "EXPIRED"
	NEW       = "NEW"
)

type Api struct {
	qrCodeCK     *model.QrCodeCK
	qrCodeResult *model.QueryQrCodeResult
	generateMux  sync.Mutex
	queryMux     sync.Mutex
}

func (_this *Api) GetQrCodeContent() (*model.QrCodeCK, error) {
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
	q := model.QrCodeGenerateResult{}
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
		_this.generateMux.Lock()
		_this.qrCodeCK = &result
		_this.generateMux.Unlock()
		vlog.Debugf("qrcode content and ck for json:\n%v", vjson.PrettifyString(result))
		return _this.qrCodeCK, nil
	}
	vlog.Debug(content.Data.TitleMsg)
	return nil, errors.New(content.Data.TitleMsg)
}

func (_this Api) GetQrCodeCK() *model.QrCodeCK {
	return _this.qrCodeCK
}

func (_this *Api) QueryQrCode() (*model.QueryQrCodeResult, bool) {
	values := url.Values{}
	_this.generateMux.Lock()
	values.Add("t", _this.qrCodeCK.T)
	values.Add("ck", _this.qrCodeCK.CK)
	_this.generateMux.Unlock()

	ticker := time.NewTicker(time.Second)
	q := &model.QueryQrCodeResult{}
	for {
		<-ticker.C
		// 默认keep-alive
		response, err := vhttp.Post("https://passport.aliyundrive.com/newlogin/qrcode/query.do?appName=aliyun_drive&fromSite=52&_bx-v=2.0.31", "application/x-www-form-urlencoded", values)
		if err != nil {
			vlog.Debugf("query qrcode request error:\n%v", err)
			return nil, false
		}
		var globalErr error
		body := response.Body
		bytes, err := ioutil.ReadAll(body)
		vlog.Debugf("query qrcode row json result:\n%v", string(bytes))

		_this.queryMux.Lock()
		globalErr = vjson.ByteArrayConvert(bytes, q)
		_this.queryMux.Unlock()

		if globalErr != nil {
			vlog.Errorf("convert body error:\n%v", globalErr)
			return nil, false
		}
		vlog.Debugf("struct:\n%v", vjson.PrettifyString(q))
		if q.Content.Success {
			if q.Content.Data.QrCodeStatus == CONFIRMED {
				ticker.Stop()
				break
			}
		}
	}
	return q, true
}
