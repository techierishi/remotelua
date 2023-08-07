package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/zyedidia/highlight"
)

func colorSyntax(inputString string) {

	syntaxFile, _ := ioutil.ReadFile("lua.yaml")

	syntaxDef, err := highlight.ParseDef(syntaxFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	h := highlight.NewHighlighter(syntaxDef)
	matches := h.HighlightString(inputString)

	lines := strings.Split(inputString, "\n")
	for lineN, l := range lines {
		for colN, c := range l {
			if group, ok := matches[lineN][colN]; ok {
				if group == highlight.Groups["statement"] {
					color.Set(color.FgGreen)
				} else if group == highlight.Groups["preproc"] {
					color.Set(color.FgHiRed)
				} else if group == highlight.Groups["special"] {
					color.Set(color.FgBlue)
				} else if group == highlight.Groups["constant.string"] {
					color.Set(color.FgCyan)
				} else if group == highlight.Groups["constant.specialChar"] {
					color.Set(color.FgHiMagenta)
				} else if group == highlight.Groups["type"] {
					color.Set(color.FgYellow)
				} else if group == highlight.Groups["constant.number"] {
					color.Set(color.FgCyan)
				} else if group == highlight.Groups["comment"] {
					color.Set(color.FgHiGreen)
				} else {
					color.Unset()
				}
			}
			fmt.Print(string(c))
		}
		if group, ok := matches[lineN][len(l)]; ok {
			if group == highlight.Groups["default"] || group == highlight.Groups[""] {
				color.Unset()
			}
		}

		fmt.Print("\n")
	}
}
