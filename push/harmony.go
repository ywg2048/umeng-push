package push

import (
	"github.com/ywg2048/umeng-push/push/notification"
)

type HarmonyRequest struct {
	notification.Notification
	Platform string                      `json:"-"`
	Payload  notification.HarmonyPayload `json:"payload"`
	Channel  notification.Channel        `json:"channel_properties,omitempty"`
}

func (r *HarmonyRequest) GetPlatform() string {
	return r.Platform
}

func (r *HarmonyRequest) GetRequestUri() string {
	return Host + SendPath
}

func NewHarmonyRequest(cast string) *HarmonyRequest {
	n := &HarmonyRequest{Platform: Harmony}
	n.SetNotificationType(cast)
	n.InitTimestamp()
	n.Payload.Initial()
	n.Policy = &notification.Policy{}
	return n
}

func NewHarmonyUnicastRequest() *HarmonyRequest {
	return NewHarmonyRequest(notification.Unicast)
}

func NewHarmonyListcastRequest() *HarmonyRequest {
	return NewHarmonyRequest(notification.Listcast)
}

func NewHarmonyGroupcastRequest() *HarmonyRequest {
	return NewHarmonyRequest(notification.Groupcast)
}

func NewHarmonyBroadcastRequest() *HarmonyRequest {
	return NewHarmonyRequest(notification.Broadcast)
}

func NewHarmonyFilecastRequest() *HarmonyRequest {
	return NewHarmonyRequest(notification.Filecast)
}

func NewHarmonyCustomizedcastRequest() *HarmonyRequest {
	return NewHarmonyRequest(notification.Customizedcast)
}
