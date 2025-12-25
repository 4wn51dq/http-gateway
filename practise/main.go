package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

type block struct {
	version       uint32
	timestamp     uint64
	merkleRoot    []byte // bytes and uint8 are the same
	prevBlockHash []byte
	difficulty    uint64
	transactions  []tx
}

func (b block) getRoot() []byte {
	return b.merkleRoot
}

type tx struct{}

func main() {

	emptyBlock := block{}
	fmt.Println(emptyBlock)

	var emptyBlock2 block
	fmt.Println(emptyBlock2)

	var someBlock = struct {
		version       uint32
		timestamp     uint64
		merkleRoot    []byte // bytes and uint8 are the same
		prevBlockHash []byte
		difficulty    uint64
		transactions  []tx
	}{ /*respective values*/ }
	fmt.Println(someBlock)

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

	myString := []rune("héllo Mffffssssssș")
	var indexed = myString[0]
	fmt.Println("\n", indexed)
	fmt.Printf("value: %v, Type: %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	strSlice := []string{"h", "e", "l", "l", "ǒ"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var concatStrSlice = strBuilder.String()
	fmt.Println(concatStrSlice)

	var p *int32               // this holds a nil value in it, because it points to literally nothing
	var i int32                // the value held here is 0
	var p1 *int32 = new(int32) // the value here here is some memory address (not nil) that points to a value of 0 somewhere in memory
	fmt.Printf("the value in pointer p is nothing, and its memory address is %v\n", p)
	fmt.Printf("the value of i is %v, pointer p1 stores the memory address: %v, value in p1 is: %v\n", i, p1, *p1)

	/* *p = 67
	fmt.Printf("after resetting p, its memory address: %v, its value: %v", p, *p)
	*/
	// cannot derefence a nil pointer:
	// dereferencing means to read/ write data in a memory address, but there literally exists no address

	*p1 = 55
	fmt.Printf("after dereferencing p1, its new value: %v, its new memory address: %v is same as old\n", *p1, p1)

	//var p2 *int32 = new(int32)
	p2 := &i
	fmt.Printf("the memory address of i is %v\n", p2)
	*p2 = 67
	fmt.Println(i)

	someSl := []int32{1, 1, 2, 3}
	sliceCopy := someSl
	sliceCopy[2] = 4 // this will change both the slice and its copy
	fmt.Println(someSl, sliceCopy)
	// this happens because slices contain pointers to an underlying array
	// both variables refer to the same data in memory

	decArr := [5]float64{8, 3.14, 9.81, 2.71, 256}
	fmt.Printf("The memory location of decArr is %p\n", &decArr)
	arrSquared := square(decArr)
	fmt.Printf("squaring everything in decArr gives %v\n", arrSquared)
	// copy=> different address and pointers=> same address

	/* t00 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("time taken for execution since t00 is %v\n", time.Since(t00))
	fmt.Printf("The results are: %v\n", results)
	*/

	c := make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
		// time.Sleep(time.Second * 1)
	}

	chickenChannel := make(chan string)
	websites := []string{"site1.com", "site2.com", "site3.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

const MAX_CHICKEN_PRICE float32 = 5

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("found a deal on chicken on website %s", <-chickenChannel)
}

func process(c chan int) {
	defer close(c) // close right before function exits
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting function")
}

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"aaj mood nahi", "diet coke", "omega 3", "beer"}
var results = []string{}

func dbCall(i int) {
	//m.Lock()
	time.Sleep(time.Second * 2)
	// fmt.Printf("The result from database is %v\n", dbData[i])
	// m.Lock()

	wg.Done()
}

func count() {
	var count int
	for i := 0; i < 1000000; i++ {
		count += 1
	}
	wg.Done()
}

func square(decArr [5]float64) [5]float64 {
	fmt.Printf("the memory location of decArr is %p\n", &decArr)
	for i := range decArr {
		decArr[i] = decArr[i] * decArr[i]
	}
	return decArr
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
