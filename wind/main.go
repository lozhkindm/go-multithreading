package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int
)

func parseToArray(textChan chan string, metarChan chan []string) {
	for text := range textChan {
		lines := strings.Split(text, "\n")
		metarSlice := make([]string, 0, len(lines))
		metarStr := ""
		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}
			if metarClose.MatchString(line) {
				metarSlice = append(metarSlice, metarStr)
				metarStr = ""
			}
		}
		metarChan <- metarSlice
	}
	close(metarChan)
}

func extractWindDirection(metarChan chan []string, windChan chan []string) {
	for metars := range metarChan {
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			if windRegex.MatchString(metar) {
				winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}
		windChan <- winds
	}
	close(windChan)
}

func mineWindDistribution(windChan chan []string, resultChan chan [8]int) {
	for winds := range windChan {
		for _, wind := range winds {
			if variableWind.MatchString(wind) {
				for i := 0; i < 8; i++ {
					windDist[i]++
				}
			} else if validWind.MatchString(wind) {
				windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
				if d, err := strconv.ParseFloat(windStr, 64); err == nil {
					dirIndex := int(math.Round(d/45.0)) % 8
					windDist[dirIndex]++
				}
			}
		}
	}
	resultChan <- windDist
	close(resultChan)
}

func main() {
	textChan := make(chan string)
	metarChan := make(chan []string)
	windChan := make(chan []string)
	resultChan := make(chan [8]int)

	//1. Change to array, each metar report is a separate item in the array
	go parseToArray(textChan, metarChan)

	//2. Extract wind direction, EGLL 312350Z 07004KT CAVOK 12/09 Q1016 NOSIG= -> 070
	go extractWindDirection(metarChan, windChan)

	//3. Assign to N, NE, E, SE, S, SW, W, NW, 070 -> E + 1
	go mineWindDistribution(windChan, resultChan)

	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()
	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(dat)
		textChan <- text
	}
	close(textChan)
	result := <-resultChan
	elapsed := time.Since(start)
	fmt.Printf("%v\n", result)
	fmt.Printf("Processing took %s\n", elapsed)
}
