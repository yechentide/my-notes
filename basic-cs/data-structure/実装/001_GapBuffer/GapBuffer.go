package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- //
// ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- //

const InitGapSize = 8

type GapBuffer struct {
	buffer   []byte
	gapStart int
	gapSize  int
}

// ---------- ---------- ---------- 构造方法 Constructor ---------- ---------- ---------- //
func newGapBuffer() *GapBuffer {
	// 创建一个新的空白的gapBuffer结构体，并返回指针
	return &GapBuffer{make([]byte, 8, 8), 0, InitGapSize}
}

// ---------- ---------- ----------     方法 Method     ---------- ---------- ---------- //

func (gb *GapBuffer) insert(index int, str string) {
	for i, c := range str {
		gb.insertAt(index+i, byte(c))
	}
}

func (gb *GapBuffer) insertAt(index int, char byte) {
	if index<0 || gb.length()<index {
		fmt.Println("ERROR: insertAt() <<< 指定的位置", index, "错误")
		return
	}

	gb.expandGap()
	gb.moveGapToPoint(index)
	gb.buffer[gb.gapStart] = char
	gb.gapStart++
	gb.gapSize--

}

func (gb *GapBuffer) removeRange(start int, size int) {
	if start > gb.length() {
		fmt.Println("ERROR: removeRange() <<< 指定的位置", start, "错误")
		return
	}
	length := size
	if (start+size)>gb.length() {
		length = gb.length()-start
	}
	for i:=0; i<length; i++ {
		gb.remove(start)
	}
}

func (gb *GapBuffer) remove(index int) {
	if index<0 || gb.length()<=index {
		fmt.Println("ERROR: remove() <<< 指定的位置", index, "错误")
		return
	}

	if gb.gapSize == 0 {
		gb.gapStart = index
		gb.gapSize += 1
		return
	}

	if (index+1)==gb.gapStart {
		gb.gapStart = index
		gb.gapSize += 1
	} else if index==gb.gapStart {
		gb.gapSize += 1
	} else {
		if index < gb.gapStart {
			gb.moveGapToPoint(index+1)
			gb.gapStart = index
			gb.gapSize += 1
		} else {
			gb.moveGapToPoint(index)
			gb.gapSize += 1
		}
	}

}

func (gb *GapBuffer) expandGap() {
	if gb.gapSize == 0 {
		gb.gapStart = gb.length()
		gb.gapSize = InitGapSize
		gb.buffer = append(gb.buffer, make([]byte, InitGapSize)...)
	}
}

func (gb *GapBuffer) moveGapToPoint(point int) {
	if point == gb.gapStart {
		return
	}

	if point < gb.gapStart {
		// 复制顺序：从后往前
		// point=0, gapStart=7, gapSize=3
		// abcdefg...
		// abcdef...g
		// abcde...fg
		// ...abcdefg
		for i:=0; i<(gb.gapStart-point); i++ {
			gb.buffer[gb.gapStart+gb.gapSize-1-i] = gb.buffer[gb.gapStart-1-i]
		}
	} else if gb.gapStart < point {
		// 复制顺序：从前往后
		// point=10, gapStart=0, gapSize=3
		// ...abcdefg
		// a...bcdefg
		// ab...cdefg
		// abcdefg...
		for i:=0; i<(point-gb.gapStart); i++ {
			gb.buffer[gb.gapStart+i] = gb.buffer[gb.gapStart+gb.gapSize+i]
		}
	}
	gb.gapStart = point
}


// ---------- ---------- 不重要的Method

func (gb *GapBuffer) length() int {
	return len(gb.buffer) - gb.gapSize
}

func (gb *GapBuffer) getString() string {
	s := ""
	for i, c := range gb.buffer {
		if gb.gapStart <= i && i < gb.gapStart+gb.gapSize {
		} else {
			s += string(c)
		}
	}
	return s
}

func (gb *GapBuffer) getStatus() string {
	s := ""
	for i, c := range gb.buffer {
		if gb.gapStart <= i && i < gb.gapStart+gb.gapSize {
			s += "_"
		} else {
			s += string(c)
		}
	}
	return s
}

func (gb *GapBuffer) debug() {
	fmt.Println("========================================")
	fmt.Println("字符串        :", gb.getString())
	fmt.Println("字符串长度     :", gb.length())
	fmt.Println("buffer       :", gb.getStatus())
	fmt.Println("buffer's size:", len(gb.buffer))
	fmt.Println("gapStart:", gb.gapStart, "gapSize:", gb.gapSize)
}

// ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- //
// ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- ---------- //

func main() {
	fmt.Println("\n\n\n===================   新建GapBuffer")
	gb := newGapBuffer()
	gb.debug()

	sc := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("\n\n\n>>> ")
		if sc.Scan() {
			slice := strings.Split(sc.Text(), " ")
			if slice[0] == "over" || len(slice[0]) == 0 {
				break
			}

			if slice[0] == "add" {
				index, _ := strconv.Atoi(slice[1])
				str := slice[2]
				fmt.Println("===================   在位置", index, "插入字符串", str)
				gb.insert(index, str)
				gb.debug()
			} else if slice[0] == "rm" {
				index, _ := strconv.Atoi(slice[1])
				size := 1
				if len(slice) == 3 {
					lll, _ := strconv.Atoi(slice[2])
					size = lll
				}
				fmt.Println("===================   删除位置", index, "开始的", size, "个字符")
				gb.removeRange(index, size)
				gb.debug()
			} else {
				fmt.Println("请输入正确的指令")
				fmt.Println("例子：")
				fmt.Println("add 1 kkk          // 在位置1插入字符串kkk")
				fmt.Println("rm 1               // 删除位于位置1的字符")
				fmt.Println("rm 1 4             // 删除位置1开始的4个字符")
			}
		}
	}


	fmt.Println("\n\n====== 结束程序 ======")
}
