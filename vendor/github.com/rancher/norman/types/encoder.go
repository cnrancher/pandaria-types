package types

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/ghodss/yaml"
	jsoniter "github.com/json-iterator/go"
)

var (
	commenter = regexp.MustCompile("(?m)^( *)zzz#\\((.*)\\)\\((.*)\\)([a-z]+.*):(.*)")
)

func JSONEncoder(writer io.Writer, v interface{}, power bool) error {
	if power {
		fastjson := jsoniter.ConfigCompatibleWithStandardLibrary
		return fastjson.NewEncoder(writer).Encode(v)
	}
	return json.NewEncoder(writer).Encode(v)
}

func JSONLinesEncoder(writer io.Writer, v interface{}, power bool) error {
	if collection, ok := v.(*GenericCollection); ok {
		encoder := json.NewEncoder(writer)

		// encode the top level object first
		err := encoder.Encode(collection.Collection)
		if err != nil {
			return err
		}

		// write collection objects 1 at a time
		for _, obj := range collection.Data {
			err = encoder.Encode(obj)
			if err != nil {
				return err
			}
		}
	} else {
		// if we receive a type that is not a collection fall back to standard json encoding
		if err := json.NewEncoder(writer).Encode(v); err != nil {
			return err
		}
	}

	// a blank newline at the end indicates the complete response was returned, if this is absent an error occurred in the middle of encoding
	_, err := writer.Write([]byte("\n"))
	return err
}

func YAMLEncoder(writer io.Writer, v interface{}, power bool) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	buf, err := yaml.JSONToYAML(data)
	if err != nil {
		return err
	}
	//buf = commenter.ReplaceAll(buf, []byte("${1}# ${2}type: ${3}\n${1}# ${4}:${5}"))
	buf = commenter.ReplaceAll(buf, []byte("${1}# ${4}:${5}"))
	_, err = writer.Write(buf)
	return err
}
