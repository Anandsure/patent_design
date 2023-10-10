package xml_file_extractor

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func extractAllFiles() error {
	directory := "datasets/I20230502/I20230502/DESIGN/"
	exportPath := "file_extraction/full_data/"

	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(directory, file.Name())
		zipReader, err := zip.OpenReader(filePath)
		if err != nil {
			return err
		}
		defer zipReader.Close()

		for _, zipFile := range zipReader.File {
			zipFileReader, err := zipFile.Open()
			if err != nil {
				return err
			}
			defer zipFileReader.Close()

			extractFilePath := filepath.Join(exportPath, zipFile.Name)

			if zipFile.FileInfo().IsDir() {
				os.MkdirAll(extractFilePath, os.ModePerm)
			} else {
				outputFile, err := os.Create(extractFilePath)
				if err != nil {
					return err
				}
				defer outputFile.Close()

				_, err = io.Copy(outputFile, zipFileReader)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func moveAllXMLFiles() error {
	exportPath := "file_extraction/all_xml/"
	directory := "file_extraction/full_data/"

	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		xmlFile := filepath.Join(directory, file.Name(), file.Name()+".XML")
		destPath := filepath.Join(exportPath, file.Name()+".XML")

		if err := os.Rename(xmlFile, destPath); err != nil {
			return err
		}
	}

	return nil
}

func _() {
	if err := extractAllFiles(); err != nil {
		fmt.Println("Error extracting files:", err)
		return
	}

	if err := moveAllXMLFiles(); err != nil {
		fmt.Println("Error moving XML files:", err)
		return
	}

	fmt.Println("XML files moved successfully.")
}
