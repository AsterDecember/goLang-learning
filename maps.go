package main
import "fmt"
func main() {
	// make(map[type of the value], type of the key)
	StudentAge := make(map[string] int)

	StudentAge["Mike"]= 25
	StudentAge["Gerardo"] = 24
	StudentAge["Edu"] = 12
	StudentAge["Em"] = 29
	StudentAge["asdsadsda"] = 4

	fmt.Println(StudentAge)
	fmt.Println(StudentAge["Mike"])
	fmt.Println(len(StudentAge))

	//map inside a map
	superhero := map[string]map[string]string{

		"Superman" : map[string]string {
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman" : map[string]string{
			"RealName" : "Bruce Wayne",
			"City"	   : "Gotham City",
		},
	}

	//fmt.Println(superhero)

	//Print the value if they exist inside a map
	if temp, hero := superhero["Superman"]; hero{
		fmt.Println(temp["RealName"],temp["City"])
	}

	if tempVar, valueFromMap := superhero["Batman"]; valueFromMap{
		fmt.Println(tempVar["RealName"], tempVar["City"])
	}

}
