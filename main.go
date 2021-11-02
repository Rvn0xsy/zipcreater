package main

import (
	"fmt"
)

func main() {
	zipFlag := ZipFileStruct{}
	SetFlag(&zipFlag)
	if zipFlag.DestFile == "" || zipFlag.Filename == "" {
		ShowBanner()
		return
	}
	if zipFlag.IsCreate {
		fmt.Println("[+]Create Zip File....")
		CreateZipFile(zipFlag.DestFile, zipFlag.SourceFile, zipFlag.Path)
		return
	}
	fmt.Println(fmt.Sprintf("[+]Replace %s to %s ...\n", zipFlag.Filename, zipFlag.Path))
	Zip(zipFlag.SourceFile, zipFlag.DestFile, zipFlag.Path, zipFlag.Filename)
}
