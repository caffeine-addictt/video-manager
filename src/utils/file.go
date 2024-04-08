package utils

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

func ValidateDirectory(dir string) (string, error) {
	dirPath, err := homedir.Expand(dir)
	if err != nil {
		return "", errors.WithStack(err)
	}

	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return "", errors.Wrapf(err, "Directory does not exist: %v\n", dirPath)
	}
	if err != nil {
		return "", errors.Wrapf(err, "Directory error: %v\n", dirPath)
	}

	if !info.IsDir() {
		return "", errors.Errorf("%#v is not a directory\n", dirPath)
	}

	// Check for read/write permissions
	mode := info.Mode()
	if mode.Perm()&0o666 == 0 {
		return "", errors.Errorf("%#v has no read/write permissions\n", dirPath)
	}

	return dirPath, nil
}
