package model

type CentrifugoPublishData struct {
	Value string `json:"value"`
}

type CentrifugoPublish struct {
	Channel string                `json:"channel"`
	Data    CentrifugoPublishData `json:"data"`
}

type CentrifugoResponse struct {
}
