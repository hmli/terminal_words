package main


import (
	"errors"
	// "os"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	errInvalidGBK = errors.New("simplifiedchinese: invalid GBK encoding")
)

// transform  converts one rune from UTF-8 to GBK
func transform(p []byte) (gbk []byte, err error) {
	if len(p) == 0 {
		return nil, errInvalidGBK
	}
	if p[0] < utf8.RuneSelf {
		return nil, errInvalidGBK
	}
	if utf8.RuneCount(p) != 1 {
		return nil, errInvalidGBK
	}
	gbk = make([]byte, len(p))
	nDst, _, err := simplifiedchinese.GBK.NewEncoder().Transform(gbk, p, false)
	if err != nil {
		return nil, err
	}
	return gbk[:nDst], nil
}

// quweima计算汉字对应的区位码
func quweima(gbk []byte) []byte {
	b := make([]byte, len(gbk))
	for i, _ := range gbk {
		b[i] = gbk[i] - 0xa0
	}
	return b
}

// offset从区位码计算汉字点阵的偏移量
func computeOffset(qwm []byte) int64 {
	offset := (94*(int64(qwm[0])-16) + (int64(qwm[1]) - 1)) * 288
	return offset
}

// Matrix生成汉字的点阵
func Matrix(p []byte) ([]byte, error) {
	gbk, err := transform(p)
	if err != nil {
		return nil, err
	}
	qwm := quweima(gbk)
	offset := computeOffset(qwm)
	// file, err := os.Open("hzk/HZK16")
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()
	// _, err = file.Seek(offset, 0)
	// if err != nil {
	// 	return nil, err
	// }
	asset, err := Asset("HZK48S")
	if err != nil {
		return nil, err
	}
	var b = make([]byte, 288)
	// _, err = file.Read(b)
	// if err != nil {
	// 	return nil, err
	// }
	b = asset[offset : offset+288]
	return b, nil
}
