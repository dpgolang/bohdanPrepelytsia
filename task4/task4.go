package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func validateFlags(path, pattern string) error {
	if path == "" {
		return errors.New("path for file must be inputed")
	}
	if pattern == "" {
		return errors.New("pattern must be inputed")
	}
	return nil
}
func readAndCount(path, pattern string) (int, error) {
	count := 0

	file, err := os.Open(path)
	if err != nil {
		return 0, errors.New("this file does not exist or it is impossible to open it")
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		count += strings.Count(scan.Text(), pattern)
	}
	if err := scan.Err(); err != nil {
		return 0, errors.New("one of the lines was too long!")
	}
	return count, nil
}
func readAndReplace(path, pattern, text string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("this file does not exist or it is impossible to open it")
	}

	textForReplace := strings.Replace(string(read), pattern, text, -1)
	err = ioutil.WriteFile(path, []byte(textForReplace), 0)
	if err != nil {
		return errors.New("it is impossible to write to this file")
	}
	return nil
}
func main() {
	path := flag.String("p", "", "input path to your file")
	pattern := flag.String("t", "", "input pattern you want to be searched in file")
	replaseText := flag.String("r", "", "input text you want to be replaced instead of pattern")
	flag.Parse()
	if *path == "" && *pattern == "" && *replaseText == "" {
		fmt.Println("HELP: this program helps you to count the number of occurrences the pattern in the text" +
			" or replace pattern in text for some another text.\nUse flag -p to input path to your file.\n" +
			"Use flag -t to input your pattern.\nUse flag -r to input text you want to be replaced instead of pattern")
		return
	}
	if err := validateFlags(*path, *pattern); err != nil {
		fmt.Println(err)
		return
	}
	if *replaseText != "" {
		if err := readAndReplace(*path, *pattern, *replaseText); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Done!")
		return
	} else {
		count, err := readAndCount(*path, *pattern)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Done! Count of occurrences:%d\n", count)
	}

}
