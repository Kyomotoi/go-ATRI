package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

func (w *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	w.Total += uint64(n)
	w.PrintProgress()
	return n, nil
}

func (w WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete ,,Ծ‸Ծ,,\r", humanize.Bytes(w.Total))
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".download")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	out.Chmod(0644)
	out.Close()
	if err = os.Rename(filepath+".download", filepath); err != nil {
		return err
	}
	return nil
}
