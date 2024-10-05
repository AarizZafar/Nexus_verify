package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SysBioMetrix struct {
	ID                primitive.ObjectID       `json:"_id                      ,omitempty"                    bson:"_id,omitempty"`
	MAC               string                   `json:"MAC                      ,omitempty"`
	CPUSerial         string				   `json:"CPUSerial                ,omitempty"`
	HardDriveSerial   string				   `json:"HardDriveSerial          ,omitempty"`
	MotherBoardSerial string		    	   `json:"MotherBoardSerial        ,omitempty"`
	BIOSSerial        string				   `json:"BIOSSerial               ,omitempty"`
	SSDSerial         string				   `json:"SSDSerial                ,omitempty"`
	TPMChipID         string				   `json:"TPMChipID                ,omitempty"`
	RAMSerial         string				   `json:"RAMSerial                ,omitempty"`
	GPUSerial         string				   `json:"GPUSerial                ,omitempty"`
	NICID             string				   `json:"NICID                    ,omitempty"`
}