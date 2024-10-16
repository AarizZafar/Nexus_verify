package model

type NetBioMetrix struct {
	SSID                   string `json:"SSID,omitempty"`
	BSSID                  string `json:"BSSID,omitempty"`
	PublicIPAdd            string `json:"PublicIPAdd,omitempty"`
	SubNetMask 			   string `json:"SubNetMask,omitempty"`
	IPV4_DG                string `json:"IPV4_DG,omitempty"`
	IPV6_DG                string `json:"IPV6_DG,omitempty"`
	Active_MAC             string `json:"Active_MAC,omitempty"`
	Inactive_MAC           string `json:"Inactive_MAC,omitempty"`
	Security_proto         string `json:"Security_proto,omitempty"`
}

type SysBioMetrix struct {
	SSID                   string `json:"SSID,omitempty"`
	MAC                    string `json:"MAC,omitempty"`
	CPUSerial              string `json:"CPUSerial,omitempty"`
	HardDriveSerial        string `json:"HardDriveSerial,omitempty"`
	MotherBoardSerial      string `json:"MotherBoardSerial,omitempty"`
	BIOSSerial             string `json:"BIOSSerial,omitempty"`
	SSDSerial              string `json:"SSDSerial,omitempty"`
	TPMChipID              string `json:"TPMChipID,omitempty"`
	RAMSerial              string `json:"RAMSerial,omitempty"`
	GPUSerial              string `json:"GPUSerial,omitempty"`
	NICID                  string `json:"NICID,omitempty"`
	BaseBoardProduct       string `json:"BaseBoardProduct,omitempty"`
	SystemUUID             string `json:"SystemUUID,omitempty"`
	OSInstallationID       string `json:"OSInstallationID,omitempty"`
	DiskVolumeSerialNumber string `json:"DiskVolumeSerialNumber,omitempty"`
	BootROMVersion         string `json:"BootROMVersion,omitempty"`
	GPUVendorID            string `json:"GPUVendorID,omitempty"`
	DeviceTreeIdentifier   string `json:"DeviceTreeIdentifier,omitempty"`
	UEFIFirmwareVersion    string `json:"UEFIFirmwareVersion,omitempty"`
}