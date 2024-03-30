package utils

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Printf(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$upslTuqNgfASXGDUffWBjewkLL45ld218RhkrfZCVvbC/W0YAe45e", "12341"))
}
