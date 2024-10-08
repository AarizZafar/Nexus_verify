package model

// import "go.mongodb.org/mongo-driver/bson/primitive"

type SysBioMetrix struct {
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
