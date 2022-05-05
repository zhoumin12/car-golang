package model

var JiaoFei []Jiaofei

type Parking struct {
	Info  []string `json:"info"`
	Local string   `json:"local"`
}

type Jiaofei struct {
	CarNumber   string `json:"car_number"`
	CarMoney    int    `json:"car_money"`
	PaymentArea string `json:"payment_area"`
	PayTime     string `json:"pay_time"`
	StopTime    string `json:"stop_time"`
}

type JiaofeiRes struct {
	Jiaofei []Jiaofei `json:"jiaofei"`
	Times   int       `json:"times"`
	Price   int       `json:"price"`
}

func GetParkingInfo(keyword string) Parking {

	for p, car := range ParkingMap {
		if car == "" {
			info := Parking{Local: p, Info: []string{"1111", "22222", "33333"}}
			ParkingMap[p] = keyword
			return info
		}
	}
	return Parking{Local: ""}
}

func GetParking() []Parking {
	var pk []Parking
	for p, car := range ParkingMap {
		if car != "" {
			info := Parking{Local: p, Info: []string{"1111", "22222", "33333"}}
			pk = append(pk, info)
		}
	}
	return pk
}

func GetMyJiaofei() JiaofeiRes {
	var jiaoFeiRes JiaofeiRes
	var sum int
	for _, v := range JiaoFei {
		sum += v.CarMoney
	}
	jiaoFeiRes.Jiaofei = JiaoFei
	jiaoFeiRes.Times = len(JiaoFei)
	jiaoFeiRes.Price = sum
	return jiaoFeiRes
}

func GetSum() int {
	var sum int
	for _, car := range ParkingMap {
		if car == "" {
			sum++
		}
	}
	return sum
}
