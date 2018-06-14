package main

import (
    "bufio"
    "fmt"
    "os"
    "crypto/md5"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"encoding/hex"
	"time"

)


const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals = "0123456789"
	Ascii    = Alphabet + Numerals //+ "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)
 
type GeneratorExprRanges [][]byte
 
func seedAndReturnRandom(n int) int {
	return rand.Intn(n)
}
 
func alphabetSlice(from, to byte) (string, error) {
	leftPos := strings.Index(Ascii, string(from))
	rightPos := strings.LastIndex(Ascii, string(to))
	if leftPos > rightPos {
		return "", fmt.Errorf("Invalid range specified: %s-%s", string(from), string(to))
	}
	return Ascii[leftPos:rightPos], nil
}
 
func replaceWithGenerated(s *string, expresion string, ranges [][]byte, length int) error {
	var alphabet string
	for _, r := range ranges {
		switch string(r[0]) + string(r[1]) {
		case `\w`:
			alphabet += Ascii
		case `\d`:
			alphabet += Numerals
		default:
			if slice, err := alphabetSlice(r[0], r[1]); err != nil {
				return err
			} else {
				alphabet += slice
			}
		}
	}
	if len(alphabet) == 0 {
		return fmt.Errorf("Empty range in expresion: %s", expresion)
	}
	result := make([]byte, length, length)
	for i := 0; i <= length-1; i++ {
		result[i] = alphabet[seedAndReturnRandom(len(alphabet))]
	}
	*s = strings.Replace(*s, expresion, string(result), 1)
	return nil
}
 
func findExpresionPos(s string) GeneratorExprRanges {
	rangeExp, _ := regexp.Compile(`([\\]?[a-zA-Z0-9]\-?[a-zA-Z0-9]?)`)
	matches := rangeExp.FindAllStringIndex(s, -1)
	result := make(GeneratorExprRanges, len(matches), len(matches))
	for i, r := range matches {
		result[i] = []byte{s[r[0]], s[r[1]-1]}
	}
	return result
}
 
func rangesAndLength(s string) (string, int, error) {
	expr := s[0:strings.LastIndex(s, "{")]
	length, err := parseLength(s)
	return expr, length, err
}
 
func parseLength(s string) (int, error) {
	lengthStr := string(s[strings.LastIndex(s, "{")+1 : len(s)-1])
	if l, err := strconv.Atoi(lengthStr); err != nil {
		return 0, fmt.Errorf("Unable to parse length from %v", s)
	} else {
		return l, nil
	}
}
 
func Generate(template string) (string, error) {
	result := template
	generatorsExp, _ := regexp.Compile(`\[([a-zA-Z0-9\-\\]+)\](\{([0-9]+)\})`)
	matches := generatorsExp.FindAllStringIndex(template, -1)
	for _, r := range matches {
		ranges, length, err := rangesAndLength(template[r[0]:r[1]])
		if err != nil {
			return "", err
		}
		positions := findExpresionPos(ranges)
		if err := replaceWithGenerated(&result, template[r[0]:r[1]], positions, length); err != nil {
			return "", err
		}
	}
	return result, nil
}


func main(){
	fmt.Println("MD5 Reset Code Bruteforcer for AppLock\nLength of Reset code is of 8 Char and contains only alpha numerics.\n")
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter MD5: ")
    md5hash, _ := reader.ReadString('\n')
    md5hash = strings.ToLower(md5hash)
    fmt.Println("Cracking the MD5: ",md5hash)
  
    for{
    	result, _ := Generate(`[a-z0-9]{8}`)
        data := []byte(result)
    	x:=md5.Sum(data)
    	hash:=hex.EncodeToString(x[:])
    	fmt.Println(time.Now().UTC(), " Generated Code : ", result," MD5: ", hash)
		cmp:=hash == md5hash
		if cmp == true {
			fmt.Print("\n\n")
			fmt.Println(time.Now().UTC(), " Reset Code Cracked!")
			fmt.Println(time.Now().UTC(), " Reset Code: ", result, " MD5 Match: ", hash )
			break
		}
		}

    
	
}
