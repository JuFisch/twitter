package main

//  "strings", "regexp"
//  "ioutil"

/*func main() {

	sourcefile, _ := os.Open("arendt.txt")
	newfile, err := os.Create("test2.txt")
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

	//flag for last line blank:
	lastLineBlank := false
	isFootnote := false

	for scanner.Scan() {

		line := scanner.Text()

		bytes := scanner.Bytes()

		// This section finds & prints footnotes as individual chunks.
		fnStart := regexp.MustCompile(`^\d+ [^\s]`)

		if fnStart.Match(bytes) && lastLineBlank == true {
			isFootnote = true
		}

		if len(bytes) == 0 {
			isFootnote = false
		}

		// need rule if -\n then delete it
		lastLineBlank = len(bytes) == 0

		//This section identifies footnotes embedded in text.
		fnText := regexp.MustCompile(`[^\s] \d{1,3} [^\s]`)

		if fnText.Match(bytes) == true {
			fmt.Println(string(bytes))
		}
		numline := regexp.MustCompile(`^\d{1,3}$`)
		ibid := regexp.MustCompile("(?i)Ibid.,")
		hyphen := regexp.MustCompile(`-$`)
		asterisk := regexp.MustCompile(`^\*`)

		// Conditionals for skipping lines we don't want
		// Skip empty lines
		if len(bytes) == 0 {
			continue
			// Skip lines less than four characters
		} else if len(line) <= 4 {
			continue
			// Skip lines with all caps
		} else if strings.ToUpper(line) == line {
			continue
			// Skip lines that are only digits
		} else if numline.Match(bytes) == true {
			continue
			// Skip ibid statements in footnotes
		} else if ibid.Match(bytes) == true {
			fmt.Println(line)
			continue
			// Skip footnotes
		} else if isFootnote == true {
			continue
			// Skip lines starting in asterisk
		} else if asterisk.MatchString(line) == true {
			continue
		} else {

			// Conditionals for cleaning text we do want

			if hyphen.MatchString(line) == true {
				line = strings.TrimSuffix(line, "-")
				bytes = []byte(line)
			}

			bits, err := newfile.Write(bytes)
			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Error printing", bits)
			}
			bits, err = newfile.WriteString("\n")
		}

	}
} */
