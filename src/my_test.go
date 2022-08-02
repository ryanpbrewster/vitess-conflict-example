package my_test

import (
	"fmt"
	"testing"

	"github.com/dolthub/go-mysql-server/memory"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
)

func Test_Collision(t *testing.T) {
	fmt.Println(memory.NewDatabase("mydb"))
	fmt.Println(topodatapb.CellInfo{})
}
