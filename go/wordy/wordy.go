package wordy

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	mathsEquationRegex := regexp.MustCompile(`What is ([\w -]+)\?$`)
	mathsEquationMatches := mathsEquationRegex.FindStringSubmatch(question)
	if mathsEquationMatches == nil {
		return 0, false
	}

	stringOfMaths := mathsEquationMatches[1]

	// put in func start
	// 5-5 matches this regex haha
	numberRegex := regexp.MustCompile(`^([\d-]+)$`)
	numberMatches := numberRegex.FindStringSubmatch(stringOfMaths)
	if numberMatches != nil {
		yah, _ := strconv.Atoi(numberMatches[1])

		return yah, true
	}

	equationRegex := regexp.MustCompile(`^([\d-]+) ([a-z ]+) ([\d-]+)`)
	equationMatches := equationRegex.FindStringSubmatch(stringOfMaths)
	if equationMatches == nil {
		return 0, false
	}

	result, err := doCalculation(equationMatches)
	if err != nil {
		return 0, false
	}

	stringOfMaths = strings.Replace(stringOfMaths, equationMatches[0], strconv.Itoa(result), 1)

	numberMatches = numberRegex.FindStringSubmatch(stringOfMaths)
	if numberMatches != nil {
		return result, true
	}

	equationMatches = equationRegex.FindStringSubmatch(stringOfMaths)
	if equationMatches == nil {
		return 0, false
	}

	result, err = doCalculation(equationMatches)
	if err != nil {
		return 0, false
	}

	return result, true
}

func doCalculation(equationParts []string) (int, error) {
	firstOperand, _ := strconv.Atoi(equationParts[1])
	secondOperand, _ := strconv.Atoi(equationParts[3])

	switch equationParts[2] {
	case "plus":
		return firstOperand + secondOperand, nil
	case "minus":
		return firstOperand - secondOperand, nil
	case "multiplied by":
		return firstOperand * secondOperand, nil
	case "divided by":
		return firstOperand / secondOperand, nil
	default:
		return 0, errors.New("bad operator 😔")
	}
}
