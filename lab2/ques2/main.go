package main

import "fmt"

func main() {

	// Slice
	students := []string{"TIWARI", "OJHA", "RAJPUT", "DAGAR", "MUDGIL"}

	fmt.Println("Original Slice:", students)

	// Append - User Input
	var name string
	fmt.Print("Enter a student name to add: ")
	fmt.Scanln(&name)

	students = append(students, name)
	fmt.Println("After Append:", students)

	// Remove by index
	students = append(students[:1], students[2:]...)
	fmt.Println("After Removing index 1:", students)

	// Update
	students[1] = "BHARDWAJ"
	fmt.Println("After Updating index 1:", students)

	// Map
	marks := map[string]int{
		"Math":    85,
		"Science": 90,
		"English": 88,
	}

	fmt.Println("\nOriginal Map:", marks)

	// Insert
	marks["Computer"] = 95
	fmt.Println("After Inserting Computer:", marks)

	// Lookup
	fmt.Println("Marks in Math:", marks["Math"])

	// Delete
	delete(marks, "Science")
	fmt.Println("After Deleting Science:", marks)
}
