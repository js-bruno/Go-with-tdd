package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adds(2, 2)
  expected := 4
  if sum != expected {
    t.Errorf("Expected '%d', result '%d'", expected, sum)
  }
}

func ExampleAdds() {
	sum := Adds(1, 7)
  fmt.Println(sum)
  // Output: 8
}
