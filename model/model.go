package model

// import "go.mongodb.org/mongo-driver/bson/primitive"

type SysBioMetrix struct {
    Test12            string `json:"test12,omitempty"`
    MAC               string `json:"MAC,omitempty"`
    CPUSerial         string `json:"CPUSerial,omitempty"`
    HardDriveSerial   string `json:"HardDriveSerial,omitempty"`
    MotherBoardSerial string `json:"MotherBoardSerial,omitempty"`
    BIOSSerial        string `json:"BIOSSerial,omitempty"`
    SSDSerial         string `json:"SSDSerial,omitempty"`
    TPMChipID         string `json:"TPMChipID,omitempty"`
    RAMSerial         string `json:"RAMSerial,omitempty"`
    GPUSerial         string `json:"GPUSerial,omitempty"`
    NICID             string `json:"NICID,omitempty"`
}

// type SysBioMetrix struct {
//     MAC               string `json:"mac,omitempty" bson:"mac,omitempty"`
//     CPUSerial         string `json:"cpuserial,omitempty" bson:"cpuserial,omitempty"`
//     HardDriveSerial   string `json:"harddriveserial,omitempty" bson:"harddriveserial,omitempty"`
//     MotherBoardSerial string `json:"motherboardserial,omitempty" bson:"motherboardserial,omitempty"`
//     BIOSSerial        string `json:"biosserial,omitempty" bson:"biosserial,omitempty"`
//     SSDSerial         string `json:"ssdserial,omitempty" bson:"ssdserial,omitempty"`
//     TPMChipID         string `json:"tpmchipid,omitempty" bson:"tpmchipid,omitempty"`
//     RAMSerial         string `json:"ramserial,omitempty" bson:"ramserial,omitempty"`
//     GPUSerial         string `json:"gpuserial,omitempty" bson:"gpuserial,omitempty"`
//     NICID             string `json:"nicid,omitempty" bson:"nicid,omitempty"`
// }

