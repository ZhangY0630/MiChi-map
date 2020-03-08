package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	filterParams string = "sites:global AND status:Published"
)

// RequestData parses the request body to JSON
type requestData struct {
	Request []requestDataItem `json:"requests"`
}
type requestDataItem struct {
	IndexName string `json:"indexName"`
	Params    string `json:"params"`
}

// RequestParams stores the parameters for request
type RequestParams struct {
	NumbersPerPage int      `name:"hitsPerPage"`
	Attributes     []string `name:"attributesToRetrieve"`
	Page           int      `name:"page"`
}

// JSON parses the request data to JSON
func (rp RequestParams) JSON(in string) []byte {
	rd := requestData{}
	rd.Request = make([]requestDataItem, 1)
	rd.Request[0].IndexName = in
	rd.Request[0].Params = rp.Encode()
	json, _ := json.Marshal(rd)
	return json
}

// Encode encodes the struct to URL format (partly)
func (rp RequestParams) Encode() string {
	ts := reflect.TypeOf(rp)
	vs := reflect.ValueOf(rp)
	nf := ts.NumField()
	params := make([]string, nf)
	for i := 0; i < nf; i++ {
		t := ts.Field(i)
		v := vs.Field(i)
		params[i] = fieldToString(v, t.Tag.Get("name"))
	}
	return strings.Join(params, "&")
}

func fieldToString(v reflect.Value, tag string) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%s=%d", tag, v.Int())
	case reflect.Slice:
		items := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			items[i] = fmt.Sprint("%22", item.String(), "%22")
		}

		s := fmt.Sprint("%5B", strings.Join(items, "%2C"), "%5D")
		return fmt.Sprintf("%s=%s", tag, s)
	}
	return ""
}
