package nbdserv

import (
	"fmt"

	"github.com/reusee/e/v2"
)

var (
	me     = e.Default.WithStack().WithName("nbdserv")
	ce, he = e.New(me)

	pt = fmt.Printf
)
