func main() {

  // This is a slice of integers
  var numbers []int

  // Appending numbers to the slice
  numbers = append(numbers, 1, 2, 3, 4, 5)

  // Check for out of bounds conditions
  index := 5
  if index >= len(numbers) {
    fmt.Println("Index out of range!")
  } else {
    fmt.Println(numbers[index]) // This will not cause a runtime error
  }

} 
