package main

import (
	"fmt"
	"github.com/huandu/xstrings"
	newcolor "github.com/z-sector/gopackages/color"
	"github.com/z-sector/gopackages/my_random"
	. "github.com/z-sector/gopackages/wordz" //Добавляем пакет wordz через точку
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")

	fmt.Println(Hello)    //Вызов переменной из пакета wordz
	fmt.Println(Random()) //Вызов функции из пакета wordz

	fmt.Println(my_random.Digit())
	randCity := my_random.City()
	fmt.Println(randCity)
	fmt.Println(xstrings.Shuffle(randCity))
}
