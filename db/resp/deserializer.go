package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type RESP struct {
	reader *bufio.Reader
}

func NewRESP(rd io.Reader) *RESP {
	return &RESP{
		reader: bufio.NewReader(rd),
	}
}

// readLine from a given command
// example: "$5\r\n
//			Ahmed\r\n"
// two lines, we read each line one at a time.
// read up to \r\n and remove the \r\n
func (r *RESP) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *RESP) Read() (Value, error) {
	Type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch Type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		return Value{}, fmt.Errorf("Unknown type: %v", Type)
	}
}

func (r *RESP) readArray() (Value, error) {
	v := Value{ Type: "array" }

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]Value, 0)
	for i := 0; i < len; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		v.Array = append(v.Array, val)
	}

	return v, nil
}

func (r *RESP) readBulk() (Value, error) {
	v := Value{ Type: "bulk" }

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, len)

	r.reader.Read(bulk)

	v.Bulk = string(bulk)

	// Read the trailing CRLF
	r.readLine()

	return v, nil
}

func (r *RESP) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}
