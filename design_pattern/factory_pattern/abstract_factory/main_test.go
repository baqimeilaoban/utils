package abstract_factory

import (
	"fmt"
	"testing"
)

func Test_abc(t *testing.T) {
	// 生成具体的对象
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	// 根据抽象方法获取具体的产品对象
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
