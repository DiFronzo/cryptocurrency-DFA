package dfa

import "unicode/utf8"

func ValidateEthAddress(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	_, _, _ = r, rlen, i
	switch {
	case i == 0:
		goto s2
	}
	return
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 48:
		goto s3
	}
	return
s3:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 120:
		goto s4
	}
	return
s4:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s5
	}
	return
s5:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s6
	}
	return
s6:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s7
	}
	return
s7:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s8
	}
	return
s8:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s9
	}
	return
s9:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s10
	}
	return
s10:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s11
	}
	return
s11:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s12
	}
	return
s12:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s13
	}
	return
s13:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s14
	}
	return
s14:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s15
	}
	return
s15:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s16
	}
	return
s16:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s17
	}
	return
s17:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s18
	}
	return
s18:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s19
	}
	return
s19:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s20
	}
	return
s20:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s21
	}
	return
s21:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s22
	}
	return
s22:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s23
	}
	return
s23:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s24
	}
	return
s24:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s25
	}
	return
s25:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s26
	}
	return
s26:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s27
	}
	return
s27:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s28
	}
	return
s28:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s29
	}
	return
s29:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s30
	}
	return
s30:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s31
	}
	return
s31:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s32
	}
	return
s32:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s33
	}
	return
s33:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s34
	}
	return
s34:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s35
	}
	return
s35:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s36
	}
	return
s36:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s37
	}
	return
s37:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s38
	}
	return
s38:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s39
	}
	return
s39:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s40
	}
	return
s40:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s41
	}
	return
s41:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s42
	}
	return
s42:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s43
	}
	return
s43:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 70 || r >= 97 && r <= 102:
		goto s44
	}
	return
s44:
	switch {
	case i == len(s):
		end = i
	}
	return
}
