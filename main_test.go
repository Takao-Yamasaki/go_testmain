package main_test

import (
	"fmt"
	"testing"
)

// 実行したいユニットテスト
func TestTableDriveParallel(t *testing.T) {
	// defer文で実行する後処理の記述
	t.Cleanup(func() {
		fmt.Println("cleanup")
	})

	// 本来のテストの記述
	tests := []struct {
		testTitle string
	}{
		{testTitle: "subtest1"},
		{testTitle: "subtest2"},
		{testTitle: "subtest3"},
	}

	for _, test := range tests {
		testcase := test
		t.Run(testcase.testTitle, func(t *testing.T) {
			// これがあることでサブテストが並列に走る
			t.Parallel()
			fmt.Println(testcase.testTitle)
		})
	}
}
