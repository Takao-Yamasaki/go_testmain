package main_test

import (
	"fmt"
	"testing"
)

func setup() {
	fmt.Println("setup")
}

func teardown() {
	fmt.Println("teardown")
}

func TestMain(m *testing.M) {
	// 前処理
	setup()

	m.Run()

	// 後処理
	teardown()
}

// 実行したいユニットテスト
func TestA(t *testing.T) {
	fmt.Println("testA")
}

func TestB(t *testing.T) {
	fmt.Println("testB")
}
