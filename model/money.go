package model

var Money int

type MoneyCard struct {
	Name string `json:"car_number"`
	Sum  int    `json:"car_money"`
}

type MyCar struct {
	Vcid      string `json:"vcid"`
	Carnumber string `json:"carnumber"`
	Opdate    string `json:"opdate"`
}

type MyCarRes struct {
	MyCar []MyCar `json:"cars_number"`
}

func GetMyMoney() MoneyCard {
	return MoneyCard{Name: "余额", Sum: Money}
}
