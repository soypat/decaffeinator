package main

import "fmt"

func main() {
	ages := map[string]int{
		"Sarah":        32,
		"Billy":        12,
		"Jeremiah":     99,
		"John Baptist": 47,
	}
	fmt.Println(ages["Sarah"])

	billyAge, billyPresent := ages["Billy"]
	fmt.Println(billyAge, billyPresent)

	x13Age, x13Present := ages["x13"]
	fmt.Println(x13Age, x13Present)

	ages["Faustus"] = 66
	fmt.Println(ages)
}
