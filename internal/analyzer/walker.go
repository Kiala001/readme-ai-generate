package analyzer

import (
	"io/fs"
	"path/filepath"
)

func WalkerProject(root string) ([]string, error) {
	var goFiles []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".go" {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	return goFiles, err
}