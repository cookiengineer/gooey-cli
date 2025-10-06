package fs

import "os"

func IsFile(path string) bool {

	stat, err0 := os.Stat(path)

	if err0 == nil && stat.IsDir() == false {
		return true
	}

	return false

}
