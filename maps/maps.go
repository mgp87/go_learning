package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string) // make(map[keyType]valueType, size)
	fmt.Println(countries)

	countries["Mexico"] = "D.F."
	countries["Argentina"] = "Buenos Aires"
	countries["España"] = "Madrid"
	fmt.Println(countries)
	fmt.Println(countries["España"])

	championship := map[string]int{
		"Barcelona":    25,
		"Real Madrid":  32,
		"Chivas":       11,
		"Boca Juniors": 18,
	}
	fmt.Println(championship)
	for team, points := range championship {
		fmt.Printf("The team %s has %d points \n", team, points)
	}
	delete(championship, "Barcelona") // If key not found, nothing happens
	fmt.Println(championship)

	points, exists := championship["Chivas"]
	fmt.Printf("Points: %d, Exists: %t \n", points, exists)

	points, exists = championship["Juventus"]
	fmt.Printf("Points: %d, Exists: %t \n", points, exists) // 0, false since the element doesn't exist in the map
}
