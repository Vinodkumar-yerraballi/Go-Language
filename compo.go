// Compostition

package main

import (
	"fmt"
)

type Vehicle struct {
	Name string
	Type string
}

type Car struct {
	Vehicle
	Maxspeed float32
	Fueltype string
}

type Movie struct {
	Actor    string
	Actors   string
	Director string
}

type AllMovie struct {
	Movie
	MusicDirector string
	SideChartecrs string
	Budget        int
}

type AgriCluture struct {
	CropName string
	LandName string
}

type Crop struct {
	AgriCluture
	Temparture  float32
	WaterSupply string
}

func main() {
	c := Car{}
	c.Name = "Ferrari"
	c.Type = "Covertable"
	c.Maxspeed = 250
	c.Fueltype = "Petrol"
	fmt.Println(c)
	a := AllMovie{}
	a.Actor = "Thalapathy vijay"
	a.Actors = "Kerrti"
	a.Director = "Atlee"
	a.SideChartecrs = "Satyaraju"
	a.Budget = 500
	fmt.Println(a)
	b := Crop{
		AgriCluture: AgriCluture{CropName: "Paddy", LandName: "Redsoil"},
		Temparture:  25.5,
		WaterSupply: "Yes",
	}
	fmt.Println(b)

}
