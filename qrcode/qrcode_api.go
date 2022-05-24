package qrcode

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"qrcode-token/qrcode/model"
	"qrcode-token/util/vhttp"
	"qrcode-token/util/vjson"
	"strconv"
	"sync"
	"time"

	"github.com/gngpp/vlog"
)

const (
	CONFIRMED = "CONFIRMED"
	EXPIRED   = "EXPIRED"
	NEW       = "NEW"
)

const WebLoginTokenApi = "https://passport.aliyundrive.com/newlogin/qrcode/generate.do?t=1653319021482&ck=1f66964aab8ee79f6296255ee5cffcee&ua=140%23%2Fc2D58y%2BzzFc7zo2LZu%2BC3Sc2l%2BjZAz%2FvRn6U9mmyQz%2BzmWCgBKn2ePuU6nBhawutVit34prDth%2Bbs5bgx%2FwAfvBKb1flp1zzqz5Gpug4bzx9H0za3luzzrb22U3lp1xzieAxoTDLzrrLDc3L6gqzzrbfF9sqjmijDapV4%2Fae%2FseTyUf9DRaZO02%2BpSgrTsMobUl6fA4BBMMb6xHKDH99LLDZcaA6XHNLwES2caq6w0Mo3PFhTl85UgdAvt4fxaYq3ID%2F2bdhJ433TD3rB%2B5hG9bHKWXaT2Vaz0SdDBY0alaQV9pCayyWytcQ0EYjunZPWd2WsD3hsxdSlrorz6KLFuNV2UZVAXLGuPKIvPXRqshuyS20REvlvlCLbHdz1O2F9t%2FHc%2BaANXWevwBTb4iU8QT4yDXtrdb6cilTjEgu1b%2Fz8Fmf4K7dOKkf5KRVDLwQeMGWUzWnWGp%2F8C8nsl%2BikeTh%2FSVP8nD842stu38V3RerJw5tw1sea7Ydfud5GcTpnxRNfzQJT0s%2BFQvoYWwc3a0Jr6KiCLqMypvmkPZaqdR4kWqLDhyaHw9DlxcouUyhhSGXa6yt2nNJ%2BSSgvfpNmCSW2VeN5ZYl8wkPfZyMUHCqrmdL4dFQm4sXMQQsY7D4Mz8QdOHyav8zJouMa6%2F9ZRvWxQh1lXTkHOwxqhjrQPX%2FJcR40LB6jwq19pQkFCWQYJeJZhL%2FlJARWfalDvhoPdGFoHadIURsuMX0XjUdklVBPEotHqQy4jj4cXNY42a4UpNfpdBixp7ETMBYWSwwE0Adz4rgG9PD04vqnG2YkDJ0fx%2BxCAoNXHJl%2FKoTFn65iIbfKQ81zmXQ%2BHJMJCP7uf%2BTRNoHaB4hZcSeC2aqUZGpObHX5NaT1pWpxfwhC%2BSp4QbH1d1RREkxAyEVftPVoxXWvvJvrUX80k0QOS2uz61%2FTJV21EU2slqY8ifXvzYlaFtg9Hgqr8Nywl2KpCkuNzC4uvIFeskt4RqV65LFofjpQowV2xgYjfZTFVj6ucHyTI4Seu9saE6srmgS0xE7CjdomgJXmxy%2FAcy8MEyRrGyzXcaQy5K37u1jPDSRiHRArVszfCh3sy57Zbnn1na1nOLW9Lh5XkjVDzZLnLabACj5R%2Bg2Bt7QzFE4K%2Fo2800WZRfWGjdsoCjXVXj8jj8FdzvjcprquE6zYeYBNNtFRlFZelPA6z69l9CgBOu%2FwZBlR1%3D&appName=aliyun_drive&appEntrance=web&_csrf_token=wn7TiM33I1h6tY9yVsy8E&umidToken=85999cf7ce14f6c63497a2b910daae9b0525262d&isMobile=false&lang=zh_CN&returnUrl=&hsiz=1f66964aab8ee79f6296255ee5cffcee&fromSite=52&bizParams=&navlanguage=zh-CN&navUserAgent=Mozilla%2F5.0%20%28Macintosh%3B%20Intel%20Mac%20OS%20X%2010.15%3B%20rv%3A100.0%29%20Gecko%2F20100101%20Firefox%2F100.0&navPlatform=MacIntel&deviceId=&pageTraceId=21310c9916533190166946737e1104&bx-ua=222!qGI7VMPrymaPok9a29aw5ZG2GCnorZiyxq8vI2YNQ8sVR+ogBhNIPuoIUTcQNxksHu08cH+m3wK5J3oSDS4dJeJgC3do5CyFPfoZuZ1q5Du68n3el1aXmtO9api6W7oy7CE4d31kHeaZwUjQpUQhH3nLldqB9wK1mno+EqzHK/yFpnqAw4WFj5ySjYqlCGlVpHyZIZiBtoWkIZsWwK7mtusCtC4Pt122KggthS8x7SEKQxYj+saNuyDCzQTi9uXxlS4wj2APsQUmn9zidg95kUQmQQMG1AFiJG8qz5GSqd43NmOYRlf2AweSYRsSiQjrb2annMAdg7YEN7z9XBnIjDQENM+5rFODPtaVi7F7cKTkSKPY2zEcgh2DSlBY1/0igKv52T2crKABShSR2YAonXFe9kOWDK6fik02g7GbulPhlFAZLgGb3kDtiNq8wjuh8I2+tnOA5Fab5tvsS7iZgZ9wKYqudhc8H2BdQcZcBjAF1xJcYuZzIMstqevFSe+2uPmu91NOKb6ig1tZUVezVl73vO7u0o9gCiZKPMF/JqIEXpbmy6oq8lCrk9JiVNHneM0Hz5uv3WgOAMgnARKp3XD1JcZiWcfT6kh+dnHWtBeAe3drmBMw9HZeHyuAsU4iPSoe4RdlgnTI/Nf6MZkGnZu2ep1Nke2Z7T9s6gdU9KcSMa0pqNy+1zaTOXmFZpMThhKk2qbY6UXcz+tA/LSg1EjASfJyIQ8iDvIDKApQINqlVyRAknGZkBsX0lQT5IKm+T8kF+gMRRr4lvzd1OVtGZcvX8gkAWNKdE1KZfD/LDVWfXCfJIClbX9IYN6T1Ee8bm8X3m5x4W4mtHkj/SAjjil9YGZZpjVacMeMgSDnK4k8I6oZq2PB4udvIi+4meotau1u9hHjXT+yKxU+gRMavJXmpirOx9dw+r35Mt/XAfYP0SkjnMX9VgMOHJgbEfPLKKqMn+4TaX+eNn1S9FkvbrsLzEm5o3DNnHurZ1xR4gaXRxvYdSxdCasok8/EGlhxURcssTYf33eMzg/4qhVNMjTFcFif2PiRI09ZlLNCsUsNUiu4TAWUydKQ7aI7BKlnfYmhRHNjGD8YAn8PQr55a1jWsDHz4gwPl1znyhOhwRkF5YE7b2z7tbwJCgkrOFcznXQeHj1yNQ0WpAJjgaPGgmhKj6aBv2nM3m3I0HIyoOPBXWoGzmN+ID5jSl3j7BkLosQOxZAz4sZR49vqIyF3xBHYKl5PJnnPWmzrK33MFsEBSXEb1VOHYJoBVzqDnk+KfJTUBFw16quVB0XVHyWDZK6fJIIa2w4uINPLtbzSFSqTsALv9gLlrBSo5FNTg76DgdwZRDzmQALk&bx-umidtoken=T2gA17PwZx1jVnczvyOFzn3Yx48F-Jh-X0VFAIE0w3Fo7AJGQkRc5n09moefCfuPx8A="
const MobileLoginTokenApi = "https://passport.aliyundrive.com/newlogin/qrcode/generate.do?appName=aliyun_drive&isMobile=true"

type Api struct {
	qrCodeCK    *model.QueryQrCodeCKForm
	generateMux sync.Mutex
	queryMux    sync.Mutex
	isMobile    bool
}

func NewApi(isMobile bool) *Api {
	return &Api{
		isMobile: isMobile,
	}
}

func (_this *Api) GetGeneratorQrCodeContent() (*model.QueryQrCodeCKForm, error) {
	var globalErr error
	var resp *http.Response
	if _this.isMobile {
		resp, globalErr = http.Get(MobileLoginTokenApi)
	} else {
		resp, globalErr = http.Get(WebLoginTokenApi)
	}
	if globalErr != nil {
		vlog.Errorf("resp qrcode error: %v", globalErr.Error())
		return nil, globalErr
	}
	body := resp.Body
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			vlog.Errorf("system error: %v\n", err)
		}
	}(body)
	bytes, _ := ioutil.ReadAll(body)
	q := model.GeneratorQrCodeResult{}
	globalErr = vjson.ByteArrayConvert(bytes, &q)
	if globalErr != nil {
		vlog.Errorf("convert body error: %v\n", globalErr)
		return nil, globalErr
	}
	vlog.Debugf("qrcode content result json:\n%v", vjson.PrettifyString(q))

	content := q.Content
	if content.Success {
		result := model.QueryQrCodeCKForm{
			T:           strconv.FormatInt(content.Data.T, 10),
			CodeContent: content.Data.CodeContent,
			CK:          content.Data.Ck,
		}
		_this.generateMux.Lock()
		_this.qrCodeCK = &result
		_this.generateMux.Unlock()
		vlog.Debugf("qrcode content and ck for json:\n%v\n", vjson.PrettifyString(result))
		return _this.qrCodeCK, nil
	}
	vlog.Debug(content.Data.TitleMsg)
	return nil, errors.New(content.Data.TitleMsg)
}

func (_this *Api) GetQrCodeCK() *model.QueryQrCodeCKForm {
	return _this.qrCodeCK
}

func (_this *Api) GetQueryQrCodeResult() (*model.QueryQrCodeResult, bool) {
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

		bytes, _ := ioutil.ReadAll(body)
		vlog.Debugf("query qrcode row json result:\n%v", string(bytes))
		_ = body.Close()
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
