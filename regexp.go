package erro

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	/*
		Note for contributors/users : these regexp have been made by me, taking my own source code as example for finding the right one to use.
		I use gofmt for source code formatting, that means this will work on most cases.
		Unfortunately, I didn't lastWriteToVar against other code formatting tools, so it may require some evolution.
		Feel free to create an issue or send a PR.
	*/
	regexpParseStack        = regexp.MustCompile(`((?:(?:[a-zA-Z._-]+)[/])*(?:[*a-zA-Z0-9_]*\.)+[a-zA-Z0-9_]+)\(((?:(?:0x[0-9a-f]+)|(?:...)[,\s]*)+)*\)[\s]+([/:\-a-zA-Z0-9\._]+)[:]([0-9]+)[\s](?:\+0x([0-9a-f]+))*`)
	regexpHexNumber         = regexp.MustCompile(`0x[0-9a-f]+`)
	regexpFindVarDefinition = func(varName string) *regexp.Regexp {
		return regexp.MustCompile(fmt.Sprintf(`%s[\s\:]*={1}([\s]*[a-zA-Z0-9\._]+)`, varName))
	}
)

// findFailingLine finds line where <var> is defined, if Debug(<var>) is present on lines[debugLine]. funcLine serves as max
func findFailingLine(lines []string, funcLine int, debugLine int) (failingLineIndex, columnStart, columnEnd int) {
	failingLineIndex = -1 //init error flag
	//find var name
	varName := MatchVarName(lines[debugLine-1])

	//build regexp for finding var definition
	reFindVar := regexpFindVarDefinition(*varName)

	//start to search for var definition
	for i := debugLine; i >= funcLine && i > 0; i-- { // going reverse from debug line to funcLine
		// early skipping some cases
		if strings.Trim(lines[i], " \n\t") == "" { // skip if line is blank
			//	(*LogTo).Debug().Msgf(color.BlueString("%d: ignoring blank line", i))
			continue
		} else if len(lines[i]) >= 2 && lines[i][:2] == "//" { // skip if line is a comment line (note: comments of type '/*' can be stopped inline and code may be placed after it, therefore we should pass line if '/*' starts the line)
			continue
		}

		//search for var definition
		index := reFindVar.FindStringSubmatchIndex(lines[i])
		if index == nil {
			continue
		}
		// At that point we found our definition

		failingLineIndex = i   //store the ressult
		columnStart = index[0] //store columnStart

		//now lets walk to columnEnd (because regexp is really bad at doing this)
		//for this purpose, we count brackets from first opening, and stop when openedBrackets == closedBrackets
		openedBrackets, closedBrackets := 0, 0
		for j := index[1]; j < len(lines[i]); j++ {
			if lines[i][j] == '(' {
				openedBrackets++
			} else if lines[i][j] == ')' {
				closedBrackets++
			}
			if openedBrackets == closedBrackets { // that means every opened brackets are now closed (the first/last one is the one from the func call)
				columnEnd = j // so we found our column end
				return        // so return the result
			}
		}

		if columnEnd == 0 { //columnEnd was not found
			if LogTo != nil {
				(*LogTo).Debug().Msgf("Fixing value of columnEnd (0). Defaulting to end of failing line.")
			}
			columnEnd = len(lines[i]) - 1
		}
		return
	}

	return
}
