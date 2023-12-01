package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func Part1(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	sum := 0

	for scanner.Scan() {
		buf := bytes.NewBuffer(scanner.Bytes())

		var (
			left  *rune
			right *rune
		)

		for {
			r, _, err := buf.ReadRune()
			if err != nil {
				// EOF (or end of line - wrap this up and add to the sum)
				if errors.Is(err, io.EOF) {
					out := &strings.Builder{}
					if left != nil {
						out.WriteRune(*left)
					}
					if right != nil {
						out.WriteRune(*right)
					} else {
						out.WriteRune(*left)
					}

					var value int
					value, err = strconv.Atoi(out.String())
					if err != nil {
						return 0, err
					}

					sum += value

					break
				}
				return 0, err
			}

			// If the rune value is a digit, determine if leftmost or rightmost
			if isDigit(r) {
				if left == nil {
					left = &r
					continue
				}
				right = &r
			}
		}
	}

	return sum, nil
}

func main() {
	f, err := os.Open("cmd/day-01/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sum, err := Part1(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sum)
}
