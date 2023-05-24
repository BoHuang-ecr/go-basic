/*
	A Map is an unordered collection of key-value pairs.
	The most important point of Map is to quickly retrieve data through the key, which is similar to an index and points to the value of the data.

	--> map_variable := make(map[KeyType]ValueType, initialCapacity)
		// create a empty Map
		m := make(map[string]int)

		// create a Map with capability 10
		m := make(map[string]int, 10)
*/

package main

import "fmt"

func main() {
	var siteMap map[string]string
	siteMap = make(map[string]string)

	siteMap["Germany"] = "Munich"
	siteMap["China"] = "Beijing"
	siteMap["UK"] = "London"
	siteMap["Italy"] = "Rom"

	for site := range siteMap {
		fmt.Println(site, "capital city", siteMap[site])
	}

	// check if exits in map
	name, ok := siteMap["Finnland"]
	if ok {
		fmt.Println("Finnland is", name)
	} else {
		fmt.Println("Finnland not exits")
	}

	// delete
	delete(siteMap, "Germany")

	// check if exits in map
	name, ok = siteMap["Germany"]
	if ok {
		fmt.Println("Germany is", name)
	} else {
		fmt.Println("Germany not exits")
	}
}
