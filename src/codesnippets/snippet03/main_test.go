package main

import (
	"testing"
	"log"
	"io/ioutil"
)

//这个测试模块仅仅测试一次
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	m.Run()
}

func TestDoAndLogSomething(t *testing.T) {
	if res := DoAndLogSomething(); res != 10 {
		t.Errorf("DoAndLogSomething resultd and an unexpcted number:got %v want %v ", res, 10)
	}
}
