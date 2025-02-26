package notification

type HarmonyPayload struct {
	DisplayType string       `json:"display_type"`
	Body        HarmonyBody  `json:"body"`
	Extra       HarmonyExtra `json:"extra,omitempty"`
}

func (p *HarmonyPayload) Initial() *HarmonyPayload {
	p.Extra = make(HarmonyExtra)
	p.DisplayType = "notification"
	return p
}

func (p *HarmonyPayload) AddExtra(key string, val string) *HarmonyPayload {
	p.Extra.AddKV(key, val)
	return p
}

func (p *HarmonyPayload) SetDisplayType(displayType string) *HarmonyPayload {
	if displayType == "message" {
		p.DisplayType = displayType
	} else {
		p.DisplayType = "notification"
	}
	return p
}

type HarmonyExtra map[string]string

func (e *HarmonyExtra) AddKV(key string, val string) *HarmonyExtra {
	(*e)[key] = val
	return e
}

type HarmonyBody struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	BigBody   string `json:"big_body,omitempty"`
	LargeIcon string `json:"largeIcon,omitempty"`
	AddBadge  string `json:"add_badge,omitempty"`
	AfterOpen string `json:"after_open,omitempty"`
	Uri       string `json:"uri"`
	Action    string `json:"action"`
	Custom    string `json:"custom"`
}

func (b *HarmonyBody) SetTitle(title string) *HarmonyBody {
	b.Title = title
	return b
}

func (b *HarmonyBody) SetText(text string) *HarmonyBody {
	b.Text = text
	return b
}

func (b *HarmonyBody) SetBigBody(bigbody string) *HarmonyBody {
	b.BigBody = bigbody
	return b
}

func (b *HarmonyBody) SetLargeIcon(icon string) *HarmonyBody {
	b.LargeIcon = icon
	return b
}

func (b *HarmonyBody) SetAddBadge(badge string) *HarmonyBody {
	b.AddBadge = badge
	return b
}

func (b *HarmonyBody) SetAfterOpen(action string) *HarmonyBody {
	b.AfterOpen = action
	return b
}

func (b *HarmonyBody) SetUri(uri string) *HarmonyBody {
	b.Uri = uri
	return b
}

func (b *HarmonyBody) SetAction(action string) *HarmonyBody {
	b.Uri = action
	return b
}

func (b *HarmonyBody) SetCustom(custom string) *HarmonyBody {
	b.Custom = custom
	return b
}
