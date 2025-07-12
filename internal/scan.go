package internal

import (
	"io/fs"
	"path/filepath"
)
func Scan(root string) ([]string,error){
	var files []string

	err := filepath.WalkDir(root,func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if !d.IsDir(){
			files = append(files, path)
		}
		return nil
	})

	return files, err
}