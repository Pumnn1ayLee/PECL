package src

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"sort"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

const unicodedataheader = `
// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

package unicodedata

// Code generated by typesettings-utils/generators/unicodedata/cmd/main.go DO NOT EDIT.


`

func sortRunes(rs []rune) {
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
}

// compacts the code more than a "%#v" directive
func printTable(rt *unicode.RangeTable, omitTypeLitteral bool) string {
	w := new(bytes.Buffer)
	if omitTypeLitteral {
		fmt.Fprintln(w, "{")
	} else {
		fmt.Fprintln(w, "&unicode.RangeTable{")
	}
	if len(rt.R16) > 0 {
		fmt.Fprintln(w, "\tR16: []unicode.Range16{")
		for _, r := range rt.R16 {
			fmt.Fprintf(w, "\t\t{Lo:%#04x, Hi:%#04x, Stride:%d},\n", r.Lo, r.Hi, r.Stride)
		}
		fmt.Fprintln(w, "\t},")
	}
	if len(rt.R32) > 0 {
		fmt.Fprintln(w, "\tR32: []unicode.Range32{")
		for _, r := range rt.R32 {
			fmt.Fprintf(w, "\t\t{Lo:%#x, Hi:%#x,Stride:%d},\n", r.Lo, r.Hi, r.Stride)
		}
		fmt.Fprintln(w, "\t},")
	}
	if rt.LatinOffset > 0 {
		fmt.Fprintf(w, "\tLatinOffset: %d,\n", rt.LatinOffset)
	}
	fmt.Fprintf(w, "}")
	return w.String()
}

func generateCombiningClasses(classes map[uint8][]rune, w io.Writer) {
	fmt.Fprint(w, unicodedataheader)

	// create and compact the tables
	var out [256]*unicode.RangeTable
	for k, v := range classes {
		if len(v) == 0 {
			return
		}
		out[k] = rangetable.New(v...)
	}

	// print them
	fmt.Fprintln(w, "var combiningClasses = [256]*unicode.RangeTable{")
	for i, t := range out {
		if t == nil {
			continue
		}
		fmt.Fprintf(w, "%d : %s,\n", i, printTable(t, true))
	}
	fmt.Fprintln(w, "}")
}

func generateEmojis(runes map[string][]rune, w io.Writer) {
	// among "Emoji", "Emoji_Presentation", "Emoji_Modifier", "Emoji_Modifier_Base", "Extended_Pictographic"
	// only Extended_Pictographic is actually used
	fmt.Fprint(w, unicodedataheader)
	for _, class := range [...]string{"Extended_Pictographic"} {
		table := rangetable.New(runes[class]...)
		s := printTable(table, false)
		fmt.Fprintf(w, "var %s = %s\n\n", class, s)
	}
}

func generateMirroring(runes map[uint16]uint16, w io.Writer) {
	fmt.Fprint(w, unicodedataheader)
	fmt.Fprintf(w, "var mirroring = map[rune]rune{ // %d entries \n", len(runes))
	var sorted []rune
	for r1 := range runes {
		sorted = append(sorted, rune(r1))
	}
	sortRunes(sorted)
	for _, r1 := range sorted {
		r2 := runes[uint16(r1)]
		fmt.Fprintf(w, "0x%04x: 0x%04x,\n", r1, r2)
	}
	fmt.Fprintln(w, "}")
}

func generateDecomposition(combiningClasses map[uint8][]rune, dms map[rune][]rune, compExp map[rune]bool, w io.Writer) {
	var (
		decompose1 [][2]rune         // length 1 mappings {from, to}
		decompose2 [][3]rune         // length 2 mappings {from, to1, to2}
		compose    [][3]rune         // length 2 mappings {from1, from2, to}
		ccc        = map[rune]bool{} // has combining class
	)
	for c, runes := range combiningClasses {
		for _, r := range runes {
			ccc[r] = c != 0
		}
	}
	for r, v := range dms {
		switch len(v) {
		case 1:
			decompose1 = append(decompose1, [2]rune{r, v[0]})
		case 2:
			decompose2 = append(decompose2, [3]rune{r, v[0], v[1]})
			var composed rune
			if !compExp[r] && !ccc[r] {
				composed = r
			}
			compose = append(compose, [3]rune{v[0], v[1], composed})
		default:
			log.Fatalf("unexpected runes for decomposition: %d, %v", r, v)
		}
	}

	// sort for determinisme
	sort.Slice(decompose1, func(i, j int) bool { return decompose1[i][0] < decompose1[j][0] })
	sort.Slice(decompose2, func(i, j int) bool { return decompose2[i][0] < decompose2[j][0] })
	sort.Slice(compose, func(i, j int) bool {
		return compose[i][0] < compose[j][0] ||
			compose[i][0] == compose[j][0] && compose[i][1] < compose[j][1]
	})

	fmt.Fprint(w, unicodedataheader)

	fmt.Fprintf(w, "var decompose1 = map[rune]rune{ // %d entries \n", len(decompose1))
	for _, vals := range decompose1 {
		fmt.Fprintf(w, "0x%04x: 0x%04x,\n", vals[0], vals[1])
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintf(w, "var decompose2 = map[rune][2]rune{ // %d entries \n", len(decompose2))
	for _, vals := range decompose2 {
		fmt.Fprintf(w, "0x%04x: {0x%04x,0x%04x},\n", vals[0], vals[1], vals[2])
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintf(w, "var compose = map[[2]rune]rune{ // %d entries \n", len(compose))
	for _, vals := range compose {
		fmt.Fprintf(w, "{0x%04x,0x%04x}: 0x%04x,\n", vals[0], vals[1], vals[2])
	}
	fmt.Fprintln(w, "}")
}

// Supported line breaking classes for Unicode 12.0.0.
// Table loading depends on this: classes not listed here aren't loaded.
var lineBreakClasses = [][2]string{
	{"BK", "Mandatory Break"},
	{"CR", "Carriage Return"},
	{"LF", "Line Feed"},
	{"NL", "Next Line"},
	{"SP", "Space"},
	{"NU", "Numeric"},
	{"AL", "Alphabetic"},
	{"IS", "Infix Numeric Separator"},
	{"PR", "Prefix Numeric"},
	{"PO", "Postfix Numeric"},
	{"OP", "Open Punctuation"},
	{"CL", "Close Punctuation"},
	{"CP", "Close Parenthesis"},
	{"QU", "Quotation"},
	{"HY", "Hyphen"},
	{"SG", "Surrogate"},
	{"GL", "Non-breaking (\"Glue\")"},
	{"NS", "Nonstarter"},
	{"EX", "Exclamation/Interrogation"},
	{"SY", "Symbols Allowing Break After"},
	{"HL", "Hebrew Letter"},
	{"ID", "Ideographic"},
	{"IN", "Inseparable"},
	{"BA", "Break After"},
	{"BB", "Break Before"},
	{"B2", "Break Opportunity Before and After"},
	{"ZW", "Zero Width Space"},
	{"CM", "Combining Mark"},
	{"EB", "Emoji Base"},
	{"EM", "Emoji Modifier"},
	{"WJ", "Word Joiner"},
	{"ZWJ", "Zero width joiner"},
	{"H2", "Hangul LV Syllable"},
	{"H3", "Hangul LVT Syllable"},
	{"JL", "Hangul L Jamo"},
	{"JV", "Hangul V Jamo"},
	{"JT", "Hangul T Jamo"},
	{"RI", "Regional Indicator"},
	{"CB", "Contingent Break Opportunity"},
	{"AI", "Ambiguous (Alphabetic or Ideographic)"},
	{"CJ", "Conditional Japanese Starter"},
	{"SA", "Complex Context Dependent (South East Asian)"},
	{"XX", "Unknown"},
}

func generateLineBreak(datas map[string][]rune, w io.Writer) {
	dict := ""

	fmt.Fprint(w, unicodedataheader)
	for i, class := range lineBreakClasses {
		className := class[0]
		table := rangetable.New(datas[className]...)
		s := printTable(table, false)
		fmt.Fprintf(w, "// %s\n", lineBreakClasses[i][1])
		fmt.Fprintf(w, "var Break%s = %s\n\n", className, s)

		dict += fmt.Sprintf("Break%s, // %s \n", className, className)
	}

	fmt.Fprintf(w, `var lineBreaks = [...]*unicode.RangeTable{
		%s}
	`, dict)
}

func generateEastAsianWidth(datas map[string][]rune, w io.Writer) {
	fmt.Fprint(w, unicodedataheader)
	// the table is used for UAX14 (LB30) : we group the classes
	// F (Fullwidth), W (Wide), H (Halfwidth)
	runes := append(append(datas["F"], datas["W"]...), datas["H"]...)
	table := rangetable.New(runes...)
	s := printTable(table, false)
	fmt.Fprintf(w, `
	// LargeEastAsian matches runes with East_Asian_Width property of 
	// F, W or H, and is used for UAX14, rule LB30.
	var LargeEastAsian = %s

	`, s)
}

func generateIndicCategories(datas map[string][]rune, w io.Writer) {
	fmt.Fprint(w, unicodedataheader)
	for _, className := range []string{"Virama", "Vowel_Dependent"} {
		table := rangetable.New(datas[className]...)
		s := printTable(table, false)
		fmt.Fprintf(w, "var Indic%s = %s\n\n", className, s)
	}
}

// only generate the table for the STerm property
func generateSTermProperty(datas map[string][]rune, w io.Writer) {
	fmt.Fprint(w, unicodedataheader)

	className := "STerm"
	table := rangetable.New(datas[className]...)
	s := printTable(table, false)
	fmt.Fprintf(w, "// SentenceBreakProperty: STerm\n")
	fmt.Fprintf(w, "var %s = %s\n\n", className, s)
}

func generateGraphemeBreakProperty(datas map[string][]rune, w io.Writer) {
	fmt.Fprint(w, unicodedataheader)

	var sortedClasses []string
	for key := range datas {
		sortedClasses = append(sortedClasses, key)
	}
	sort.Strings(sortedClasses)

	list := ""
	var allGraphemes []*unicode.RangeTable
	for _, className := range sortedClasses {
		runes := datas[className]
		table := rangetable.New(runes...)
		s := printTable(table, false)
		fmt.Fprintf(w, "// GraphemeBreakProperty: %s\n", className)
		fmt.Fprintf(w, "var GraphemeBreak%s = %s\n\n", className, s)

		list += fmt.Sprintf("GraphemeBreak%s, // %s \n", className, className)
		allGraphemes = append(allGraphemes, table)
	}

	// generate a union table to speed up lookup
	allTable := rangetable.Merge(allGraphemes...)
	fmt.Fprintf(w, "// contains all the runes having a non nil grapheme break property\n")
	fmt.Fprintf(w, "var graphemeBreakAll = %s\n\n", printTable(allTable, false))

	fmt.Fprintf(w, `var graphemeBreaks = [...]*unicode.RangeTable{
	%s}
	`, list)
}

func generateEmojisTest(sequences [][]rune, w io.Writer) {
	fmt.Fprintln(w, `
	// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

	package harfbuzz

	// Code generated by typesettings-utils/generators/unicodedata/cmd/main.go DO NOT EDIT.

	`)

	fmt.Fprintln(w, "var emojisSequences = [][]rune{")
	for _, seq := range sequences {
		if len(seq) < 2 {
			continue
		}
		fmt.Fprint(w, "{")
		for _, r := range seq {
			fmt.Fprintf(w, "0x%x,", r)
		}
		fmt.Fprintln(w, "},")
	}
	fmt.Fprintln(w, "}")
}

type item struct {
	start, end rune
	script     uint32
}

func compactScriptLookupTable(scripts map[string][]runeRange, scriptNames map[string]uint32) []item {
	var crible []item
	for scriptName, runes := range scripts {
		script, ok := scriptNames[scriptName]
		if !ok {
			check(fmt.Errorf("unknown script name %s", scriptName))
		}
		for _, ra := range runes {
			crible = append(crible, item{script: script, start: ra.Start, end: ra.End})
		}
	}

	sort.Slice(crible, func(i, j int) bool { return crible[i].start < crible[j].start })

	for i := range crible {
		if i == len(crible)-1 {
			continue
		}
		if crible[i].end >= crible[i+1].start {
			check(fmt.Errorf("inconsistent crible index %d: %v %v", i, crible[i], crible[i+1]))
		}
	}

	var compacted []item
	for _, v := range crible {
		if len(compacted) == 0 {
			compacted = append(compacted, v)
			continue
		}

		last := &compacted[len(compacted)-1]

		if v.script == last.script && v.start == last.end+1 { // merge
			last.end = v.end
		} else {
			compacted = append(compacted, v)
		}
	}

	return compacted
}

func scriptToString(s uint32) string {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], s)
	return string(buf[:])
}

func generateScriptLookupTable(scripts map[string][]runeRange, scriptNames map[string]uint32, w io.Writer) {
	crible := compactScriptLookupTable(scripts, scriptNames)

	fmt.Fprintln(w, `
	// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

	package language

	// Code generated by typesettings-utils/generators/unicodedata/cmd/main.go DO NOT EDIT.

	`)

	// scripts constants
	fmt.Fprintln(w, "const (")

	var sortedKeys []string
	for k := range scriptNames {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		v := scriptNames[k]
		fmt.Fprintf(w, "%s = Script(0x%08x) // %s \n", k, v, scriptToString(v))
	}
	fmt.Fprintln(w, ")")

	fmt.Fprintln(w, "var scriptToTag = map[string]Script{")
	for _, k := range sortedKeys {
		v := scriptNames[k]
		fmt.Fprintf(w, "%q : %d,\n", k, v)
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintln(w, `type scriptItem struct {
		start, end rune
		script     Script
	}
	
	var scriptRanges = [...]scriptItem{`)
	for _, item := range crible {
		fmt.Fprintf(w, "{start: 0x%x, end: 0x%x, script: 0x%08x},\n", item.start, item.end, item.script)
	}
	fmt.Fprintln(w, "}")
}

func generateGeneralCategories(m map[rune]string, w io.Writer) {
	// reverse the rune->category mapping
	cats, keys := map[string][]rune{}, []string{}
	for r, cat := range m {
		cats[cat] = append(cats[cat], r)
	}
	for key := range cats {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Fprint(w, unicodedataheader)

	for _, cat := range keys {
		runes := cats[cat]
		rt := rangetable.New(runes...)
		code := printTable(rt, false)
		fmt.Fprintln(w, fmt.Sprintf("var %s = %s\n", cat, code))
	}
}