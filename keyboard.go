package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloat takes a float from stin and returns it to a variable.
func GetFloat() (float64, error) {
	myReader := bufio.NewReader(os.Stdin)
	input, err := myReader.ReadString('\n')

	//Validates that error is not returned
	if err != nil {
		return 0, err
	}

	//Trims the space out of the input from standard in.
	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
