//   Copyright 2021 Lars Lehtonen
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//   http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// https://superuser.com/questions/380772/removing-ansi-color-codes-from-text-stream
	// https://stackoverflow.com/questions/14693701/how-can-i-remove-the-ansi-escape-sequences-from-a-string-in-python
	rxp := regexp.MustCompile(`\x1B\[[0-Z;]*[a-zA-Z]`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		input := scanner.Bytes()
		os.Stdout.Write(rxp.ReplaceAll(input, []byte{}))
		fmt.Print("\n")
	}
}
