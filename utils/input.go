package utils

import (
	"bufio"
	"io"
)

//ReadPipe reads data from a pipe
func ReadPipe(pipe io.ReadCloser) []byte {
	output := []byte{}
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			output = append(output, b)
		}
		output = append(output, byte('\n'))
	}
	return output
}
