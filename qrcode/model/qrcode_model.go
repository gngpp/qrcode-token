package model

import "time"

type QrCodeGenerateResult struct {
	Content struct {
		Data struct {
			T           int64  `json:"t,omitempty"`
			CodeContent string `json:"codeContent,omitempty"`
			Ck          string `json:"ck,omitempty"`
			ResultCode  int    `json:"resultCode,omitempty"`
			TitleMsg    string `json:"titleMsg,omitempty"`
			TraceId     string `json:"traceId,omitempty"`
			ErrorCode   string `json:"errorCode,omitempty"`
			IsMobile    bool   `json:"isMobile,omitempty"`
		} `json:"data"`
		Status  int  `json:"status"`
		Success bool `json:"success"`
	} `json:"content"`
	HasError bool `json:"hasError"`
}

type QrCodeCK struct {
	T           string
	CodeContent string
	CK          string
}

type QueryQrCodeResult struct {
	Content struct {
		Data struct {
			LoginResult          string `json:"loginResult,omitempty"`
			LoginSucResultAction string `json:"loginSucResultAction,omitempty"`
			St                   string `json:"st,omitempty"`
			QrCodeStatus         string `json:"qrCodeStatus,omitempty"`
			LoginType            string `json:"loginType,omitempty"`
			BizExt               string `json:"bizExt,omitempty"`
			LoginScene           string `json:"loginScene,omitempty"`
			ResultCode           int    `json:"resultCode,omitempty"`
			AppEntrance          string `json:"appEntrance,omitempty"`
			Smartlock            bool   `json:"smartlock,omitempty"`
		} `json:"data,omitempty"`
		Status  int  `json:"status,omitempty"`
		Success bool `json:"success,omitempty"`
	} `json:"content,omitempty"`
	HasError bool `json:"hasError,omitempty"`
}

type LoginResult struct {
	PdsLoginResult struct {
		Role     string `json:"role,omitempty"`
		UserData struct {
			DingDingRobotUrl string `json:"DingDingRobotUrl,omitempty"`
			EncourageDesc    string `json:"EncourageDesc,omitempty"`
			FeedBackSwitch   bool   `json:"FeedBackSwitch,omitempty"`
			FollowingDesc    string `json:"FollowingDesc,omitempty"`
			BackUpConfig     struct {
				ʖ뺱距 struct {
					FolderId      string `json:"folder_id,omitempty"`
					PhotoFolderId string `json:"photo_folder_id,omitempty"`
					SubFolder     struct {
					} `json:"sub_folder,omitempty"`
					VideoFolderId string `json:"video_folder_id,omitempty"`
				} `json:"ʖ뺱距,omitempty"`
			} `json:"back_up_config,omitempty"`
			DingDingRobotUrl1 string `json:"ding_ding_robot_url,omitempty"`
			EncourageDesc1    string `json:"encourage_desc,omitempty"`
			FeedBackSwitch1   bool   `json:"feed_back_switch,omitempty"`
			FollowingDesc1    string `json:"following_desc,omitempty"`
		} `json:"userData,omitempty"`
		IsFirstLogin   bool          `json:"isFirstLogin,omitempty"`
		NeedLink       bool          `json:"needLink,omitempty"`
		LoginType      string        `json:"loginType,omitempty"`
		NickName       string        `json:"nickName,omitempty"`
		NeedRpVerify   bool          `json:"needRpVerify,omitempty"`
		Avatar         string        `json:"avatar,omitempty"`
		AccessToken    string        `json:"accessToken,omitempty"`
		UserName       string        `json:"userName,omitempty"`
		UserId         string        `json:"userId,omitempty"`
		DefaultDriveId string        `json:"defaultDriveId,omitempty"`
		ExistLink      []interface{} `json:"existLink,omitempty"`
		ExpiresIn      int           `json:"expiresIn,omitempty"`
		ExpireTime     time.Time     `json:"expireTime,omitempty"`
		RequestId      string        `json:"requestId,omitempty"`
		DataPinSetup   bool          `json:"dataPinSetup,omitempty"`
		State          string        `json:"state,omitempty"`
		TokenType      string        `json:"tokenType,omitempty"`
		DataPinSaved   bool          `json:"dataPinSaved,omitempty"`
		RefreshToken   string        `json:"refreshToken,omitempty"`
		Status         string        `json:"status,omitempty"`
	} `json:"pds_login_result,omitempty"`
}
