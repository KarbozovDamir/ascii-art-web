package dirs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	fileLen = 855
)

func GetArt(text, font string) (string, error) {

	arr := []string{}

	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()
	// if !isValid(text) {
	// 	fmt.Println("None-valid characters")
	// }

	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}
	if len(arr) != fileLen {
		fmt.Println("File is corrupted")
		fmt.Println()
	}
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	ans := ""
	for _, v := range argsArr {
		for x := 0; x < 8; x++ {
			for _, el := range v {
				n := (el-32)*9 + 1
				ans += arr[int(n)+x] + "\n"
			}
		}
	}
	return ans, nil
}

func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}
