// Package collector provides a way to collect objects from a reader, file or directory.
package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

// Collector is a struct that can collect objects from a reader, file or directory
type Collector[T any] struct {
	Decoder  func(io.Reader) ([]T, error)
	Vaidlate func(T) error
	Timout   time.Duration
}

var (
	// ErrSkip is an error to skip the current obj.
	ErrSkip        = errors.New("skip")
	defaultTimeout = 30 * time.Second
)

// NewCollector returns a new collector
// if decoder is nil it will use the default JSON decoder.
// if validate is nil it will use a no-op validate function.
func NewCollector[T any](decoder func(io.Reader) ([]T, error), vaidlate func(T) error) *Collector[T] {
	if vaidlate == nil {
		vaidlate = func(T) error { return nil }
	}
	if decoder == nil {
		var newObj T
		objType := getType(newObj)
		decoder = func(reader io.Reader) ([]T, error) {
			newObjPtr := reflect.New(objType).Interface().(T)
			err := json.NewDecoder(reader).Decode(newObjPtr)
			if err != nil {
				return nil, err
			}
			return []T{newObjPtr}, nil
		}
	}
	return &Collector[T]{
		Decoder:  decoder,
		Vaidlate: vaidlate,
	}
}

func getType(obj any) reflect.Type {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Pointer {
		objType = objType.Elem()
	}
	return objType
}

// FromDir takes in a directory path and returns a list of objects.
// If recurse is true it will recurse into the directory.
func (c *Collector[T]) FromDir(dirName string, recurse bool) ([]T, error) {
	retList := []T{}
	// read all entries in the directory
	dirFiles, err := os.ReadDir(dirName)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir '%s': %w", dirName, err)
	}

	for _, dirEntry := range dirFiles {
		fullPath := filepath.Join(dirName, dirEntry.Name())
		if !dirEntry.IsDir() {
			fileObjs, err := c.FromFile(fullPath)
			if err != nil {
				return nil, err
			}
			retList = append(retList, fileObjs...)
			continue
		}
		if !recurse {
			continue
		}
		// if the entry is the dir recurse into that folder to get all crds
		dirObjs, err := c.FromDir(fullPath, recurse)
		if err != nil {
			return nil, err
		}
		retList = append(retList, dirObjs...)
	}
	return retList, nil
}

// FromFile takes in a file path and returns a list of objects.
func (c *Collector[T]) FromFile(fullPath string) ([]T, error) {
	// read the file and convert it to a crd object
	file, err := os.Open(filepath.Clean(fullPath))
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", fullPath, err)
	}

	//nolint: errcheck // we don't care about the close error
	defer file.Close()
	retList, err := c.FromReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to convert file '%s': %w", fullPath, err)
	}
	return retList, nil
}

// FromReader takes in a reader and returns a list objects.
func (c *Collector[T]) FromReader(reader io.Reader) ([]T, error) {
	objs, err := c.Decoder(reader)
	if err != nil {
		return nil, fmt.Errorf("failed decode: %w", err)
	}
	for _, obj := range objs {
		err := c.Vaidlate(obj)
		if err != nil {
			if errors.Is(err, ErrSkip) {
				continue
			}
			return nil, fmt.Errorf("failed to validate: %w", err)
		}
	}
	return objs, nil
}

// FromURL takes in a URL and returns a list of objects.
func (c *Collector[T]) FromURL(inputURL string) ([]T, error) {
	if c.Timout == 0 {
		c.Timout = defaultTimeout
	}
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get request YAML: %w", err)
	}

	//nolint: errcheck // we don't care about the close error
	defer resp.Body.Close()
	retList, err := c.FromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}
	return retList, nil
}
