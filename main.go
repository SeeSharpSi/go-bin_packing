package main

import (
	"fmt"

	"github.com/gedex/bp3d"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p := bp3d.NewPacker()

	p.AddBin(bp3d.NewBin("Small Bin", 10, 15, 20, 100))
	p.AddBin(bp3d.NewBin("Medium Bin", 100, 150, 200, 1000))

	p.AddItem(bp3d.NewItem("Item 1", 2, 2, 1, 2))
	p.AddItem(bp3d.NewItem("Item 2", 3, 3, 2, 3))

	err := p.Pack()
	check(err)

    displayPacked(p.Bins)
}

func displayPacked(bins []*bp3d.Bin) {
	for _, b := range bins {
		fmt.Println(b)
		fmt.Println(" packed items:")
		for _, i := range b.Items {
			fmt.Println("  ", i)
		}
	}
}
