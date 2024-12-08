package main

import (
	"math"
	"strconv"
)

func generateOperators(operators []string, count int) [][]string {
	comboCount := int(math.Pow(float64(len(operators)), float64(count)))
	//fmt.Println("Total combos:", comboCount)
	operatorSets := make([][]string, comboCount)
	for i := 0; i < count; i++ {
		for j := 0; j < comboCount; j++ {
			threshold := int(math.Pow(float64(len(operators)), float64(count-i-1)))
			operatorSets[j] = append(operatorSets[j], operators[int(j/threshold)%len(operators)])
		}
	}
	return operatorSets
}

func assemble(nums []string, sets [][]string) [][]string {
	equations := [][]string{}
	for _, set := range sets {
		equation := []string{}
		nIdx := 0
		oIdx := 0
		for i := 0; i < len(nums)+len(set); i++ {
			if i%2 == 0 { // numbers at even positions
				m := nums[nIdx]
				nIdx++
				equation = append(equation, m)
			} else { // operators at odd positions
				m := set[oIdx]
				oIdx++
				equation = append(equation, m)
			}
		}
		equations = append(equations, equation)
	}
	return equations
}

func check(expected string, equations [][]string, opsMap map[string]func(int, int) int) (bool, int) {
	for _, eqn := range equations {
		//fmt.Println(eqn)
		result, err := strconv.Atoi(eqn[0])
		if err != nil {
			panic(err)
		}
		for i := 1; i < len(eqn); i += 2 {
			rhs, err := strconv.Atoi(eqn[i+1])
			if err != nil {
				panic(err)
			}
			result = opsMap[eqn[i]](result, rhs)
		}
		xpected, err := strconv.Atoi(expected)
		if err != nil {
			panic(err)
		}
		if result == xpected {
			return true, result
		}
	}
	return false, 0
}
