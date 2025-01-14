package query

import (
	"net/url"
	"reflect"
)

const queryStructTag = "query"
const pathParamStructTag = "pathparam"

func MarshalWithSettings(value interface{}, settings QuerySettings) url.Values {
	e := encoder{settings}
	kv := url.Values{}
	val := reflect.ValueOf(value)
	if !val.IsValid() {
		return nil
	}
	typ := val.Type()
	for _, pair := range e.typeEncoder(typ)("", val) {
		kv.Add(pair.key, pair.value)
	}
	return kv
}

func Marshal(value interface{}) url.Values {
	return MarshalWithSettings(value, QuerySettings{})
}

type Queryer interface {
	URLQuery() url.Values
}

type QuerySettings struct {
	NestedFormat NestedQueryFormat
	ArrayFormat  ArrayQueryFormat
}

type NestedQueryFormat int

const (
	NestedQueryFormatBrackets NestedQueryFormat = iota
	NestedQueryFormatDots
)

type ArrayQueryFormat int

const (
	ArrayQueryFormatComma ArrayQueryFormat = iota
	ArrayQueryFormatRepeat
	ArrayQueryFormatIndices
	ArrayQueryFormatBrackets
)
