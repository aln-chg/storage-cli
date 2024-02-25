package main

import (
	"flag"
	"fmt"
	"syscall"
)

var version = "development"

func main() {
	var stat syscall.Statfs_t
	var versionFlag = flag.Bool("version", false, "Print the version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", version)
		return
	}

	fs := "/" // Root file system. Change as needed.

	err := syscall.Statfs(fs, &stat)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	all := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := all - free

	fmt.Printf("Total: %.2f GB\n", float64(all)/1e9)
	fmt.Printf("Used: %.2f GB\n", float64(used)/1e9)
	fmt.Printf("Free: %.2f GB\n", float64(free)/1e9)
}
