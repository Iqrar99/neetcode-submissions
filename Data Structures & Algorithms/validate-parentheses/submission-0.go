type BracketStack struct {
    Length  int
    Stack   []rune
}

func (bs *BracketStack) Push(p rune) {
    bs.Stack = append(bs.Stack, p)
    bs.Length++
}

func (bs *BracketStack) Pop() {
    bs.Stack = bs.Stack[:len(bs.Stack)-1]
    bs.Length--
}

func (bs *BracketStack) Top() rune {
    return bs.Stack[len(bs.Stack)-1]
}

func isValid(s string) bool {
    var stack BracketStack
    reverseBracket := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, p := range s {
        if stack.Length == 0 {
            stack.Push(p)
            continue
        }

        if stack.Top() == reverseBracket[p] {
            stack.Pop()
        } else {
            stack.Push(p)
        }
    }

    if stack.Length == 0 {
        return true
    }
    return false
}
