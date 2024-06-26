package internal

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	/*
		Note for contributors/users : these regexp have been made by me, taking my own source code as example for finding the right one to use.
		I use gofmt for source code formatting, that means this will work on most cases.
		Unfortunately, I didn't GoLastWriteToVar against other code formatting tools, so it may require some evolution.
		Feel free to create an issue or send a PR.
	*/
	regexpParseStack        = regexp.MustCompile(`((?:(?:[a-zA-Z._-]+)[/])*(?:[*a-zA-Z0-9_]*\.)+[a-zA-Z0-9_]+)\(((?:(?:0x[0-9a-f]+)|(?:...)[,\s]*)+)*\)[\s]+([/:\-a-zA-Z0-9\._]+)[:]([0-9]+)[\s](?:\+0x([0-9a-f]+))*`)
	regexpHexNumber         = regexp.MustCompile(`0x[0-9a-f]+`)
	regexpFindVarDefinition = func(varName string) *regexp.Regexp {
		return regexp.MustCompile(fmt.Sprintf(`%s[\s\:]*={1}([\s]*[a-zA-Z0-9\._]+)`, varName))
	}
)

// findFailingLine finds line where <var> is defined, if Debug(<var>) is present on lines[debugLine]. funcLine serves as max
func findFailingLine(lines []string, funcLine int, debugLine int) (int, int, int) {
	failingLineIndex, columnStart, columnEnd := -1, -1, -1
	//find var name
	varName := MatchVarName(lines[debugLine-1])

	//build regexp for finding var definition
	reFindVar := regexpFindVarDefinition(*varName)

	//start to search for var definition
	for i := debugLine; i >= funcLine && i > 0; i-- { // going reverse from debug line to funcLine
		// early skipping some cases
		if strings.Trim(lines[i], " \n\t") == "" { // skip if line is blank
			continue
		} else if len(lines[i]) >= 2 && lines[i][:2] == "//" { // skip if line is a comment line (note: comments of type '/*' can be stopped inline and code may be placed after it, therefore we should pass line if '/*' starts the line)
			continue
		}
		failingLineIndex = i
		failingLine := lines[i]

		//search for var definition
		index := reFindVar.FindStringSubmatchIndex(failingLine)
		if index == nil {
			continue
		}
		// At that point we found our definition
		columnStart = index[0]
		columnEnd = findStartEnd(failingLine, columnStart)
		return failingLineIndex, columnStart, columnEnd
	}
	return failingLineIndex, columnStart, columnEnd
}

func findStartEnd(failingLine string, start int) int {
	//now lets walk to columnEnd (because regexp is really bad at doing this)
	//for this purpose, we count brackets from first opening, and stop when openedBrackets == closedBrackets
	openedBrackets, closedBrackets := 0, 0
	for j := start; j < len(failingLine); j++ {
		if failingLine[j] == '(' {
			openedBrackets++
		} else if failingLine[j] == ')' {
			closedBrackets++
		}
		if openedBrackets > 0 && openedBrackets == closedBrackets { // that means every opened brackets are now closed (the first/last one is the one from the func call)
			return j // so return the result
		}
	}
	return len(failingLine) - 1
}
