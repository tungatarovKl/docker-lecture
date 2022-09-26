package main

import (
	"log"
	"net/http"
	"fmt"
	"strconv"
)

func main() {
	
	http.HandleFunc("/", handler)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func handler(w http.ResponseWriter, r *http.Request) {
	firstOperand := r.URL.Query().Get("a")
	secondOperand := r.URL.Query().Get("b")
	
	if len(firstOperand) == 0 || len(secondOperand) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Не передан один из операндов а или b"))
		
		return
	}
	
	operation := r.URL.Query().Get("operation")
	if len(operation) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Не передана операция"))
		
		return
	}
	
	firstNormalized, err := normalize(firstOperand)
	secondNormalized, err := normalize(secondOperand)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Ошибка при обработке операндов"))
		
		return
	}
	
	result, err := calculate(firstNormalized, secondNormalized, operation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при вычислений результата"))
		
		return
	}
	
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Результат: %d", result)))
}

func normalize(operand string) (int, error) {
	i, err := strconv.Atoi(operand)
	if err != nil {
		return 0, err
	}
	
	return i, nil
}
