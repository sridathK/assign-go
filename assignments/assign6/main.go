package main

import "fmt"

func main() {

	simpleMapOperation()
	mapIteration()
	mapOfMap()
}

func simpleMapOperation() {
	studentData := make(map[string]float32)

	studentData["virat"] = 10.00
	studentData["rohit"] = 9.00
	studentData["rahul"] = 8.00

	for k, v := range studentData {
		fmt.Printf("student %s, grade is %f", k, v)
		fmt.Println()
	}
	delete(studentData, "rohit")

	for k, v := range studentData {
		fmt.Printf("student %s, grade is %f", k, v)
		fmt.Println()
	}
}

func mapIteration() {
	fruitsMap := make(map[string]int)
	fruitsMap["mango"] = 1
	fruitsMap["grape"] = 2
	fruitsMap["banana"] = 3
	fruitsMap["pome"] = 4
	for k, v := range fruitsMap {
		fmt.Printf("fruit %s, quant is %d", k, v)
		fmt.Println()
	}

}

func mapOfMap() {
	studentData := make(map[string]map[string]string)
	studentDetails := make(map[string]string)
	studentDetails["age"] = "20"
	studentDetails["grade"] = "5"
	studentDetails["city"] = "hyd"
	studentData["ramu"] = studentDetails

	studentDetails["age"] = "25"
	studentDetails["grade"] = "15"
	studentDetails["city"] = "guj"

	studentData["raju"] = studentDetails

	studentDetails["age"] = "32"
	studentDetails["grade"] = "8"
	studentDetails["city"] = "bang"

	studentData["ravi"] = studentDetails

	for k, v := range studentData {
		fmt.Printf("name %s, details are %s", k, v)
		fmt.Println()
	}

}
