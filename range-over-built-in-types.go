package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _,num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	
	// range on arrays and slices provides 
	// both the index and value for each entry. 
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points. 
	// The first value is the starting byte index 
	// of the rune and the second the rune itself. 
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}