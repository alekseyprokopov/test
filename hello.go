package main

import "fmt"

func main() {
	//type Person struct {
	//	Name        string `json:"Имя"`
	//	Email       string `json:"Почта"`
	//	DateOfBirth time.Time
	//}
	//person := Person{
	//	Name:  "Алекс",
	//	Email: "alex@yandex.ru",
	//}
	//jsPerson, _ := json.Marshal(person)

	//fmt.Println(unintuitive())
	//fmt.Println(intuitive())

	var Global = 5

	defer func() {
		Global = 10
		fmt.Println(Global)
		Global = 5
	}()

}

func find(arr []int, k int) []int {
	storage := make(map[int]int)

	for i, v := range arr {
		if j, ok := storage[k-v]; ok {
			return []int{i, j}
		}
		storage[v] = i
	}
	return nil

}

func RemoveDuplicates(input []string) []string {
	storage := make(map[string]int)
	output := make([]string, 0)
	for _, v := range input {
		if _, ok := storage[v]; !ok {
			output = append(output, v)
		}
		storage[v]++

	}
	return output

}

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}
func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}
