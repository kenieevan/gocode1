package main

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//can't invoke fun in const because there is no in compile time
const (
	kb = 1000
	mb = 1000 * kb
	gb = 1000 * mb
)

func main() {
	fmt.Println(kb)
	fmt.Println(mb)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	var tmp int8 = -65
	//var tmp1 int8
	//tmp1 = tmp << 1

	fmt.Println(tmp)
	var f float64 = 1677721600
	var f1 float64 = 2e2

	fmt.Println(f == (f + 1))
	fmt.Println(f1)
	s := "hello, \u4e16\u754c"
	for i, r := range s {
		fmt.Printf("%d  %q  %d\n", i, r, r)
	}

	n := 0
	for range s {
		n++
	}
	fmt.Printf("len is %d\n", n)
	fmt.Printf("len is %d\n", utf8.RuneCountInString(s))
	spath := path.Base("http://www.baidu.com/")

	fmt.Printf("%q \n", spath)

	spath = path.Base("/root/a")

	fmt.Printf("%q \n", spath)
	spath = path.Base("c:\test\a")

	fmt.Printf("%q \n", spath)
	spath = filepath.Base("c:\foo\bar")
	fmt.Printf("%q \n", spath)

	s = "hello"
	sarray := strings.Fields("h e ll l o")
	for _, stmp := range sarray {
		fmt.Println(stmp)
	}

	b := []byte(s) //slice is not comparable
	b[0] = 'a'
	b[0] = 'h'

	s1 := string(b)

	if s1 == s {
		fmt.Println("two string of different reference but same content")
	} else {

		fmt.Println("two string of different reference and differet content")
	}

	b1 := []byte(s) //slice is not comparable
	b1[0] = 'a'
	fmt.Printf("%c ", b[0]) //h

	fmt.Printf("bytes compare %d\n", bytes.Compare(b, b1))

	s1 = int2String([]int{1, 2, 3})
	fmt.Println(s1)

	fmt.Println(comma("12456"))

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 10))

	yint, err := strconv.ParseInt("1235", 10, 8)
	fmt.Printf("%d %v", yint, err)

}

//1,2,3 int  to  1:2:3 in string
func int2String(values []int) string {

	var buf bytes.Buffer

	for _, v := range values {
		fmt.Fprintf(&buf, "%d", v)
		buf.WriteString(":")
	}

	return buf.String()
}

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		return s
	}

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		count++
		if count == 3 {
			buf.WriteRune(',')
			count = 0
		}
	}

	b1 := buf.Bytes()
	reverse(b1)
	return string(b1)
	//	//reverse the bytes
	//	tmps := buf.String()
	//	var rbuf bytes.Buffer
	//	for i := len(tmps) - 1; i >= 0; i-- {
	//		rbuf.WriteByte(tmps[i])
	//	}
	//
	//	return rbuf.String()
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
