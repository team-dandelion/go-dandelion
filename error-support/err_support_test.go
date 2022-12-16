package error_support

import (
	"fmt"
	"gitlab.outer.staruniongame.com/platform/toolbox/proto/gm_manager"
	"testing"
)

func TestFormat(t *testing.T) {
	Init(".")
	out := &gm_manager.CommonResp{}
	Format(&Error{Code: 1001}, out)
	fmt.Println(out)
}
