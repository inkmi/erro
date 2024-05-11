package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/StephanSchmidt/erro/internal"

	"github.com/fatih/color"
)

func readFilesInDirectory(dir string) (map[string][]string, error) {
	files := make(map[string][]string)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filename := filepath.Base(path)
			files[filename] = append(files[filename], path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}

// LogEntry represents the structure of the log entry
type LogEntry struct {
	Level LogLevel        `json:"level"`
	Data  json.RawMessage `json:"data"`
}

type LogLevel struct {
	Level int             `json:"level"`
	Color color.Attribute `json:"color"`
}

func extractFilename(fullPath string) string {
	// Split the path by the "/" character
	parts := strings.Split(fullPath, "/")

	// The last element will be the filename with line number
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}

	// If there's no "/" or input is empty, return the original input
	return fullPath
}

func IntToLogLevel(level int) string {
	switch level {
	case 6:
		return "PNC"
	case 5:
		return "FAT"
	case 4:
		return "ERR"
	case 3:
		return "WRN"
	case 2:
		return "INF"
	case 1:
		return "DGG"
	case 0:
		return "TRC"
	default:
		return "???" // Return "unknown" for invalid integer values
	}
}

func LogLevelToInt(logLevel string) LogLevel {
	switch strings.ToLower(logLevel) {
	case "panic":
		return LogLevel{6, color.FgRed}
	case "fatal":
		return LogLevel{5, color.FgRed}
	case "error":
		return LogLevel{4, color.FgRed}
	case "warn":
		return LogLevel{3, color.FgYellow}
	case "info":
		return LogLevel{2, color.FgGreen}
	case "debug":
		return LogLevel{1, color.FgWhite}
	case "trace":
		return LogLevel{0, color.FgWhite}
	default:
		return LogLevel{-1, color.FgWhite} // Return -1 for unknown log levels
	}
}

func formatJsonMap(jsonData map[string]interface{}) string {

	// Create a cyan color for keys and "="
	cyan := color.New(color.FgCyan).SprintFunc()

	// Build the string with colors
	var b strings.Builder
	for key, value := range jsonData {
		// Add key and "=" in cyan
		b.WriteString(cyan(key))
		b.WriteString(cyan("="))
		b.WriteString(fmt.Sprint(value))
		b.WriteString(" ")
	}

	// Remove the trailing space and convert to a string
	result := strings.TrimSpace(b.String())
	return result
}

func main() {

	files, _ := readFilesInDirectory(".")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var jsonMap map[string]interface{}

		decoder := json.NewDecoder(strings.NewReader(line))
		decoder.UseNumber()
		err := decoder.Decode(&jsonMap)
		if err != nil {
			fmt.Fprintln(os.Stderr, line)
			continue
		}

		// Extract data preserving other fields as a sub-JSON
		level, ok := jsonMap["level"].(string)
		if !ok {
			level = "unknown"
		}

		delete(jsonMap, "level") // Remove the level so only data remains
		dataBytes, err := json.Marshal(jsonMap)
		if err != nil {
			fmt.Fprintln(os.Stderr, line)
			continue
		}

		entry := LogEntry{
			Level: LogLevelToInt(level),
			Data:  dataBytes,
		}

		message := ""
		m, ok := jsonMap["message"].(string)
		if ok {
			message = m
			delete(jsonMap, "message") // Remove the level so only data remains
		}

		caller := ""
		call, ok := jsonMap["caller"].(string)
		if ok {
			caller = call
			delete(jsonMap, "caller") // Remove the level so only data remains
		}

		errorStr := ""
		errorParse, ok := jsonMap["error"].(string)
		if ok {
			errorStr = errorStr
			delete(jsonMap, "error") // Remove the level so only data remains
		}

		ts := int64(0)
		tn, ok := jsonMap["time"].(json.Number)
		if ok {
			ts, err = tn.Int64()
			if err == nil {
				delete(jsonMap, "time") // Remove the level so only data remains
			}
		}

		t, ok := jsonMap["time"].(int64)
		if ok {
			ts = t
			delete(jsonMap, "time") // Remove the level so only data remains
		}

		tstr, ok := jsonMap["time"].(string)
		if ok {
			layout := "2006-01-02T15:04:05-07:00"
			parsedTime, err := time.Parse(layout, tstr)
			if err == nil {
				ts = parsedTime.Unix()
			}
			delete(jsonMap, "time") // Remove the level so only data remains
		}

		data := formatJsonMap(jsonMap)

		c := color.New(entry.Level.Color).SprintFunc()

		cyan := color.New(color.FgCyan).SprintFunc()
		black := color.New(color.FgHiBlack).SprintFunc()

		stamp := time.Unix(ts, 0).Format("3:04PM")

		// Now you can process each log entry
		fmt.Printf("%s %s %s %s %s %s\n", black(stamp), extractFilename(caller), c(IntToLogLevel(entry.Level.Level)), cyan(">"), message, data)

		if entry.Level.Level == 4 {
			red := color.New(color.FgHiRed).SprintFunc()
			parts := strings.Split(caller, ":")
			lineNumber, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}

			fileName := parts[0]
			if strings.Index(fileName, "/") < 0 {
				fileName = files[fileName][0]
			}

			fmt.Println(red("ERROR: " + errorParse))

			internal.PrintError(fileName, lineNumber)

		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input: %s\n", err)
	}
}
