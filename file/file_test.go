package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestIterAllLines(t *testing.T) {

	want_str := "First Line\nSecond Line\nThird Line\n"

	input := strings.NewReader(want_str)

	scanner := bufio.NewScanner(input)

	var output bytes.Buffer

	iter(&output, scanner, 0, true)

	got_str := output.String()

	fmt.Println(got_str)

	if got_str != want_str {
		t.Errorf("Wanted %s, got %s", want_str, got_str)
	}

}

func TestIterOneLine(t *testing.T) {

	full_str := "First Line\nSecond Line\nThird Line\n"

	want_str := "First Line\n"

	input := strings.NewReader(full_str)

	scanner := bufio.NewScanner(input)

	var output bytes.Buffer

	iter(&output, scanner, 1, false)

	got_str := output.String()

	fmt.Println(got_str)

	if got_str != want_str {
		t.Errorf("Wanted %s, got %s", want_str, got_str)
	}

}

func TestIterAllLinesEqualsMaxNumLines(t *testing.T) {
	// checks if setting AllLines == true results in the same
	// output as setting NumLines to length of file

	want_str := "First Line\nSecond Line\nThird Line"

	var buffer io.Reader

	var scanner *bufio.Scanner

	buffer = strings.NewReader(want_str)

	scanner = bufio.NewScanner(buffer)

	if err := scanner.Err(); err != nil {
		t.Errorf("Scanner error: %v", err)
	}

	var buf1 bytes.Buffer

	var buf2 bytes.Buffer

	iter(&buf1, scanner, 3, false)

	buf1_str := buf1.String()

	buffer = strings.NewReader(want_str)

	scanner = bufio.NewScanner(buffer)

	iter(&buf2, scanner, 0, true)

	buf2_str := buf2.String()

	if buf2_str != buf1_str {
		t.Errorf("Result from allLines and numLines not equal: allLines: %s, numLines %s", buf2_str, buf1_str)
	}

}
