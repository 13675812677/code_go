package model

type ResourceCar struct {
//GENERATE_START
	 
	CarId string`json:"carId"`
	 
	CarNo string`json:"carNo"`
	 
	PlateNum string`json:"plateNum"`
	 
	CarType int`json:"carType"`
	 
	Status int`json:"status"`
	 
	CarDesc string`json:"carDesc"`
	
	
	DataType string `json:"dataType"`
	
//GENERATE_END
}