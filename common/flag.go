package common

import (
	"flag"
	"fmt"
)

func Flag(options *ApCOptions) {
	fmt.Print(Banner)
	flag.IntVar(&options.QueryPageCount, "P", 1, "查询页数")
	flag.Parse()
}
