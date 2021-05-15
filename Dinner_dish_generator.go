package Dinner_dish_generator

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main()  {
	food := [15]string{
		"Суши",
		"Пицца",
		"Паста Карбонара",
		"Паста с морепродуктами",
		"Лазанья",
		"Курочка в соевом соусе",
		"Чебуречки",
		"Зразы",
		"Тефтельки",
		"Пельмешки",
		"Плов",
		"Котлетки",
		"Домашние бургеры",
		"Колбаски",
		"Макароны по-флотски",
	}

	var point string
	count := 1
	fla := true

	for fla == true{
		fmt.Println("1. Сгенерировать блюдо \n2. Выйти")
		fmt.Println("Выберите пункт: ")
		fmt.Scan(&point)

		if point == "1" {
			if count <= 2{
				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(len(food))

				fmt.Println("Вы можете сгенерировать ещё раз, но результат второй генерации вы должны принять!")
				fmt.Println("Попытка номер: ", count)
				fmt.Println("Ваш ужин: ", food[n])
				count ++
			} else {
				fmt.Print("Попыток больше нет! До встречи!")
				fla = false
			}
		}

		if point == "2" {
			fmt.Println("До встречи!")
			os.Exit(0)
		}
	}
}
