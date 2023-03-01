package main

import (
	"fmt"
	"strings"
)

func main() {

	slc := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//slc := 992
	//slc := 195
	//slc := 16
	res := CreatePhoneNumber(slc)

	fmt.Println(res)

}

func CreatePhoneNumber(numbers [10]uint) string {

	var res string = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ""), "[]")

	return "(" + res[:3] + ") " + res[3:6] + "-" + res[6:]
}

/*
func DigitalRoot(n int) int {

	if n <= 9 {
		return n
	}
	weAreDone := true
	res := 0

	for weAreDone {
		str := fmt.Sprintf("%v", n)
		for i := 0; i < len(str); i++ {
			num, err := strconv.Atoi(string(str[i]))
			if err == nil {
				res = res + num
			}
		}
		if res <= 9 {
			return res
		} else {
			n = res
			res = 0
		}
	}
	return 0
}

/func SpinWords(str string) string {

	strSlc := strings.Split(str, " ")
	res := ""
	var revertWord []byte
	for _, x := range strSlc {
		if len(x) >= 5 {
			for i := len(x) - 1; i >= 0; i-- {
				revertWord = append(revertWord, x[i])
			}
			res += string(revertWord) + " "
			revertWord = nil
		} else {
			res += x + " "
		}
	}
	return res[:len(res)-1]
} // SpinWords


func FindOdd(seq []int) int {
	dict := make(map[int]int)
	for _, num := range seq {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
	for k, v := range dict {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}


func Multiple3And5(number int) int {
	var sum int = 0

	if number < 3 {
		return 0
	}

	for i := 3; i < number; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	return sum

}
*/
