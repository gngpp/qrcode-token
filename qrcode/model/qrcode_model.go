package model

type QrCodeResult struct {
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
