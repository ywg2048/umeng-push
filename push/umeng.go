package push

import (
	"strings"
)

const (
	Host        string = "https://msgapi.umeng.com"
	SendPath    string = "/api/send"
	StatusPath  string = "/api/status"
	ChannelPath string = "/api/channel/data"
	QuotaPath   string = "/api/quota/query"
	CancelPath  string = "/api/cancel"
	UploadPath  string = "/upload"

	TagAddPath    string = "/api/tag/add"
	TagListPath   string = "/api/tag/list"
	TagSetPath    string = "/api/tag/set"
	TagDeletePath string = "/api/tag/delete"
	TagClearPath  string = "/api/tag/clear"

	TmplAddPath    string = "/api/template/add"
	TmplDeletePath string = "/api/template/delete"
	TmplGetPath    string = "/api/template/get"
	TmplListPath   string = "/api/template/list"
	TmplSendPath   string = "/api/template/send"
	TmplMsgPath    string = "/api/template/msg"

	IOS     string = "ios"
	Android string = "android"
)

var Proxy string

func UseProxy(addr string) {
	Proxy = addr
}

func UnsetProxy() {
	UseProxy("")
}

type Umeng struct {
	Android *Client
	IOS     *Client
}

func NewUmeng() *Umeng {
	return &Umeng{
		Android: &Client{Platform: Android},
		IOS:     &Client{Platform: IOS},
	}
}

func (u *Umeng) InitAndroid(appkey string, secret string) *Umeng {
	u.Android.AppKey = appkey
	u.Android.MasterSecret = secret
	return u
}

func (u *Umeng) InitIOS(appkey string, secret string) *Umeng {
	u.IOS.AppKey = appkey
	u.IOS.MasterSecret = secret
	return u
}

func (u *Umeng) Debug(debug bool) *Umeng {
	u.Android.Debug = debug
	u.IOS.Debug = debug
	return u
}

func (u *Umeng) GetClient(platform string) *Client {
	platform = strings.ToLower(platform)
	if platform == IOS {
		return u.IOS
	} else if platform == Android {
		return u.Android
	}
	return &Client{}
}

func (u *Umeng) UseProxy(addr string) {
	UseProxy(addr)
}

func (u *Umeng) Send(request CastRequester) (taskOrMsgId string, err error) {
	return u.GetClient(request.GetPlatform()).Send(request)
}
