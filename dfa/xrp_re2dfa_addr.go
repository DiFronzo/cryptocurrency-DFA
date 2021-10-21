package dfa

import "unicode/utf8"

func ValidateXrpAddress(s string) (end int) {
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
	case r == 114:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
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
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s27
	}
	return
s27:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s29
	}
	return
s29:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s30
	}
	return
s30:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s31
	}
	return
s31:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s32
	}
	return
s32:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s33
	}
	return
s33:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s34
	}
	return
s34:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s35
	}
	return
s35:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s36
	}
	return
s36:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s37
	}
	return
s37:
	switch {
	case i == len(s):
		end = i
		return
	}
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122:
		goto s38
	}
	return
s38:
	switch {
	case i == len(s):
		end = i
	}
	return
}
