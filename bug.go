func main() {

  // This is a slice of integers
  var numbers []int

  // Appending numbers to the slice
  numbers = append(numbers, 1, 2, 3, 4, 5)

  // Accessing the slice out of bounds leads to a runtime panic
  fmt.Println(numbers[5]) // This will cause a runtime error

}