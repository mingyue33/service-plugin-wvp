package model

import "encoding/json"

type DeviceItem struct {
	DeviceNumber string `json:"device_number"`
	DeviceName   string `json:"device_name"`
	Description  string `json:"description"`
}

// 事件/命令
type EventInfo struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

type WvpForm struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	ApiToken string `json:"apiToken"`
}

func (w *WvpForm) MarshalBinary() (data []byte, err error) {
	return json.Marshal(w)
}

func (w *WvpForm) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}
