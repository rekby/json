package json

import (
	"testing"
)

func TestSimpleValuesJson(t *testing.T){
	var err error
	var j *Json

	// 0
	j, err = Parse([]byte(`0`))
	if err != nil {
		t.Error(err)
	}
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
	if j.IsNullv() == true {
		t.Error()
	}

	// true
	j, err = Parse([]byte(`true`))
	if err != nil {
		t.Error(err)
	}
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
	if j.IsNullv() == true {
		t.Error()
	}

	// false
	j, err = Parse([]byte(`false`))
	if err != nil {
		t.Error(err)
	}
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
	if j.IsNullv() == true {
		t.Error()
	}

	// string
	j, err = Parse([]byte(`"asd"`))
	if err != nil {
		t.Error(err)
	}
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
	if j.IsNullv() == true {
		t.Error()
	}

	// null
	j, err = Parse([]byte(`null`))
	if err != nil {
		t.Error(err)
	}
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
	if j.IsNullv() != true {
		t.Error()
	}
}

func TestArrJson(t *testing.T){
	j, err := Parse([]byte(`[1,2,"asd",true,null, "null"]`))
	if err != nil {
		t.Error(err)
	}
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
	j, err := Parse([]byte(`{"asd":123,"vvv":"test", "null":123}`))
	if err != nil {
		t.Error(err)
	}
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