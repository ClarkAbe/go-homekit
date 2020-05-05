package util

import (
	"unicode"

	"ccms/x/text/secure/precis"
	"ccms/x/text/transform"
	"ccms/x/text/unicode/norm"
)

// RemoveAccentsFromString removes accent characters from string
// From https://stackoverflow.com/a/40405242/424814
func RemoveAccentsFromString(v string) string {
	var loosecompare = precis.NewIdentifier(
		precis.AdditionalMapping(func() transform.Transformer {
			return transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
				return unicode.Is(unicode.Mn, r)
			}))
		}),
		precis.Norm(norm.NFC), // This is the default; be explicit though.
	)
	p, _ := loosecompare.String(v)
	return p
}
