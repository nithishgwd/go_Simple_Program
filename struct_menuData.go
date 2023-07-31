package main

type menuItems struct {
	name   string
	prices map[string]float64
}

func menuData() []menuItems {
	menu := []menuItems{
		{name: "coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "espresso", prices: map[string]float64{"single": 2.1, "double": 2.7, "triple": 3.4}},
	}

	return menu
}
