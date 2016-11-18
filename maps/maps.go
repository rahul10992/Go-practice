package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68, -74.399}
	fmt.Println(m["Bell Labs"]) // we have a nil concept here as well
	//Map literals below
	var map2 = map[string]Vertex{
		"Bell Labs": Vertex{40.68, -74.399},
		"Google":    Vertex{37.42, -122.08},
	}
	fmt.Println("Map2: ", map2)

	//Mutating maps
	map3 := make(map[string]int)
	map3["Answer"] = 3                        // inserting or updating: m[key] = value
	fmt.Println("The value:", map3["Answer"]) //retreiving a value elem = m[key]

	map3["Answer"] = 23
	fmt.Println("The value:", map3["Answer"])

	delete(map3, "Answer") // just guess what this is for
	fmt.Println("The value:", map3["Answer"])
	//to test that key is present with 2-value assignment elem, ok = m[key]
	v, ok1 := map3["Answer"] // alternative syntax: elem, ok = m[key] only if they have been declared before
	fmt.Println("The value:", v, "Present?", ok1)
}
