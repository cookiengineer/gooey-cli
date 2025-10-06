package files

import "embed"
import "io/fs"

//go:embed boilerplate/**
var embedded_filesystem embed.FS

var Boilerplate fs.FS = nil

func init() {

	tmp, err := fs.Sub(embedded_filesystem, "boilerplate")

	if err == nil {
		Boilerplate= tmp
	}

}
