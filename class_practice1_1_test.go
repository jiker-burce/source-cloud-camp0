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
	arms [2]Arm
}

type Arm struct {
	kind  string
	color string
}

func (h *Human) eat(sth string) {
	var _abc int

	fmt.Println(h.name, ": you are eating: "+sth, _abc)
}

func (h *Human) run() {
	fmt.Println(h.name, ": you are running now~")
}

/**
上面定义了2组接口，其中eat方法是2组里面都包含的，
编译时会检查
只会在将其赋值给某组接口时，如果发现调用的某个方法存在，且入参和返回值结构都一致，就会畅通无阻的执行；否则就会报错
*/
func TestInterface(t *testing.T) {
	h := Human{name: "Bruce", arms: [2]Arm{
		{"left", "yellow"},
		{"right", "red"},
	}}

	a2 := h            // 创建副本
	var a1 Action = &h // 指针引用，同时显式转换为某个接口指向

	a1.eat("apple")
	a1.run()

	h.name = "John"
	a1.eat("apple")
	a1.run()

	a2.eat("orange")
	fmt.Println("a2 to string: ", a2)
}

type Dog struct {
}

func (d *Dog) eat() {
	fmt.Println("eating")
}

type Pet interface {
	eat()
}

func TestDog(t *testing.T) {
	var dog1 *Dog
	fmt.Printf("The first dog is nil. [wrap1] %T => %v \n", dog1, dog1)
	dog2 := dog1
	fmt.Printf("The second dog is nil. [wrap1] %T => %v \n", dog2, dog2)
	var pet Pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Printf("The pet is not nil. [wrap1] %T => %v \n", pet, pet)
	}
}
