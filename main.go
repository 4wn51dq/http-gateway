package main

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	var intNum uint32
	fmt.Println(intNum)

	var flt float32
	fmt.Println(flt)

	var said string = "Hello" + " " + "World"
	fmt.Println(said)

	fmt.Println(len("șl̥"))
	fmt.Println(utf8.RuneCountInString(said))

	var truthValue bool = false
	fmt.Println(truthValue)

	myVar := ""
	fmt.Println(myVar)

	var1, var2 := "this", "that"

	fmt.Println(var1, "and", var2)

	const someConst string = "literally some constant"
	fmt.Println(someConst)

	printMe("cake", 8)

	add(-2387423, 32423)
	quotient, remainder, err := divide(-139443, 0)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("the value of division is %v", quotient)
	default:
		fmt.Printf("the value of division is %v and the remainder is %v", quotient, remainder)
	}
	/*	if err != nil {
			fmt.Println("could not do division")
		} else if remainder == 0 {
			fmt.Printf("the result of divison is %v", quotient)
		} else {
			fmt.Printf("The result of division is %v and remainder is %v", quotient, remainder)
		}
	*/

	switch remainder {
	case 0:
		fmt.Println("the division was exact")
	case 1, 2:
		fmt.Println("the division was close")
	default:
		fmt.Println("the division was not close")
	}

	intArr := [3]int32{1, 2, 3}
	intArr1 := [...]int32{0, 0}
	var intArr2 [3]int32
	intArr2[0] = -2
	intArr2[1] = 33
	intArr3 := [...]int32{}
	// intArr4 := [...]int32{}
	// intArr4.push(xyz) cannot push to an array in go.. so we have slices
	fmt.Println(intArr, intArr1, intArr2, intArr3)

	var intSlice []int32 = []int32{1, 1, 2}
	intSlice = append(intSlice, 3, 5, 8, 13)
	intSlice1 := []int32{}
	intSlice1 = append(intSlice1, 21, 34, 55)
	intSlice2 := make([]int32, 0) // a slice of length 0 and capacity is also 0
	fmt.Println(intSlice)
	fmt.Printf("the length of the slice is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("the length of the slice1 is %v and the capacity is %v\n", len(intSlice1), cap(intSlice1))
	fmt.Println(intSlice2)

	intSlice = append(intSlice, intSlice1...)
	fmt.Println(intSlice)
	fmt.Printf("after appending second slice to first, the slice length is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	for i := 0; i < 9; i++ {
		intSlice2 = append(intSlice2, int32(i))
	}
	fmt.Println(intSlice2, len(intSlice2), cap(intSlice2))

	var ageMap map[string]uint16 = make(map[string]uint16)
	fmt.Println(ageMap)
	ageMap2 := map[string]uint16{"father": 55, "son": 20, "daughter": 27}
	fmt.Println(ageMap2["father"])

	delete(ageMap2, "daughter")

	var age, ok = ageMap2["daughter"]
	if ok {
		fmt.Printf("the age of the son i s %v", age)
	} else {
		fmt.Println("invalid name")
	}

	for name, age := range ageMap2 {
		fmt.Printf("[ Name: %v, Age: %v\n ]", name, age)
	}

	for i, v := range intSlice {
		fmt.Printf("\nIndex: %v, Value: %v", i, v)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice1 = make([]int, 0, n)
	fmt.Printf("\nTotal time without preallocation of capacity: %v", timeLoop(testSlice, n))
	fmt.Printf("\nTotal time with preallocation: %v", timeLoop(testSlice1, n))
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func printMe(printValue string, intValue int32) {
	fmt.Println(printValue, "with", intValue, "slices")
}

func add(a int64, b uint32) {
	fmt.Println(a + int64(b))
}

func divide(a int64, b uint32) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	result := int(int32(a) / int32(b))
	remainder := int(a % int64(b))
	return result, remainder, err
}
