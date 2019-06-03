package arn

import (
	"net/http"

	webpush "github.com/akyoto/webpush-go"
	jsoniter "github.com/json-iterator/go"
)

// PushSubscription ...
type PushSubscription struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"userAgent"`
	Screen    struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"screen"`
	Endpoint    string `json:"endpoint" private:"true"`
	P256DH      string `json:"p256dh" private:"true"`
	Auth        string `json:"auth" private:"true"`
	Created     string `json:"created"`
	LastSuccess string `json:"lastSuccess"`
}

// ID ...
func (sub *PushSubscription) ID() string {
	return sub.Endpoint
}

// SendNotification ...
func (sub *PushSubscription) SendNotification(notification *PushNotification) (*http.Response, error) {
	// Define endpoint and security tokens
	s := webpush.Subscription{
		Endpoint: sub.Endpoint,
		Keys: webpush.Keys{
			P256dh: sub.P256DH,
			Auth:   sub.Auth,
		},
	}

	// Create notification
	data, err := jsoniter.Marshal(notification)

	if err != nil {
		return nil, err
	}

	// Send Notification
	return webpush.SendNotification(data, &s, &webpush.Options{
		Subscriber:      APIKeys.VAPID.Subject,
		TTL:             60,
		VAPIDPrivateKey: APIKeys.VAPID.PrivateKey,
	})
}
