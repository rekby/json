package json

import (
	"testing"
	"encoding/json"
)

func TestSimpleValuesJson(t *testing.T){
	var js interface{}
	var j *Json

	// 0
	json.Unmarshal([]byte(`0`), &js)
	j = FromJson(js)
	if j.Av() == nil || len(j.Av()) != 0 {
		t.Error()
	}
	if j.Bv() != false {
		t.Error()
	}
	if j.Fv() != 0 {
		t.Error()
	}
	if j.Iv() != 0 {
		t.Error()
	}
	if v, ok := j.Jv().val.(float64); !ok || v != 0 {
		t.Error()
	}
	if j.Jv().m != nil {
		t.Error()
	}
	if j.Sv() != "0" {
		t.Error()
	}

	// true
	json.Unmarshal([]byte(`true`), &js)
	j = FromJson(js)
	if j.Av() == nil || len(j.Av()) != 0 {
		t.Error()
	}
	if j.Bv() != true {
		t.Error()
	}
	if j.Fv() != 1 {
		t.Error()
	}
	if j.Iv() != 1 {
		t.Error()
	}
	if v, ok := j.Jv().val.(bool); !ok || v != true {
		t.Error()
	}
	if j.Jv().m != nil {
		t.Error()
	}
	if j.Sv() != "true"{
		t.Error()
	}

	// false
	json.Unmarshal([]byte(`false`), &js)
	j = FromJson(js)
	if j.Av() == nil || len(j.Av()) != 0 {
		t.Error()
	}
	if j.Bv() != false {
		t.Error()
	}
	if j.Fv() != 0 {
		t.Error()
	}
	if j.Iv() != 0 {
		t.Error()
	}
	if v, ok := j.Jv().val.(bool); !ok || v != false {
		t.Error()
	}
	if j.Jv().m != nil {
		t.Error()
	}
	if j.Sv() != "false"{
		t.Error()
	}

	// string
	json.Unmarshal([]byte(`"asd"`), &js)
	j = FromJson(js)
	if j.Av() == nil || len(j.Av()) != 0 {
		t.Error()
	}
	if j.Bv() != true {
		t.Error()
	}
	if j.Fv() != 0 {
		t.Error()
	}
	if j.Iv() != 0 {
		t.Error()
	}
	if v, ok := j.Jv().val.(string); !ok || v != "asd" {
		t.Error()
	}
	if j.Jv().m != nil {
		t.Error()
	}
	if j.Sv() != "asd"{
		t.Error()
	}

	// string
	json.Unmarshal([]byte(`null`), &js)
	j = FromJson(js)
	if j.Av() == nil || len(j.Av()) != 0 {
		t.Error()
	}
	if j.Bv() != false {
		t.Error()
	}
	if j.Fv() != 0 {
		t.Error()
	}
	if j.Iv() != 0 {
		t.Error()
	}
	if j.Jv().val != nil {
		t.Error()
	}
	if j.Jv().m != nil {
		t.Error()
	}
	if j.Sv() != ""{
		t.Error()
	}
}

func TestArrJson(t *testing.T){
	var js interface{}
	var j *Json

	json.Unmarshal([]byte(`[1,2,"asd",true,null, "null"]`), &js)
	j = FromJson(js)
	arr := j.Av()
	if len(arr) != 6 {
		t.Error()
	}
	if arr[0].Iv()!= 1 {
		t.Error()
	}
	if arr[1].Iv()!=2 {
		t.Error()
	}
	if arr[2].Sv() != "asd" {
		t.Error()
	}
	if arr[3].Bv() != true {
		t.Error()
	}
	if arr[4].Jv().val != nil {
		t.Error()
	}
	if arr[5].Sv() != "null" {
		t.Error()
	}
}

func TestObjJson(t *testing.T){
	var js interface{}
	var j *Json

	json.Unmarshal([]byte(`{"asd":123,"vvv":"test", "null":123}`), &js)
	j = FromJson(js)
	arr := j.Av()
	if len(arr) != 0 {
		t.Error()
	}
	if j.I("asd") != 123 {
		t.Error()
	}
	if j.S("vvv")!= "test"{
		t.Error()
	}
	if j.I("null") != 123 {
		t.Error()
	}
}