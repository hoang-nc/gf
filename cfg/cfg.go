package cfg

import (
	"bufio"
	"github.com/goframework/gf/ext"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cfg struct {
	Data map[string]string
}

func (this *Cfg) Str(key string, defaultValue string) string {
	v, ok := this.Data[key]
	if ok {
		return v
	}
	return defaultValue
}

func (this *Cfg) Int(key string, defaultValue int) int {
	v, ok := this.Data[key]
	if ok {
		num, err := strconv.Atoi(v)
		if err != nil {
			return num
		}
	}
	return defaultValue

}

func (this *Cfg) Load(file string) {

	// Open an input file, exit on error.
	inputFile, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}

	// Closes the file when we leave the scope of the current function,
	// this makes sure we never forget to close the file if the
	// function can exit in multiple places.
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	this.Data = map[string]string{}
	// scanner.Scan() advances to the next token returning false if an error was encountered
	for scanner.Scan() {
		line := scanner.Text()

		eqIndex := strings.Index(line, "=")
		if eqIndex > 0 {

			key := strings.Trim(line[:eqIndex], " 	　")
			if key[0] != '#' {
				value := strings.Trim(line[eqIndex+1:], " 　")
				value = ext.ReplaceEnv(value)
				this.Data[key] = value
			}
		}
	}

	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}
