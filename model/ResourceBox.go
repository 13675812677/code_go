package model

type ResourceBox struct {
//GENERATE_START
	 
	BoxId string`json:"boxId"`
	 
	BoxNo string`json:"boxNo"`
	 
	BoxVolume string`json:"boxVolume"`
	 
	BoxStatus int`json:"boxStatus"`
	
	
	DataType string `json:"dataType"`
	
//GENERATE_END
}