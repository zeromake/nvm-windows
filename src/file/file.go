package file

import (
	"archive/zip"
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Function courtesy http://stackoverflow.com/users/1129149/swtdrgn
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	var symlinks [][2]string
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			isSymlink := f.Mode()&os.ModeSymlink != 0
			if isSymlink {
				var symlinkBytes []byte
				symlinkBytes, err = io.ReadAll(rc)
				if err != nil {
					return err
				}
				symlink := string(symlinkBytes)
				symlinks = append(symlinks, [2]string{symlink, fpath})
			} else {
				f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}
				defer f.Close()

				_, err = io.Copy(f, rc)
				if err != nil {
					return err
				}
			}
		}
	}
	for _, symlink := range symlinks {
		err = os.Symlink(symlink[0], symlink[1])
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
