package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ImportInts a file
func ImportInts(path string) (ints []int, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read credential file")
	}

	buf := bytes.NewBuffer(data)
	for {
		strWithDelim, err := buf.ReadString([]byte("\n")[0])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.Wrap(err, "unable to readstring")
		}

		str := strings.TrimSuffix(strWithDelim, "\n")

		i64, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse data: '%s'", str)
		}

		if i64 > math.MaxInt32 {
			return nil, errors.New("int64 exceeds max int")
		}

		ints = append(ints, int(i64))
	}

	return ints, nil
}
