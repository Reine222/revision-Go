package main
import ("fmt")

func add(a *int, b *int) int {
  var som = *a + *b
  return som
}

func main() {
  a := 5
  b := 6
  fmt.Println(add(&a, &b))
}