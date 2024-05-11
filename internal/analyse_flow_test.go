package internal

import (
	"reflect"
	"testing"
)

func TestFindErrorOrigin(t *testing.T) {
	tests := []struct {
		name          string
		lines         []string
		logLine       int
		expected      []int
		expectedError string
	}{
		{name: "",
			lines: []string{
				"// some comment",
				"err := someBigFunction(2)",
				"logger.Error().Err(err).Str(\"Test\", \"Test\").Msg(\"Error example\")",
			},
			logLine:       3,
			expected:      []int{1, 1},
			expectedError: "",
		},
		{
			name: "Simple Declaration",
			lines: []string{
				"err := errors.New(\"failed\")",
				"if err != nil {",
				"logger.Error().Err(err).Msg(\"Error occurred\")",
			},
			logLine:       3,
			expected:      []int{0, 0},
			expectedError: "",
		},
		{
			name: "Multi-Variable Declaration",
			lines: []string{
				"ok, err := someCall()",
				"if err != nil {",
				"logger.Error().Err(err).Msg(\"Failure\")",
			},
			logLine:       3,
			expected:      []int{0, 0},
			expectedError: "",
		},
		{
			name: "Existing Variable Assignment",
			lines: []string{
				"var err error",
				"err = someCall()",
				"if err != nil {",
				"logger.Error().Err(err).Msg(\"Updated error\")",
			},
			logLine:       4,
			expected:      []int{1, 1},
			expectedError: "",
		},
		{
			name: "No err in Log Statement",
			lines: []string{
				"var err error",
				"err = someCall()",
				"logger.Info().Msg(\"Just logging info\")",
			},
			logLine:       3,
			expected:      nil,
			expectedError: "error variable not found in the log statement",
		},
		{
			name: "err Not Found Before Log",
			lines: []string{
				"if something {",
				"logger.Error().Err(err).Msg(\"err used but not declared\")",
			},
			logLine:       2,
			expected:      nil,
			expectedError: "no assignment for err found before the log statement",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GolangFindErrorOrigin(tt.lines, tt.logLine)
			if (err != nil) != (tt.expectedError != "") || (err != nil && err.Error() != tt.expectedError) {
				t.Errorf("GolangFindErrorOrigin() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GolangFindErrorOrigin() got = %v, expected %v", got, tt.expected)
			}
		})
	}
}
