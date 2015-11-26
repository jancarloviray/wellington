package handlers

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/net/context"

	libsass "github.com/wellington/go-libsass"
)

func init() {
	libsass.RegisterSassFunc("font-url($path, $raw: false)", FontURL)
}

// FontURL builds a relative path to the requested font file from the built CSS.
func FontURL(ctx context.Context, usv libsass.SassValue) (*libsass.SassValue, error) {

	var (
		path, format string
		csv          libsass.SassValue
		raw          bool
	)
	comp, err := libsass.CompFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	libctx := comp.Context()
	err = libsass.Unmarshal(usv, &path, &raw)

	if err != nil {
		return nil, err
	}

	// Enter warning
	if libctx.FontDir == "." || libctx.FontDir == "" {
		s := "font-url: font path not set"
		return nil, errors.New(s)
	}

	rel, err := filepath.Rel(libctx.BuildDir, libctx.FontDir)
	if err != nil {
		return nil, err
	}

	if raw {
		format = "%s%s"
	} else {
		format = `url("%s%s")`
	}
	// TODO: FontDir() on compiler
	fontdir := libctx.FontDir
	abspath := filepath.Join(fontdir, path)
	qry, err := qs(comp.CacheBust(), abspath)
	if err != nil {
		return nil, err
	}
	csv, err = libsass.Marshal(fmt.Sprintf(format,
		filepath.ToSlash(filepath.Join(rel, path)),
		qry,
	))

	return &csv, err
}

func sumHash(f io.Reader) (string, error) {
	hdr := make([]byte, 50*1024)
	if _, err := f.Read(hdr); err != nil {
		return "", err
	}
	h := crc32.NewIEEE()
	_, err := h.Write(hdr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%x", h.Sum(nil)), nil
}

func modHash(info os.FileInfo) (string, error) {
	mod := info.ModTime()
	bs, err := mod.MarshalText()
	if err != nil {
		return "", err
	}
	ts := sha1.Sum(bs)
	return "?" + fmt.Sprintf("%x", ts[:4]), nil
}

func qs(method string, abs string) (string, error) {
	var qry string
	var err error
	switch method {
	case "timestamp":
		var fileinfo os.FileInfo
		fileinfo, err = os.Stat(abs)
		if err != nil {
			return "", err
		}
		qry, err = modHash(fileinfo)
	case "sum":
		var r io.Reader
		r, err := os.Open(abs)
		if err != nil {
			return "", err
		}
		qry, err = sumHash(r)
		if err != nil {
			return "", err
		}
	}
	return qry, err
}
