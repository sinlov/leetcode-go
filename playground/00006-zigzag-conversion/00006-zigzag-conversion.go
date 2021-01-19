package _00006_zigzag_conversion

import "bytes"

func convert(s string, numRows int) string {
	if len(s) < 1 || numRows == 1 { // special case
		return s
	}

	var temp = make([]bytes.Buffer, numRows)
	length := len(s)

	for i := 0; i < length; i++ {
		ans := i / (numRows - 1)
		cur := i % (numRows - 1)
		if ans%2 == 0 { // even or 0
			temp[cur].WriteByte(s[i]) // in positive order
		} else { // uneven
			temp[numRows-cur-1].WriteByte(s[i]) // in reverse order
		}
	}

	var res bytes.Buffer
	for _, v := range temp {
		res.Write(v.Bytes())
	}
	return res.String()
}
