package model

type NetBioMetrix struct {
	SSID            string `json:"SSID,omitempty"`
	BSSID           string `json:"BSSID,omitempty"`
	SubNetMask      string `json:"SubNetMask,omitempty"`
	NetInterfaceMAC string `json:"NetInterfaceMAC,omitempty"`
}

type SysBioMetrix struct {
	SSID               string `json:"SSID,omitempty"`
	MAC                string `json:"MAC,omitempty"`
	SystemSerialNumber string `json:"SystemSerialNumber,omitempty"`
	UUID               string `json:"UUID,omitempty"`
}
