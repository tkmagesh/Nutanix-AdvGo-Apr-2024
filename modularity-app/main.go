package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tkmagesh/nutanix-advgo-apr-2024/modularity-app/calculator"
	"github.com/tkmagesh/nutanix-advgo-apr-2024/modularity-app/calculator/utils"
)

func main() {
	color.Red("modularity app executed")
	greet("Magesh")

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.OpCount())

	fmt.Println(utils.IsEven(21))
}
