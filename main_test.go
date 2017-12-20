package main

import "testing"

func TestMain(t *testing.T) {
	//main()
}

func TestDo(t *testing.T) {
	err := do(list)
	if err != nil {
		t.Error("Error occurs: ", err)
	}
}
