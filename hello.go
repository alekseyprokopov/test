package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type Person struct {
		Name        string `json:"Имя"`
		Email       string `json:"Почта"`
		DateOfBirth time.Time
	}
	person := Person{
		Name:  "Алекс",
		Email: "alex@yandex.ru",
	}
	jsPerson, _ := json.Marshal(person)

	fmt.Println(string(jsPerson))
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
