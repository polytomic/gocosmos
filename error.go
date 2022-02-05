package gocosmos

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type OverLimitError struct {
	StatusCode int
	Properties []ResponseProperty
	msg        string
}

func NewOverLimitError(statusCode int, response []byte) OverLimitError {
	envelope := map[string]string{}
	json.Unmarshal(response, &envelope)
	return OverLimitError{
		StatusCode: statusCode,
		Properties: responseProperties(strings.NewReader(envelope["message"])),
	}
}

func (e OverLimitError) Error() string {
	if e.msg != "" {
		return e.msg
	}
	for _, prop := range e.Properties {
		if prop.Key == "Message" {
			// parse the message as JSON to get the error
			message := map[string][]string{}
			err := json.Unmarshal([]byte(prop.Value), &message)
			if err != nil {
				break
			}
			if len(message["Errors"]) > 0 {
				e.msg = message["Errors"][0]
				return e.msg
			}
		}
	}

	return ""
}

type ResponseProperty struct {
	Key   string
	Value string
}

func responseProperties(resp io.Reader) []ResponseProperty {
	result := []ResponseProperty{}

	scanner := bufio.NewScanner(resp)
	scanner.Split(scanResponse())

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "{") {
			// this property is just a JSON blob; it happens
			result = append(result, ResponseProperty{Value: scanner.Text()})
			continue
		}

		kv := strings.SplitN(scanner.Text(), ":", 2)
		prop := ResponseProperty{Key: kv[0]}
		if len(kv) > 1 {
			prop.Value = strings.TrimSpace(kv[1])
		}
		result = append(result, prop)
	}

	return result
}

func scanResponse() bufio.SplitFunc {
	objDepth := 0

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i, c := range data {
			if c == '{' {
				objDepth++
			}
			if c == '}' {
				objDepth--
			}
			if objDepth == 0 {
				switch c {
				case '\n', ',', ';':
					return i + 1, bytes.TrimSpace(data[0:i]), nil
				}
			}
		}

		return 0, nil, nil
	}
}
