package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Operators interface {
	Generate(int, int) string
	Degenerate(string) (int, int, error)
}

type keyOp struct {
	Template string
}

func GetKeyOperator() Operators {
	op := &keyOp{
		Template: "%v_%v",
	}

	return op
}

func (kp *keyOp) Generate(x, y int) string {
	return fmt.Sprintf(kp.Template, x, y)
}

func (kp *keyOp) Degenerate(s string) (int, int, error) {

	if idx := strings.IndexByte(s, '_'); idx >= 0 {
		before, _ := strconv.Atoi(s[:idx])
		after, _ := strconv.Atoi(s[idx+1:])

		return before, after, nil
	} else {
		return 0, 0, fmt.Errorf("not_a_valid_key")
	}
}
