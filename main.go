package main

import (
	"flag"
	"fmt"
	"syscall"
)

func main() {
	var Version = "development"
	var stat syscall.Statfs_t

	versionFlag := flag.Bool("version", false, "prints the version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Verison:", Version)
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
