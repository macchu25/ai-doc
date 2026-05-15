package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			continue
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func ScanCodebase(root string) ([]string, []string) {
	var structure []string
	var frameworks []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden and node_modules
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") || info.Name() == "node_modules" || info.Name() == "__pycache__" {
				return filepath.SkipDir
			}
		}

		relPath, _ := filepath.Rel(root, path)
		if info.IsDir() {
			structure = append(structure, relPath+"/")
		} else {
			structure = append(structure, "  "+relPath)
			
			// Detect frameworks
			name := info.Name()
			if name == "package.json" {
				frameworks = append(frameworks, "Node.js")
			} else if name == "pom.xml" {
				frameworks = append(frameworks, "Java/Maven")
			} else if name == "go.mod" {
				frameworks = append(frameworks, "Go")
			}
		}

		return nil
	})

	return structure, frameworks
}
