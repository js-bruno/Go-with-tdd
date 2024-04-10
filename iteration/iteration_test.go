package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	repetions := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repetions != expected {
		t.Errorf("Expected '%s', result '%s'", expected, repetions)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}

}

func ExampleRepeat() {
	repetions := Repeat("a", 10)
	fmt.Println(repetions)
	//Output: aaaaaaaaaa
}

