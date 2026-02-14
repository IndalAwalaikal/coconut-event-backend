package util

import (
	"bytes"
	"encoding/csv"
)

// CSVFromRows takes a 2D slice of strings and returns CSV bytes.
func CSVFromRows(rows [][]string) ([]byte, error) {
    buf := &bytes.Buffer{}
    w := csv.NewWriter(buf)
    for _, r := range rows {
        if err := w.Write(r); err != nil {
            return nil, err
        }
    }
    w.Flush()
    if err := w.Error(); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
