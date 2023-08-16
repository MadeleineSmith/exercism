package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.Replace(question, "What is ", "", 1)
	question = strings.Replace(question, "?", "", 1)
	question = strings.Replace(question, "multiplied by", "multiplied", -1)
	question = strings.Replace(question, "divided by", "divided", -1)

	parts := strings.Fields(question)

	runningTotal := 0
	operator := ""
	i := 0
	numParts := len(parts)

	for i < numParts {
		part := parts[i]

		if part == "plus" || part == "minus" || part == "multiplied" || part == "divided" {
			if operator != "" {
				return 0, false
			}

			operator = part
			i++
			continue
		}

		num, err := strconv.Atoi(part)
		if err != nil {
			return 0, false
		}

		// extract out below into func:
		switch operator {
		case "":
			if i == 0 {
				runningTotal = num
			} else {
				return 0, false // two num in a row
			}
		case "plus":
			runningTotal += num
		case "minus":
			runningTotal -= num
		case "multiplied":
			runningTotal *= num
		case "divided":
			runningTotal /= num
		default:
			return 0, false
		}

		i++
		operator = ""
	}

	if operator != "" { // we're missing an operand here:
		return 0, false
	}

	return runningTotal, true
}
