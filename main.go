package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX = 100

type City struct {
	name string
	long int
	lat  int
}

func main() {
	file, err := os.Open("cities.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var cities []City

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		long, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		lat, _ := strconv.Atoi(strings.TrimSpace(parts[2]))

		point := City{parts[0], long, lat}
		cities = append(cities, point)
	}
	//fmt.Printf("%v", cities)
	build(cities)
}

func build(cities []City) {
	degree := len(cities) - 1
	var matrix [][]int

	for j := 0; j < len(cities); j++ {
		equation := make([]int, degree+1)
		for i := 0; i < degree; i++ {
			equation[i] = pow(cities[j].long, degree-i)
		}
		equation[degree] = cities[j].lat
		matrix = append(matrix, equation)
	}
	fmt.Println(matrix)
	result := matMult(matrix, matrix)
	fmt.Println(result)
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func matMult(A, B [][]int) [][]int {
	var result [][]int
	if len(A[0]) != len(B) {
		fmt.Println("ERROR")
		return nil
	}

	for row := 0; row < len(A); row++ {
		line := []int{}
		for col := 0; col < len(B[0]); col++ {
			line = append(line, 0)
			for i := 0; i < len(B[0]); i++ {
				line[col] += A[row][i] * B[i][col]
			}
		}
		result = append(result, line)
	}

	return result
}
