package calculator

import (
	"fmt"
)

func CalcOperation(operator string, left float64, right float64) (float64, error) {
	switch operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func Calc(expression string) (float64, error) {
	operators := []string{}
	numbers := []float64{}

	operatorPriority := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, char := range expression {
		if char == ' ' {
			continue
		}
		if char >= '0' && char <= '9' {
			numbers = append(numbers, float64(char-'0'))
			continue
		}

		if char == '(' {
			operators = append(operators, string(char))
			continue
		}

		if char == ')' {
			flag := false
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				flag = true
				operator := operators[len(operators)-1]
				operators = operators[:len(operators)-1]

				if len(numbers) < 2 {
					return 0, fmt.Errorf("invalid expression")
				}

				right := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				left := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]

				result, err := CalcOperation(operator, left, right)
				if err != nil {
					return 0, err
				}

				numbers = append(numbers, result)
			}

			if len(operators) == 0 {
				return 0, fmt.Errorf("invalid expression")
			}

			if operators[len(operators)-1] == "(" {
				operators = operators[:len(operators)-1]
			} else {
				return 0, fmt.Errorf("invalid expression")
			}
			if !flag {
				return 0, fmt.Errorf("invalid expression")
			}
			continue
		}

		if char == '+' || char == '-' || char == '*' || char == '/' {
			currentOp := string(char)
			for len(operators) > 0 && operators[len(operators)-1] != "(" &&
				operatorPriority[operators[len(operators)-1]] >= operatorPriority[currentOp] {
				operator := operators[len(operators)-1]
				operators = operators[:len(operators)-1]

				if len(numbers) < 2 {
					return 0, fmt.Errorf("invalid expression")
				}

				right := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				left := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]

				result, err := CalcOperation(operator, left, right)
				if err != nil {
					return 0, err
				}
				numbers = append(numbers, result)
			}

			operators = append(operators, currentOp)
			continue
		}

		return 0, fmt.Errorf("invalid expression")
	}

	for len(operators) > 0 {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		if operator == "(" || operator == ")" {
			return 0, fmt.Errorf("invalid expression")
		}

		if len(numbers) < 2 {
			return 0, fmt.Errorf("invalid expression")
		}

		right := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		left := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]

		result, err := CalcOperation(operator, left, right)
		if err != nil {
			return 0, err
		}

		numbers = append(numbers, result)
	}

	if len(numbers) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return numbers[0], nil
}
