package main

import (
	"fmt"
	"strings"
)

func main() {
	test := ""

	test_list := strings.Split(test, ",")

	fmt.Println(test_list)	// []

	ret := SplitTextSegment("hello,世界热名号解耦", 5, 1, 10)

	ret1 := SplitTextSegmentV1("hello,世界热名号解耦", 5, 1, 10)


	strings.Compare()

	fmt.Printf("ret: %+v", ret)
	fmt.Printf("ret: %+v", ret1)

}

func SplitTextSegmentV1(src string, checkLimit, checkOffset, reportLimit int) []string {
	var (
		dstlist []string
		start   = 0
		end     = 0
	)

	//src := []rune(srcByte)

	fmt.Printf("src: %v", src)

	if len(src) <= 0 {
		return dstlist
	}

	if len(src) > reportLimit {
		src = src[:reportLimit]
	}

	for {
		end = start + checkLimit
		if end > len(src) {
			end = len(src)
		}

		dstlist = append(dstlist, string(src[start:end]))

		start = end - checkOffset

		if end >= len(src) {
			break
		}
	}

	return dstlist
}

func SplitTextSegment(srcByte string, checkLimit, checkOffset, reportLimit int) []string {
	var (
		dstlist []string
		start   = 0
		end     = 0
	)

	src := []rune(srcByte)

	fmt.Printf("src: %v", src)

	if len(src) <= 0 {
		return dstlist
	}

	if len(src) > reportLimit {
		src = src[:reportLimit]
	}

	for {
		end = start + checkLimit
		if end > len(src) {
			end = len(src)
		}

		dstlist = append(dstlist, string(src[start:end]))

		start = end - checkOffset

		if end >= len(src) {
			break
		}
	}

	return dstlist
}