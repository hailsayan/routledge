package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

type St struct {
	s   string
	op  string
	idx int
}

type Stack []St

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string, idx int, op string) {
	*s = append(*s, St{str, op, idx})
}

func (s *Stack) Pop() (St, bool) {
	if s.IsEmpty() {
		return St{}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Head() (St, bool) {
	if s.IsEmpty() {
		return St{}, false
	} else {
		return (*s)[len(*s)-1], true
	}
}

type History struct {
	s         []string
	undoStack Stack
}

func (h *History) insert(i int, x string, undo bool) {
	if i <= len(h.s) {
		h.s = append(h.s[:i], append([]string{x}, h.s[i:]...)...)
		if !undo {
			h.undoStack = append(h.undoStack, St{x, "insert", i})
		}
	}
}

func (h *History) delete(i int, undo bool) {
	deleteValue := (h.s)[i]
	h.s = append((h.s)[:i], (h.s)[i+1:]...)
	if !undo {
		h.undoStack.Push(deleteValue, i, "delete")
	}
}

func (h *History) undo() {
	if len(h.undoStack) == 0 {
		return
	}
	undoValue, e := h.undoStack.Pop()
	if !e {
		return
	}
	if undoValue.op == "insert" {
		h.delete(undoValue.idx, true)
	} else {
		h.insert(undoValue.idx, undoValue.s, true)
	}
}

func main() {
	var history History
	var n, idx int
	var op, str string
	scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		scanf("%s %d %s\n", &op, &idx, &str)
		if op == "insert" {
			history.insert(idx-1, str, false)
		} else if op == "delete" {
			history.delete(idx-1, false)
		} else {
			history.undo()
		}
	}

	result := history.s
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
