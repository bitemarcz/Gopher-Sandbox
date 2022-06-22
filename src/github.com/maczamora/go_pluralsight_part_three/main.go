package main

import "fmt"

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Arrays -*-*-*-*-*-*-*-*-*-*-*")
	var arr [3]int // long form to create an array and may be verbose

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)    // printing array
	fmt.Println(arr[1]) // printing out a single value in the array

	arr2 := [3]int{4, 5, 6} // less typing does not remove clarity !!! Cannot change the array it's already set and immutable
	fmt.Println(arr2)       // printing array

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Slices  -*-*-*-*-*-*-*-*-*-*-*")
	mySlice := arr[:] // colon operators creates a slice of the array collection type -- start to finish of the array
	fmt.Println(arr, mySlice)

	anotherSlice := []int{7, 8, 9} // slice not fixed size or is mutable
	fmt.Println(anotherSlice)
	anotherSlice = append(anotherSlice, 10, 11, 12)
	fmt.Println(anotherSlice)

	s2 := anotherSlice[1:]
	s3 := anotherSlice[:2]
	s4 := anotherSlice[1:2]

	fmt.Println(s2, s3, s4)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Maps -*-*-*-*-*-*-*-*-*-*-*")
	myMap := map[string]int{"foo": 23} // key value relationships
	fmt.Println(myMap)
	fmt.Println(myMap["foo"])

	myMap["foo"] = 21
	fmt.Println(myMap)
	delete(myMap, "foo")
	fmt.Println(myMap)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Structs -*-*-*-*-*-*-*-*-*-*-*") // only collection that allows us to use disparate data types
	type User struct {                                                   // in scope of main function and can be used in package level or anywhere that makes sense
		ID        int
		FirstName string
		LastName  string
	}

	var u User
	u.ID = 1
	u.FirstName = "Mario"
	u.LastName = "Zamora"
	fmt.Println(u)

	u2 := User{ID: 1,
		FirstName: "Elizabeth",
		LastName:  "Dent",
	}
	fmt.Println(u2)
}
