package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMultiline(t *testing.T) {
	str1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	fmt.Println(len(str1))
	res1 := StringToMultiline(str1, 20)
	res1s := strings.Split(res1, "\n")
	assert.Equal(t, 19, len(res1s))

	str2 := "Lorem ipsum dolor sit amet,"
	res2 := StringToMultiline(str2, 20)
	res2s := strings.Split(res2, "\n")
	assert.Equal(t, 2, len(res2s))

	str3 := "Lorem ipsum dolor"
	res3 := StringToMultiline(str3, 20)
	res3s := strings.Split(res3, "\n")
	assert.Equal(t, 1, len(res3s))
}
