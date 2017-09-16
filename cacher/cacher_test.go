package cacher

import (
	"encoding/json"
	"fmt"
	"testing"
)

var c Cacher

type testObj struct {
	ID   int64
	Name string
}

func TestCacher_NewCache(t *testing.T) {
	res := c.NewCache(1)
	if res != true {
		t.Fail()
	}
}

func TestCacher_Set(t *testing.T) {
	// test simple string val
	err := c.Set("test", "test value", 60)
	if err != nil {
		t.Fail()
	}
}

func TestCacher_Get(t *testing.T) {
	// test simple string val
	val, err := c.Get("test")
	fmt.Print("value: ")
	fmt.Println(val)
	if err != nil || val != "test value" {
		t.Fail()
	}
}

func TestCacher_Set2(t *testing.T) {
	// test simple string val
	err := c.Set("test", "test value 2", 60)
	if err != nil {
		t.Fail()
	}
}

func TestCacher_Get2(t *testing.T) {
	// test simple string val
	val, err := c.Get("test")
	fmt.Print("value: ")
	fmt.Println(val)
	if err != nil || val != "test value 2" {
		t.Fail()
	}
}

func TestCacher_Set3(t *testing.T) {
	// test simple string val
	o := new(testObj)
	//var o testObj
	o.ID = 2
	o.Name = "test value 3"
	oo, err := json.Marshal(o)
	if err != nil {
		t.Fail()
	}
	err2 := c.Set("test3", string(oo), 60)
	if err2 != nil {
		t.Fail()
	}
}

func TestCacher_Get3(t *testing.T) {
	// test simple string val
	val, err := c.Get("test3")
	fmt.Print("value3: ")
	t3 := new(testObj)
	err2 := json.Unmarshal([]byte(val), t3)
	fmt.Println(val)
	if err != nil || err2 != nil || t3.Name != "test value 3" {
		t.Fail()
	}
}

func TestCacher_SetObj(t *testing.T) {
	// test opj val
	o := new(testObj)
	//var o testObj
	o.ID = 2
	o.Name = "test value"
	oo, err := json.Marshal(o)
	if err != nil {
		t.Fail()
	}
	err2 := c.SetObj("133:test:obj", oo, 60)
	if err2 != nil {
		t.Fail()
	}
}

func TestCacher_GetObj(t *testing.T) {
	// test simple string val
	val, err := c.GetObj("133:test:obj")
	var o testObj
	err2 := json.Unmarshal(val, &o)
	if err2 != nil {
		t.Fail()
	}
	fmt.Print("obj value: ")
	fmt.Println(o)
	if err != nil || o.Name != "test value" {
		t.Fail()
	}
}

func TestCacher_Delete(t *testing.T) {
	// test simple string val
	val1 := c.Delete("test")
	val2 := c.Delete("133:test:obj")
	if val1 != true || val2 != true {
		t.Fail()
	}
}
