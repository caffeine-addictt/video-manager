package utils

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// Path Flag
type PathFlag string

func (e *PathFlag) String() string {
	return string(*e)
}

func (e *PathFlag) Set(value string) error {
	if _, err := os.Stat(filepath.Clean(value)); err != nil {
		return errors.New("must be a valid path")
	}

	*e = PathFlag(value)
	return nil
}

func (e *PathFlag) Type() string {
	return "<path>"
}

// File Path Flag
type FilePathFlag string

func (e *FilePathFlag) String() string {
	return string(*e)
}

func (e *FilePathFlag) Set(value string) error {
	filePath, err := ValidateFile(value)
	if err != nil {
		return err
	}

	*e = FilePathFlag(filePath)
	return nil
}

func (e *FilePathFlag) Type() string {
	return "<path>"
}

// Directory Path Flag
type DirPathFlag string

func (e *DirPathFlag) String() string {
	return string(*e)
}

func (e *DirPathFlag) Set(value string) error {
	dirPath, err := ValidateDirectory(value)
	if err != nil {
		return err
	}

	*e = DirPathFlag(dirPath)
	return nil
}

func (e *DirPathFlag) Type() string {
	return "<path>"
}

// Functions
func ValidateFile(path string) (string, error) {
	filePath, err := homedir.Expand(path)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return "", errors.Wrapf(err, "File does not exist: %v\n", filePath)
	}
	if err != nil {
		return "", errors.Wrapf(err, "File error: %v\n", filePath)
	}

	if info.IsDir() {
		return "", errors.Errorf("%#v is a directory\n", filePath)
	}

	// Check for read/write permissions
	mode := info.Mode()
	if mode.Perm()&0o666 == 0 {
		return "", errors.Errorf("%#v has no read/write permissions\n", filePath)
	}

	return filePath, nil
}

func ValidateDirectory(path string) (string, error) {
	dirPath, err := homedir.Expand(path)
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

func ReadDirectory(dir string, n int) ([]string, error) {
	dirPath, err := homedir.Expand(dir)
	if err != nil {
		return []string{}, errors.WithStack(err)
	}

	file, err := os.Open(filepath.Clean(dirPath))
	if err != nil {
		return []string{}, errors.WithStack(err)
	}
	defer file.Close()

	names, err := file.Readdirnames(n)
	if err != nil {
		return nil, err
	}

	return names, nil
}
