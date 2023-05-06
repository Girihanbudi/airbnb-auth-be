package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Casers []cases.Caser

var (
	CaseLower = cases.Lower(language.Und)
	CaseTitle = cases.Title(language.Und)
	CaseUpper = cases.Upper(language.Und)
)

func Case(text string, casers ...cases.Caser) string {
	for _, caser := range casers {
		text = caser.String(text)
	}

	return text
}
