package cncamp

import (
	"fmt"
	"testing"
)

func TestReplaceByRepArr(t *testing.T) {
	oriArr := []string{"I", "am", "stupid", "and", "weak"}
	replaceArr := []string{"", "", "smart", "", "strong"}

	target := replaceByRepArr2(oriArr, replaceArr)

	t.Log(oriArr)
	t.Log(target)
}

func TestReplaceByIndex(t *testing.T) {
	oriArr := []string{"I", "am", "stupid", "and", "weak"}
	replaceArr := map[int]string{2: "smart", 4: "strong"}

	target := replaceByIndex(oriArr, replaceArr)

	t.Log(oriArr)
	t.Log(target)
}

func TestSwapValue(t *testing.T) {
	swapValue()
}

type Action interface {
	eat(food string)
	run()
}

type Action2 interface {
	eat(vegetable string)
}

type Human struct {
	name string
}

func (h *Human) eat(sth string) {
	var _abc int

	fmt.Println("you are eating: "+sth, _abc)
}

//func TestInterface()
