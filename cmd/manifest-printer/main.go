package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFromInput() []byte {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	var out []string

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {

			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		out = append(out, fmt.Sprintf("%s", string(buf)))

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	js := strings.Join(out, " ")

	return []byte(js)
}

func main() {

	s := ReadFromInput()
	fmt.Println(string(s))

}
