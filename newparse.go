package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//  "strings", "regexp"
//  "ioutil"

func tweetSplit(s string) []string {

	line := s
	tweetables := make([]string, 0)

	tweetSeparator := " [...]"
	charLimit := 140 - len(tweetSeparator)

	for len(line) > charLimit {

		firstChunk := line[0:charLimit]
		lastspace := strings.LastIndex(firstChunk, " ")
		tweetChunk := line[0:lastspace]
		startTrimmed := lastspace + 1
		trimmed := firstChunk[startTrimmed:charLimit]
		tweet := tweetChunk + tweetSeparator
		tweetables = append(tweetables, tweet)

		line = strings.TrimSuffix(trimmed, " ") + line[charLimit:len(line)]

	}

	tweetables = append(tweetables, line)
	return tweetables

}

func main() {

	// Create a list of footnotes from footnotes.txt
	// This will help us later on with cleaning & appending footnotes.

	// Open footnotes file & defer close.
	fnSourcefile, err := os.Open("arendtfootnotes.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(files []*os.File) {
		for _, file := range files {
			file.Close()
		}
	}([]*os.File{fnSourcefile})

	// Create a line-by-line scanner for footnotes file.
	fnScanner := bufio.NewScanner(fnSourcefile)
	fnScanner.Split(bufio.ScanLines)

	// Set flags for footnote parsing.
	firstNote := true
	emptylinecount := 0
	asterisks := false

	// Set footnote regex rule.
	digit := regexp.MustCompile(`(^\d+a?)`)

	// Initialize empty footnote storage.
	var footnoteStore []int

	// Initalize empty list of footnotes that met the a-suffix rule (e.g. 62a)
	var aSuffixStore []int

	// Scan file for footnotes.
	for fnScanner.Scan() {

		line := fnScanner.Text()

		if firstNote == true {
			if digit.FindString(line) != "" {
				fnNum, err := strconv.Atoi(digit.FindString(line))
				if err != nil {
					fmt.Println("Footnote1 err: ", err)
				}
				footnoteStore = append(footnoteStore, fnNum)
				firstNote = false
			}
			continue
		}

		if line == "* * *" {
			asterisks = true
			continue
		}

		if asterisks && len(line) == 0 {
			emptylinecount++
			continue
		}

		if emptylinecount == 3 && asterisks {
			emptylinecount = 0
			asterisks = false

			// Create footnoteStore by adding in all instances of footnotes found
			if digit.FindString(line) != "" {
				fnNum, err := strconv.Atoi(digit.FindString(line))
				//fmt.Println("Now we print the next footnote...", fnNum)
				if err == nil {
					footnoteStore = append(footnoteStore, fnNum)
				}

				if err != nil {
					if strings.HasSuffix(digit.FindString(line), "a") {
						//Trim the suffix 'a' off the string
						fnValue := strings.TrimSuffix(digit.FindString(line), "a")
						//Assign the resulting value to num as an int
						fnNumA, _ := strconv.Atoi(fnValue)
						footnoteStore = append(footnoteStore, fnNumA)
						aSuffixStore = append(aSuffixStore, fnNumA)
						//Indicate that this is a suffixed instance of a footnote
						//isSuffix = true
					}
				}
			}

		}
	}

	// Set last value flag for checking footnotes fit expected pattern
	lastValue := 0

	// Make the list of items that break our rules (we're expecting only a-suffixed notes here)
	var aSuffixExpected []int

	// Check that all footnotes meet expected patern (==1, ==lastValue+1, panic for 4 a-suffix footnotes)
	for _, num := range footnoteStore {
		if num == 1 {
			lastValue = num
			continue
		}

		if num == lastValue+1 {
			lastValue = num
			continue
		}

		aSuffixExpected = append(aSuffixExpected, num)
	}

	for i, item := range aSuffixStore {
		if item != aSuffixExpected[i] {
			fmt.Println("Error! Unexpected item in list of asuffix notes")
		}
	}

	// Do the rest of the parsing work
	sourcefile, _ := os.Open("arendtclean.txt")
	newfile, err := os.Create("testoutput.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(files []*os.File) {
		for _, file := range files {
			file.Close()
		}
	}([]*os.File{sourcefile, newfile})

	scanner := bufio.NewScanner(sourcefile)

	scanner.Split(bufio.ScanLines)

	count := 0

	// Regex rules to catch footnotes.
	footnote := regexp.MustCompile(`([\.!,:;"”\)]|[a-z]+)([1-9][0-9]*a?)`)

	// Set flag for footnoteNum (previous in list check) to 0.
	//footnoteNum := 0

	// Empty list that will be used to check that we addressed all expected footnotes from above list.
	var fnCheck []int

	var fnStoreConvert []string
	for _, item := range footnoteStore {
		str := strconv.Itoa(item)
		fnStoreConvert = append(fnStoreConvert, str)
	}

	var fnIndex = 0

	for scanner.Scan() {

		line := scanner.Text()

		if footnote.MatchString(line) == true {

			matches := footnote.FindAllStringSubmatch(line, -1)

			// For every match in range matches for this line, including a-suffixed notes:
			for _, match := range matches {

				// Assign num to the int value of match
				num, err := strconv.Atoi(match[2])
				// If you get an error, check if it's because it's a-suffixed and assign num to int of unsuffixed note
				if err != nil {
					if strings.HasSuffix(match[2], "a") {
						//Trim the suffix 'a' off the string
						value := strings.TrimSuffix(match[2], "a")
						//Assign the resulting value to num as an int
						num, err = strconv.Atoi(value)
						if err != nil {
							fmt.Println("Error", err)
						}
					}
				}

				if footnoteStore[fnIndex] == num {
					fnCheck = append(fnCheck, num)
					fnIndex++
				}

				/*if footnoteStore[fnIndex] == 12 {
					//fmt.Println(line)
				}
				if num == footnoteStore[fnIndex] {
					fnIndex++
					fnCheck = append(fnCheck, num)
				}*/

			}

		}

		tweets := make([]string, 0)

		if len(line) == 0 {
			continue
		} else if line == "[back]" {
			continue
		} else if line == "* * *" {
			continue
		} else if len(line) < 140 {
			tweets = append(tweets, line)
			count = count + 1
		} else {
			splitLine := tweetSplit(line)

			for _, item := range splitLine {
				tweets = append(tweets, item)
				count = count + 1
			}
		}

		for _, tweet := range tweets {
			byteTweet := []byte(tweet)
			bits, err := newfile.Write(byteTweet)
			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Error printing", bits)
			}
			bits, err = newfile.WriteString("\n")
		}

	}
	fmt.Println(len(fnCheck))
	fmt.Println(len(footnoteStore))

	for i, v := range footnoteStore {
		if v != fnCheck[i] {
			fmt.Printf("Break in iteration %d, value %d", i, v)
		}
	}
}
