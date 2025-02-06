package task2

import (
	"bytes"
	"strings"
	"testing"
)

const concatTimes = 50000

func BenchmarkPlus(b *testing.B) {
	var str string
	for i := 0; i < concatTimes; i++ {
		str += "q"
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < concatTimes; i++ {
		buffer.WriteString("q")
	}
}

func BenchmarkBuilder(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < concatTimes; i++ {
		sb.WriteString("q")
	}
}
