package ziputil

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/grokify/simplego/io/ioutilmore"
	"github.com/pkg/errors"
)

func ZipFilesRx(zipfile, dir string, rx *regexp.Regexp, removePaths bool) error {
	_, filepaths, err := ioutilmore.ReadDirMore(dir, rx, true, false)
	if err != nil {
		return err
	}
	return ZipFiles(zipfile, removePaths, filepaths)
}

// ZipFiles compresses one or many files into a single zip archive file.
func ZipFiles(zipfile string, removePaths bool, srcfiles []string) error {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	zfile, err := os.OpenFile(zipfile, flags, 0644)
	if err != nil {
		return fmt.Errorf("E_FAILED_TO_OPEN_FILE [%s]", err)
	}

	zipw := zip.NewWriter(zfile)

	for _, filename := range srcfiles {
		if err := AddFileToZip(zipw, filename, removePaths); err != nil {
			return closeFileAndZipOnError(
				zfile,
				zipw,
				errors.Wrap(err, fmt.Sprintf("Failed to add file %s to zip", filename)))
		}
	}
	err = zipw.Close()
	if err != nil {
		return closeFileOnError(zfile, err)
	}
	return zfile.Close()
}

func closeFileAndZipOnError(f *os.File, zipw *zip.Writer, err error) error {
	if zipw != nil {
		zipwErr := zipw.Close()
		if zipwErr != nil {
			err = errors.Wrap(err, zipwErr.Error())
		}
	}
	if f != nil {
		fErr := f.Close()
		if fErr != nil {
			err = errors.Wrap(err, fErr.Error())
		}
	}
	return err
}

func closeFileOnError(f *os.File, err error) error {
	if f != nil {
		fErr := f.Close()
		if fErr != nil {
			err = errors.Wrap(err, fErr.Error())
		}
	}
	return err
}

func AddFileToZip(zipWriter *zip.Writer, filename string, removePaths bool) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	if !removePaths {
		header.Name = filename
	}

	// See http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return closeFileOnError(fileToZip, err)
	}
	return fileToZip.Close()
}
