package common

import (
	"flag"
	"fmt"
)

func Flag(options *ApCOptions) {
	fmt.Print(Banner)
	flag.IntVar(&options.QueryPageCount, "P", 1, "查询页数，例如 -P 1")
	flag.IntVar(&options.QueryStartFrom, "S", 1, "查询起始页，例如 -S 2")
	flag.StringVar(&options.HostLocation, "L", "", "查询地址，默认全部，例如 -L 香港")
	flag.Parse()
}
