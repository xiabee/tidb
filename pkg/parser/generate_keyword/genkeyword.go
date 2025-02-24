// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	reservedKeywordStart   = "The following tokens belong to ReservedKeyword"
	unreservedkeywordStart = "The following tokens belong to UnReservedKeyword"
	notKeywordStart        = "The following tokens belong to NotKeywordToken"
	tiDBKeywordStart       = "The following tokens belong to TiDBKeyword"
)

const (
	sectionNone = iota
	sectionReservedKeyword
	sectionUnreservedKeyword
	sectionTiDBKeyword
)

const (
	fileStart = `// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

// WARNING: This file is generated by 'genkeyword'

// KeywordsType defines the attributes of keywords
type KeywordsType struct {
	Word     string
	Reserved bool
	Section  string
}

// Keywords is used for all keywords in TiDB
var Keywords = []KeywordsType{
`
	fileEnd = `}
`
)

var keywordRe *regexp.Regexp

// parseLine extracts a keyword from a line of parser.y
// returns an empty string if there is no match.
//
// example data:
//
//	add               "ADD"
//
// Note that all keywords except `TiDB_CURRENT_TSO` are fully uppercase.
func parseLine(line string) string {
	if keywordRe == nil {
		keywordRe = regexp.MustCompile(`^\s+\w+\s+"(\w+)"$`)
	}
	m := keywordRe.FindStringSubmatch(line)
	if len(m) != 2 {
		return ""
	}
	return m[1]
}

func main() {
	parserData, err := os.ReadFile("parser.y")
	if err != nil {
		log.Fatalf("Failed to open parser.y: %s", err)
	}
	keywordsFile, err := os.OpenFile("keywords.go", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed to create keywords.go: %s", err)
	}
	_, err = keywordsFile.WriteString(fileStart)
	if err != nil {
		log.Fatalf("Failed to write fileStart to keywords.go: %s", err)
	}

	section := sectionNone
	for _, line := range strings.Split(string(parserData), "\n") {
		if line == "" { // Empty line indicates section end
			section = sectionNone
		} else if strings.Contains(line, reservedKeywordStart) {
			section = sectionReservedKeyword
		} else if strings.Contains(line, unreservedkeywordStart) {
			section = sectionUnreservedKeyword
		} else if strings.Contains(line, tiDBKeywordStart) {
			section = sectionTiDBKeyword
		} else if strings.Contains(line, notKeywordStart) {
			section = sectionNone
		}

		switch section {
		case sectionReservedKeyword:
			word := parseLine(line)
			if len(word) > 0 {
				fmt.Fprintf(keywordsFile, "\t{\"%s\", true, \"reserved\"},\n", word)
			}
		case sectionTiDBKeyword:
			word := parseLine(line)
			if len(word) > 0 {
				fmt.Fprintf(keywordsFile, "\t{\"%s\", false, \"tidb\"},\n", word)
			}
		case sectionUnreservedKeyword:
			word := parseLine(line)
			if len(word) > 0 {
				fmt.Fprintf(keywordsFile, "\t{\"%s\", false, \"unreserved\"},\n", word)
			}
		}
	}

	_, err = keywordsFile.WriteString(fileEnd)
	if err != nil {
		log.Fatalf("Failed to write fileEnd to keywords.go: %s", err)
	}
}
