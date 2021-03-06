package mysql

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func readString(src []byte, srcCharSet charSet) string {
	switch srcCharSet {
	case charSetUjis, charSetEucjpms:
		buff := bytes.NewBuffer(src)
		dst, err := ioutil.ReadAll(transform.NewReader(buff, japanese.EUCJP.NewDecoder()))
		if err != nil {
			return strings.TrimRight(string(src), "\x00")
		}
		return strings.TrimRight(string(dst), "\x00")
	case charSetSjis, charSetCp932:
		buff := bytes.NewBuffer(src)
		dst, err := ioutil.ReadAll(transform.NewReader(buff, japanese.ShiftJIS.NewDecoder()))
		if err != nil {
			return strings.TrimRight(string(src), "\x00")
		}
		return strings.TrimRight(string(dst), "\x00")
	default:
		return strings.TrimRight(string(src), "\x00")
	}
}
