package my_test

import (
	"fmt"
	"testing"

	"github.com/dolthub/go-mysql-server/memory"
	_ "vitess.io/vitess/go/vt/vterrors"
)

func Test_Collision(t *testing.T) {
	fmt.Println(memory.NewDatabase("mydb"))
}
