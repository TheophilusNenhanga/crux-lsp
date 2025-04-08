package parser

import (
	"testing"
)

const source = "use _time_ms from \"crux:time\";\n\nlet max = 440000000;\n\nlet cache = [0];\n\nfor let i = 1; i < 10; i += 1 {\n    cache.push(i**i);\n}\n\nfn is_munchausen(number) {\n    let n = number;\n    let total = 0;\n\n    while n > 0 {\n        total += cache[n % 10];\n        if total > number {\n            return false;\n        }\n        n = n \\ 10;\n    }\n    return total == number;\n}\n\nfn main() {\n    let time_start = _time_ms();\n    for let i = 0; i < max; i += 1 {\n        if is_munchausen(i) {\n            println(i);\n        }\n    }\n    let time_end = _time_ms();\n    let time_diff = time_end - time_start;\n    println(\"Time taken: \" + time_diff + \" seconds\");\n}\n\nmain();\n"

func TestNewScanner(t *testing.T) {
	scanner := NewScanner(source)
	if scanner == nil {
		t.Fatal("NewScanner failed")
	}
	if scanner.source != source {
		t.Fatal("NewScanner failed")
	}
}

func TestScanner_Advance(t *testing.T) {
	scanner := NewScanner(source)
	if scanner.current != 0 {
		t.Fatal("NewScanner failed")
	}
	scanner.Advance()
	if scanner.current != 1 {
		t.Fatal("Advance failed")
	}
}
