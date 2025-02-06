package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Data struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var data = []Data{
	{ID: 1, Name: "Alice", Value: "Value1"},
	{ID: 2, Name: "Bob", Value: "Value2"},
	{ID: 3, Name: "Charlie", Value: "Value3"},
}

func getData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	search := query.Get("search")
	filter := query.Get("filter")
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	var filteredData []Data
	for _, d := range data {
		if (search == "" || strings.Contains(d.Name, search)) &&
			(filter == "" || d.Value == filter) {
			filteredData = append(filteredData, d)
		}
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredData) {
		start = len(filteredData)
	}
	if end > len(filteredData) {
		end = len(filteredData)
	}

	paginatedData := filteredData[start:end]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedData)
}

func RemoveChar(word string) string {
	str := strings.Split(word, "")
	str = str[1 : len(str)-1]
	return strings.Join(str, "")
}

func StringToArray(str string) []string {
	// your code here
	return strings.Split(str, " ")
}

func Movie(card, ticket int, perc float64) int {
	sumB := 0.0
	currentTerm := 0.0
	n := 1

	for {
		if n == 1 {
			currentTerm = float64(ticket) * perc
			sumB = currentTerm
		} else {
			currentTerm *= perc
			sumB += currentTerm
		}

		totalB := float64(card) + sumB
		roundedB := math.Ceil(totalB)
		systemA := ticket * n

		if roundedB < float64(systemA) {
			return n
		}

		n++
	}

}

func ReverseLetters(s string) string {
	var letters []rune
	for _, c := range s {
		if isLetter(c) {
			letters = append(letters, c)
		}
	}

	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	return string(letters)
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
    area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
    return area
}


func Arithmetic(a int, b int, operator string) int{
  //your code here
   switch operator {
    case "add":
        return a + b
    case "subtract":
        return a - b
    case "multiply":
        return a * b
    case "divide":
        return a / b
    default:
        return 0  
    }
}

func main() {
	// http.HandleFunc("/data", getData)
	// http.ListenAndServe(":8080", nil)
	// fmt.Println(RemoveChar("word"))

	// fmt.Println(StringToArray("word komik anime"))
	// fmt.Println(ReverseLetters("krishan"))
	// fmt.Println(ReverseLetters("ultr53o?n"))

	fmt.Println(Heron(18.00, 6.00, 13.00));
}
