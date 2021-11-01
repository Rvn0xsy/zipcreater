package main

import (
	"flag"
	"os"
)

// ZipFileStruct ...
type ZipFileStruct struct {
	SourceFile string
	DestFile   string
	Path       string
	IsCreate   bool
}

// SetFlag ...
func SetFlag(ZipFileFlag *ZipFileStruct) {
	flag.BoolVar(&ZipFileFlag.IsCreate, "create", false, "Create Zip File.")
	flag.StringVar(&ZipFileFlag.DestFile, "dest", "output.zip", "Output Zip File.")
	flag.StringVar(&ZipFileFlag.SourceFile, "source", "", "Source Zip File.")
	flag.StringVar(&ZipFileFlag.Path, "path", "", "File Path. e.g. ..//")
	if len(os.Args) < 2 {
		ShowBanner()
		os.Args = append(os.Args, "-h")
	}
	flag.Parse()
}
