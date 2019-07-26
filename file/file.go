package file

import (
	"io/ioutil"
	"os"
)

// exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func ListAllFilesInDir(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	return files, err
}

func ListAllFilesOnlyInDir(path string) ([]os.FileInfo, error) {
	rawFiles, err := ListAllFilesInDir(path)
	files := filter(rawFiles, func(info os.FileInfo) bool {
		return !info.IsDir()
	})
	return files, err
}
func ListAllDirOnlyInDir(path string) ([]os.FileInfo, error) {
	rawFiles, err := ListAllFilesInDir(path)
	files := filter(rawFiles, func(info os.FileInfo) bool {
		return info.IsDir()
	})
	return files, err
}

func mapp(vs []os.FileInfo, f func(os.FileInfo) os.FileInfo) []os.FileInfo {
	vsm := make([]os.FileInfo, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func filter(vs []os.FileInfo, f func(os.FileInfo) bool) []os.FileInfo {
	vsf := make([]os.FileInfo, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func any(vs []os.FileInfo, f func(os.FileInfo) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func all(vs []os.FileInfo, f func(os.FileInfo) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}
