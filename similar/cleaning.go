package similar

import (
	"github.com/caneroj1/stemmer"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)


var reg00, _ = regexp.Compile(`\n?\r`)
var reg01, _ = regexp.Compile(`\n`)
var reg02, _ = regexp.Compile(`\[.*?]`)
var reg03, _ = regexp.Compile(`\(source: [^)]*\)`)
var reg04, _ = regexp.Compile(`\(from: [^)]*\)`)
var reg05, _ = regexp.Compile(`<[^>]*>`)
var reg06, _ = regexp.Compile(`^https?:\/\/.*[\r\n]*`)
var reg07, _ = regexp.Compile(`^http?:\/\/.*[\r\n]*`)
var reg08, _ = regexp.Compile(`[\w\.-]+@[\w\.-]+`)
var reg09, _ = regexp.Compile(`[^a-zA-Z0-9 ]+`)
var reg10, _ = regexp.Compile(`\s+`)

func CleanTitle(strRaw string) string {

	// Remove all symbols (clean to normal english)
	strRaw = reg09.ReplaceAllString(strRaw, "")

	// Remove multiple spaces
	strRaw = reg10.ReplaceAllString(strRaw, " ")

	// Stemming (porter)
	strRawArray := strings.Split(strRaw, " ")
	stemmer.StemMultipleMutate(&strRawArray)
	strRaw = strings.Join(strRawArray, " ")

	// To lowercase
	strRaw = strings.ToLower(strRaw)

	// Finally return
	return strRaw

}

func CleanDescription(strRaw string) string {

	// Remove all non-english descriptions
	// This assumes the english one is first
	// https://github.com/CarlosEsco/Neko/blob/master/app/src/main/java/eu/kanade/tachiyomi/source/online/utils/MdUtil.kt
	for _, tag := range DescriptionLanguages {
		strRaw = strings.Split(strRaw, tag)[0]
	}

	// Remove "rune" / umlauts / diacritics
	// https://stackoverflow.com/a/26722698/7718197
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	strRaw, _, _ = transform.String(t, strRaw)

	// To lowercase
	strRaw = strings.ToLower(strRaw)

	// Replace new lines with space
	strRaw = reg00.ReplaceAllString(strRaw, " ")
	strRaw = reg01.ReplaceAllString(strRaw, " ")

	// Now remove all english tags which are no longer needed
	for _, tag := range EnglishDescriptionTags {
		strRaw = strings.ReplaceAll(strRaw, tag, "")
	}

	// Next clean the string from any bbcodes
	for _, tag := range BBCodes {
		strRaw = strings.ReplaceAll(strRaw, tag, "")
	}
	strRaw = reg02.ReplaceAllString(strRaw, "")

	// Remove source parentheses typical of anilist
	// Eg: (source: solitarycross), (source: eat manga)
	strRaw = reg03.ReplaceAllString(strRaw, " ")
	strRaw = reg04.ReplaceAllString(strRaw, " ")

	// Remove any html codes
	strRaw = reg05.ReplaceAllString(strRaw, " ")

	// Remove emails and urls
	strRaw = reg06.ReplaceAllString(strRaw, " ")
	strRaw = reg07.ReplaceAllString(strRaw, " ")
	strRaw = reg08.ReplaceAllString(strRaw, " ")

	// Replace apostrophes with standard lexicons
	strRaw = strings.ReplaceAll(strRaw, "isn't", "is not")
	strRaw = strings.ReplaceAll(strRaw, "aren't", "are not")
	strRaw = strings.ReplaceAll(strRaw, "ain't", "am not")
	strRaw = strings.ReplaceAll(strRaw, "won't", "will not")
	strRaw = strings.ReplaceAll(strRaw, "didn't", "did not")
	strRaw = strings.ReplaceAll(strRaw, "shan't", "shall not")
	strRaw = strings.ReplaceAll(strRaw, "haven't", "have not")
	strRaw = strings.ReplaceAll(strRaw, "hadn't", "had not")
	strRaw = strings.ReplaceAll(strRaw, "hasn't", "has not")
	strRaw = strings.ReplaceAll(strRaw, "don't", "do not")
	strRaw = strings.ReplaceAll(strRaw, "wasn't", "was not")
	strRaw = strings.ReplaceAll(strRaw, "weren't", "were not")
	strRaw = strings.ReplaceAll(strRaw, "doesn't", "does not")
	strRaw = strings.ReplaceAll(strRaw, "'s", " is")
	strRaw = strings.ReplaceAll(strRaw, "'re", " are")
	strRaw = strings.ReplaceAll(strRaw, "'m", " am")
	strRaw = strings.ReplaceAll(strRaw, "'d", " would")
	strRaw = strings.ReplaceAll(strRaw, "'ll", " will")

	// Remove all symbols (clean to normal english)
	strRaw = reg09.ReplaceAllString(strRaw, "")

	// Remove multiple spaces
	strRaw = reg10.ReplaceAllString(strRaw, " ")

	// Stemming (porter)
	strRawArray := strings.Split(strRaw, " ")
	stemmer.StemMultipleMutate(&strRawArray)
	strRaw = strings.Join(strRawArray, " ")

	// To lowercase (again since stemmer makes upper)
	strRaw = strings.ToLower(strRaw)

	// Finally return
	return strRaw
}