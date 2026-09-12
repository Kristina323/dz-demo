//Описание: Создайте программу на Go, которая объявляет две строковые переменные и выводит их объединение (конкатенацию) в консоль.
//Входные данные: Нет (все данные встроены в код)
//Выходные данные: Объединённая строка, выведенная в консоль
//Ограничения: Используйте объявление двух строковых переменных в одной строке с помощью var и выполните их конкатенацию

//Примеры:
//Output: greeting = Hello World
//Output: greeting = Good Morning

/*package main

import "fmt"

	func main() {
		a := "Я"
		b := "Учу"
		c := "Go"
		d := a + " " + b + " " + c + " " + "!"
		fmt.Println(d)
	}

/*package main

import "fmt"

func main() {
	var a, b string = "Hello World", "Good Morning"
	fmt.Println(a + " " + b + "!")
}*/

//Задание 2:
//Объявите переменную типа int с помощью var
//Объявите переменную типа float64 с помощью :=
//Выполните явное преобразование типа int к float64 перед сложением
//Сохраните результат в отдельную переменную и выведите его
//Примеры:

//Output: result = 8.5
//Output: result = 11.2

/*
package main

import "fmt"

	func main() {
		var a int = 5
		b := 3.5
		sum := float64(a) + b
		fmt.Println("result =", sum)
	}
*/
package main

import "fmt"

func main() {
	var a int = 7
	b := 4.2
	sum := float64(a) + b
	fmt.Println("result =", sum)
}
