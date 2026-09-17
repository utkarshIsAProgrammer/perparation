package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"Henry":   31,
		"Charlie": 34}
	fmt.Println(ages)
	fmt.Println(ages["Henry"], len(ages))

	// empty map
	temp := make(map[string]float64)
	fmt.Println(temp)

	temp["Monday"] = 35
	temp["Tuesday"] = 32
	temp["Wednesday"] = 44
	temp["Thursday"] = 29
	temp["Friday"] = 22
	fmt.Println(temp)

	fmt.Println(temp["Friday"])

	// update value
	temp["Friday"] = 25
	fmt.Println(temp["Friday"])

	// delete key & value
	delete(temp, "Tuesday")
	fmt.Println(temp)

	// check value exists
	value, exists := temp["Tuesday"]
	fmt.Println(value, exists)

	// check value exists (preferred due to short var name)
	res, ok := temp["Wednesday"]
	fmt.Println(res, ok)

}
