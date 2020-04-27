package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "Dwyane",
		"course": "golang",
		"age":    "eighteen",
		"sex":    "boy",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key dose not exist")
	}

	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
