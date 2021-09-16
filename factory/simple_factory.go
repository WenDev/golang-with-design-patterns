package factory

import "fmt"

func CreateHuman(t string) Human {
	switch t {
	case "b":
		return &BlackHuman{}
	case "y":
		return &YellowHuman{}
	case "w":
		return &WhiteHuman{}
	default:
		fmt.Println("人种生成错误!")
		return nil
	}
}
