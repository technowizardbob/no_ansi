package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func removeCtrlKeyEscapeCodes(input string) string {
	// Define regular expressions for ANSI escape codes
	pattern1 := `\^\[\[([0-9]{1,2}(;[0-9]{1,2})*)?[mGK]`
	pattern2 := `\^\[\[([0-9;]+)m|\x1B\[\d+[A-Za-z]`
	pattern3 := `\^\[\[[0-9;]*[mGK]`
	pattern4 := `\^\[(?:[[][^A-Za-z]*[A-Za-z]|[=<>!](?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[()#;?]*[@-~])`
	pattern5 := `\^\[\[[0-9;]*[A-Za-z]`

	// Compile the regular expressions
	re1 := regexp.MustCompile(pattern1)
	re2 := regexp.MustCompile(pattern2)
	re3 := regexp.MustCompile(pattern3)
	re4 := regexp.MustCompile(pattern4)
	re5 := regexp.MustCompile(pattern5)

	// Remove ANSI escape codes step by step
	noColors := re1.ReplaceAllString(input, "")
	output := re2.ReplaceAllString(noColors, "")
	safeOutput := re3.ReplaceAllString(output, "")
	saferOutput := re4.ReplaceAllString(safeOutput, "")
	out := re5.ReplaceAllString(saferOutput, "")

	return out
}

func removeOctalEscapeCodes(input string) string {
	// Define regular expressions for ANSI escape codes
	pattern1 := `\033\[([0-9]{1,2}(;[0-9]{1,2})*)?[mGK]`
	pattern2 := `\033\[([0-9;]+)m|\x1B\[\d+[A-Za-z]`
	pattern3 := `\033\[[0-9;]*[mGK]`
	pattern4 := `\033(?:[[][^A-Za-z]*[A-Za-z]|[=<>!](?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[()#;?]*[@-~])`
	pattern5 := `\033\[[0-9;]*[A-Za-z]`

	// Compile the regular expressions
	re1 := regexp.MustCompile(pattern1)
	re2 := regexp.MustCompile(pattern2)
	re3 := regexp.MustCompile(pattern3)
	re4 := regexp.MustCompile(pattern4)
	re5 := regexp.MustCompile(pattern5)

	// Remove ANSI escape codes step by step
	noColors := re1.ReplaceAllString(input, "")
	output := re2.ReplaceAllString(noColors, "")
	safeOutput := re3.ReplaceAllString(output, "")
	saferOutput := re4.ReplaceAllString(safeOutput, "")
	out := re5.ReplaceAllString(saferOutput, "")

	return out
}

func removeUnicodeEscapeCodes(input string) string {
	// Define regular expressions for ANSI escape codes
	pattern1 := `\\\\u001b\[([0-9]{1,2}(;[0-9]{1,2})*)?[mGK]`
	pattern2 := `\\\\u001b\[([0-9;]+)m|\x1B\[\d+[A-Za-z]`
	pattern3 := `\\\\u001b\[[0-9;]*[mGK]`
	pattern4 := `\\\\u001b(?:[[][^A-Za-z]*[A-Za-z]|[=<>!](?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[()#;?]*[@-~])`
	pattern5 := `\\\\u001b\[[0-9;]*[A-Za-z]`

	// Compile the regular expressions
	re1 := regexp.MustCompile(pattern1)
	re2 := regexp.MustCompile(pattern2)
	re3 := regexp.MustCompile(pattern3)
	re4 := regexp.MustCompile(pattern4)
	re5 := regexp.MustCompile(pattern5)

	// Remove ANSI escape codes step by step
	noColors := re1.ReplaceAllString(input, "")
	output := re2.ReplaceAllString(noColors, "")
	safeOutput := re3.ReplaceAllString(output, "")
	saferOutput := re4.ReplaceAllString(safeOutput, "")
	out := re5.ReplaceAllString(saferOutput, "")

	return out
}

func removeHexadecimalANSIEscapeCodes(input string) string {
	// Define regular expressions for ANSI escape codes
	pattern1 := `\x1B\[([0-9]{1,2}(;[0-9]{1,2})*)?[mGK]`
	pattern2 := `\x1B\[([0-9;]+)m|\x1B\[\d+[A-Za-z]`
	pattern3 := `\x1b\[[0-9;]*[mGK]`
	pattern4 := `\x1b(?:[[][^A-Za-z]*[A-Za-z]|[=<>!](?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[()#;?]*[@-~])`
	pattern5 := `\x1b\[[0-9;]*[A-Za-z]`

	// Compile the regular expressions
	re1 := regexp.MustCompile(pattern1)
	re2 := regexp.MustCompile(pattern2)
	re3 := regexp.MustCompile(pattern3)
	re4 := regexp.MustCompile(pattern4)
	re5 := regexp.MustCompile(pattern5)

	// Remove ANSI escape codes step by step
	noColors := re1.ReplaceAllString(input, "")
	output := re2.ReplaceAllString(noColors, "")
	safeOutput := re3.ReplaceAllString(output, "")
	saferOutput := re4.ReplaceAllString(safeOutput, "")
	out := re5.ReplaceAllString(saferOutput, "")

	return out
}

func removeBell(input string) string {
	pattern := `\x07|\033\[|\a`
	re := regexp.MustCompile(pattern)
	out := re.ReplaceAllString(input, "")
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		
		cleanLine := removeCtrlKeyEscapeCodes(line)
		cleanLine2 := removeOctalEscapeCodes(cleanLine)
		cleanLine3 := removeUnicodeEscapeCodes(cleanLine2)
		cleanLine4 := removeHexadecimalANSIEscapeCodes(cleanLine3)
		cleanLine5 := removeBell(cleanLine4)
		fmt.Println(cleanLine5)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
	}
}

