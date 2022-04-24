package model

import "time"

var CarMap = make(map[string]MyCar)
var ParkingMap = make(map[string]string)

//初始化数据
func InitCar() {
	car1 := MyCar{Vcid: "赣B11332", Carnumber: "赣B11332", Opdate: "2022年1月30日"}
	car2 := MyCar{Vcid: "赣B41332", Carnumber: "赣B11332", Opdate: "2022年1月30日"}
	CarMap["赣B11332"] = car1
	CarMap["赣B41332"] = car2

	ParkingMap["3排1号"] = ""
	ParkingMap["1排1号"] = ""
	ParkingMap["2排15号"] = ""
	ParkingMap["6排1号"] = ""
	ParkingMap["34排145号"] = ""
	ParkingMap["21排5号"] = ""
}

//查看我的爱车
func GetMyCar() MyCarRes {
	var res MyCarRes
	var cars []MyCar
	for _, v := range CarMap {
		cars = append(cars, v)
	}
	res = MyCarRes{MyCar: cars}
	return res
}

//添加我的爱车
func AddCar(number string) error {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	CarMap[number] = MyCar{
		Vcid:      number,
		Carnumber: number,
		Opdate:    timeStr,
	}
	return nil
}

func DeleteCar(card string) error {
	delete(CarMap, card)
	return nil
}
