package test

import (
	"go-sample/internal/tool/clone"
	"testing"
)

type User struct {
	Name string
	Age int
	Locked bool  `copier:"Lock"`  // 以别名为准
	IgnoringVal string `copier:"-"` // 忽略此属性
	EmptyVal string 				//忽略空值
}

type Employee struct {
	Name string
	Age int
	Lock bool
	EmptyVal string
}


// Test Copy
func TestStructCopy(t *testing.T) {
	u := User{
		Name:        "Admin",
		Age:         18,
		Locked:      true,
		IgnoringVal: "xxx",
	}
	e := Employee{}
	err := clone.Copy(&e, &u)
	if err != nil {
		t.Error(err)
	}
	t.Log(e)

	u2, err := clone.Clone[User](&e)
	if err != nil {
		t.Error(err)
	}
	t.Log(*u2)
}
