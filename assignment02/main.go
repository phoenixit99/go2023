package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
)

func (f *CustomFlag) String() string {
	return fmt.Sprintf("%v", f.value)
}

func (f *CustomFlag) Set(s string) error {
	if intValue, err := strconv.Atoi(s); err == nil {
		f.value = intValue
		return nil
	} else {
		// Try parsing as float
		if floatValue, err := strconv.ParseFloat(s, 64); err == nil {
			f.value = floatValue
			return nil
		}
		f.value = s
		return nil
	}
}

type CustomFlag struct {
	value interface{}
}

func sortListNumber(flagInt int) {
	numberArray := []int{flagInt}
	for _, arg := range flag.Args() {
		if intValue, err := strconv.Atoi(arg); err == nil {
			numberArray = append(numberArray, intValue)
		} else {
			fmt.Println("Error : Invalid format input")
			return
		}
	}
	sort.Ints(numberArray)
	fmt.Println("--- Sorted Array Integer ---", numberArray)
}

func sortListString(flagStr string) {
	stringArray := []string{flagStr}
	for _, arg := range flag.Args() {
		if _, err := strconv.Atoi(arg); err == nil {
			fmt.Println("Error : Invalid format input")
			return
		} else {
			stringArray = append(stringArray, arg)
		}
	}
	sort.Strings(stringArray)
	fmt.Println("---Sort String Array ---", stringArray)
}

func handleSortInMix(flagMix CustomFlag) {
	numberArray := []interface{}{}
	stringArray := []string{}
	switch flagMix.value.(type) {
	case int:
		numberArray = append(numberArray, flagMix.value)
	case float64:
		numberArray = append(numberArray, flagMix.value)
	case string:
		stringArray = append(stringArray, flagMix.value.(string))
	}
	for _, arg := range flag.Args() {
		if intValue, err := strconv.Atoi(arg); err == nil {
			numberArray = append(numberArray, intValue)
		} else {
			if floatValue, err := strconv.ParseFloat(arg, 64); err == nil {
				numberArray = append(numberArray, floatValue)
			} else {
				stringArray = append(stringArray, arg)
			}
		}
	}
	sortMixedNumbers(numberArray)
	sort.Strings(stringArray)
	for _, str := range stringArray {
		numberArray = append(numberArray, str)
	}
	fmt.Println("---arrayGeneric ---", numberArray)
}

func sortMixedNumbers(numbers []interface{}) {
	sort.Slice(numbers, func(i, j int) bool {
		switch numbers[i].(type) {
		case int:
			if numJ, ok := numbers[j].(int); ok {
				return numbers[i].(int) < numJ
			} else {
				if numJ, ok := numbers[j].(float64); ok {
					return float64(numbers[i].(int)) < numJ
				}
			}
		case float64:
			if numJ, ok := numbers[j].(float64); ok {
				return numbers[i].(float64) < numJ
			} else {
				if numJ, ok := numbers[j].(int); ok {
					return numbers[i].(float64) < float64(numJ)
				}
			}
		}
		return false
	})
}

func main() {

	var flagNum CustomFlag
	flag.Var(&flagNum, "int", " ")

	var flagStr CustomFlag
	flag.Var(&flagStr, "string", " ")

	var flagMix CustomFlag
	flag.Var(&flagMix, "mix", " ")

	flag.Parse()

	flagActions := map[string]func(){
		"int": func() {
			if numJ, ok := flagNum.value.(int); ok {
				sortListNumber(numJ)
			} else {
				fmt.Println("Error : Invalid format input")
			}
		},
		"string": func() {
			if numJ, ok := flagStr.value.(string); ok {
				sortListString(numJ)
			} else {
				fmt.Println("Error : Invalid format input")
			}
		},
		"mix": func() {
			handleSortInMix(flagMix)
		},
	}
	// Iterate through the flags and execute associated functions
	flag.Visit(func(f *flag.Flag) {
		if action, ok := flagActions[f.Name]; ok {
			action()
		}
	})

}
