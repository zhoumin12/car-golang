package model

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

func GetParkingInfo() Parking {
	info := Parking{Local: "三排一号", Info: []string{"1111", "22222", "33333"}}
	return info
}

func GetMyJiaofei() JiaofeiRes {
	var jiaoFeiRes JiaofeiRes
	var jiaofei []Jiaofei
	info1 := Jiaofei{CarMoney: 12, CarNumber: "2022年1月30日", PaymentArea: "支付宝", PayTime: "11:22", StopTime: "12:22"}
	info2 := Jiaofei{CarMoney: 15, CarNumber: "2022年2月12日", PaymentArea: "微信", PayTime: "15:22", StopTime: "16:12"}
	jiaofei = append(jiaofei, info1)
	jiaofei = append(jiaofei, info2)
	var sum int
	for _, v := range jiaofei {
		sum += v.CarMoney
	}
	jiaoFeiRes.Jiaofei = jiaofei
	jiaoFeiRes.Times = len(jiaofei)
	jiaoFeiRes.Price = sum
	return jiaoFeiRes
}