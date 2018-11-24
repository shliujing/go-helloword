package learn

import (
	"fmt"
)

type Man interface {
	name() string;
	age() int;
}

type Woman struct {
	sexy bool;
}

func (woman Woman) name() string {
	if woman.sexy {
		return "Jin Yawei, very sexy"
	} else {
		return "Jin Yawei, normal sexy"
	}
}

func (woman Woman) age() int {
	return 23;
}

type Men struct {
}

func (men Men) name() string {
	return "liweibin";
}
func (men Men) age() int {
	return 27;
}

func main() {
	var man Man;
	//man = new(Woman);
	man = Woman{
		sexy: true,
	}

	fmt.Println(man.name());
	fmt.Println(man.age());
	man = new(Men);
	fmt.Println(man.name());
	fmt.Println(man.age());
}
