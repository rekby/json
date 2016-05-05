package json

import (
	"log"
	"strconv"
	"strings"
	"encoding/json"
)

type Json struct {
	val interface{}
	m   map[string]*Json
}

func New() *Json {
	return &Json{}
}

func FromJson(js interface{}) *Json {
	res := New()
	switch v := js.(type) {
	case bool, float64, string, nil:
		res.Setv(v)
	case []interface{}:
		arr := make([]*Json, len(v))
		for i := range arr {
			j := FromJson(v[i])
			arr[i] = j
		}
		res.Setv(arr)
	case map[string]interface{}:
		res.m = make(map[string]*Json)
		for key, val := range v {
			j := FromJson(val)
			res.m[key] = j
		}
	default:
		log.Printf("Json FromJson unknown type '%v'", v)
	}
	return res
}

func (this *Json) ToJson() interface{}{
	if this == nil {
		log.Println("Json nil ToJson")
		return nil
	}
	if this.m != nil {
		res := make(map[string]interface{})
		for k, v := range this.m {
			res[k] = v.ToJson()
		}
		return res
	}
	if this.val == nil {
		return nil
	}

	switch v := this.val.(type) {
	case bool,float64,string:
		return v
	case []*Json:
		arr := make([]interface{}, len(v))
		for i := range v {
			arr[i] = v[i].ToJson()
		}
		return arr
	default:
		log.Printf("Json unknow type ToJson '%v'\n", v)
		return nil
	}
}

func (this *Json) Equal(other *Json)bool {
	if this == nil && other == nil {
		return true
	}
	if this == nil || other == nil {
		return false
	}

	if this.m == nil && this.val == nil && other.m == nil && other.val == nil {
		return true
	}

	if this.m != nil && other.m != nil {
		if len(this.m) != len(other.m){
			return false
		}
		for k, v := range this.m{
			if otherv, ok := other.m[k]; ok {
				if !v.Equal(otherv) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}

	if this.val != nil && other.val != nil {
		switch v := this.val.(type) {
		case bool:
			if ov, ok := other.val.(bool); ok && v == ov {
				return true
			} else {
				return false
			}
		case float64:
			if ov, ok := other.val.(float64); ok && v == ov {
				return true
			} else {
				return false
			}
		case string:
			if ov, ok := other.val.(string); ok && v == ov {
				return true
			} else {
				return false
			}
		case []*Json:
			if ov, ok := other.val.([]*Json); ok {
				if len(v) != len(ov){
					return false
				}
				for i := range v {
					if !v[i].Equal(ov[i]) {
						return false
					}
				}
				return true
 			} else {
				return false
			}
		default:
			log.Printf("Json compare unknow type '%v'\n", v)
			return false
		}
	}

	return false
}

func (this *Json) Marshal() []byte {
	j := this.ToJson()
	res, err := json.Marshal(j)
	if err != nil {
		log.Println("Json err while Marshal:", err)
		return nil
	}
	return res
}

func (this *Json) MarshalIdent(prefix, ident string) []byte {
	j := this.ToJson()
	res, err := json.MarshalIndent(j, prefix, ident)
	if err != nil {
		log.Println("Json err while MarshalIdent:", err)
		return nil
	}
	return res
}

func Unmarshal(data []byte) (*Json, error){
	var js interface{}
	err := json.Unmarshal(data, &js)
	if err != nil {
		return nil, err
	}
	return FromJson(js), nil
}

func (this *Json) A(name string) []*Json {
	return this.J(name).Av()
}

func (this *Json) Av() []*Json {
	if this == nil || this.val == nil {
		log.Println("Json cast nil to array")
		return []*Json{}
	}
	if v, ok := this.val.([]*Json); ok && v != nil {
		return v
	}
	return []*Json{}
}

func (this *Json) B(name string) bool {
	return this.J(name).Bv()
}
func (this *Json) Bv() bool {
	if this == nil {
		log.Println("Json cast nil to bool")
		return false
	}
	if this.m != nil {
		return len(this.m) > 0
	}
	switch v := this.val.(type) {
	case bool:
		return v
	case float64:
		return !(-0.5 < v || v < 0.5)
	case string:
		val := strings.TrimSpace(strings.ToLower(v))
		for _, falseString := range []string{"", "false", "0"} {
			if val == falseString {
				return false
			}
		}
		return true
	case []*Json:
		return len(v) > 0
	case nil:
		return false
	default:
		log.Printf("Json can't cast to bool '%v'", v)
		return false
	}
}

func (this *Json) F(name string) float64 {
	return this.J(name).Fv()
}
func (this *Json) Fv() float64 {
	if this == nil {
		log.Println("Json cast nil to float")
		return 0
	}
	switch v := this.val.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case float64:
		return v
	case string:
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Printf("Json can't strconv to float '%v'\n", val)
			return 0
		}
		return val
	case nil:
		log.Println("Json cast nil val to float")
		return 0
	default:
		log.Printf("Json can't cast to float '%v'\n", v)
		return 0
	}
}

func (this *Json) I(name string) int {
	return this.J(name).Iv()
}

func (this *Json) Iv() int {
	if this == nil {
		log.Println("Json cast nil to int")
		return 0
	}
	switch v := this.val.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case float64:
		return int(v)
	case string:
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Json can't strconv to int: '%v'\n", v)
			return 0
		}
		return val
	case nil:
		log.Println("Json cast nil val to int")
		return 0
	default:
		log.Printf("Json can't cast to int: '%v'\n", v)
		return 0
	}
}

func (this *Json) IsNull(name string) bool {
	return this.J(name).IsNullv()
}

func (this *Json) IsNullv() bool {
	if this == nil {
		log.Println("Json isnull on nil object")
		return false
	}
	return this.val == nil && this.m == nil
}

func (this *Json) J(name string) *Json {
	if this == nil {
		log.Println("Json cast nil to json")
		return &Json{}
	}
	if this.m == nil {
		this.m = make(map[string]*Json)
	}
	if v, ok := this.m[name]; ok {
		return v
	} else {
		v = New()
		this.m[name] = v
		return v
	}
}

func (this *Json) Jv() *Json {
	if this == nil {
		return &Json{}
	}
	return &Json{val: this.val}
}

func (this *Json) S(name string) string {
	return this.J(name).Sv()
}

func (this *Json) Sv() string {
	if this == nil {
		log.Println("Json cast nil to string")
		return ""
	}
	switch v := this.val.(type) {
	case bool:
		if v {
			return "true"
		} else {
			return "false"
		}
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case nil:
		log.Println("Json cast nil val to string")
		return ""
	default:
		log.Printf("Json can't cast to string: '%v'\n", v)
		return ""
	}
}

func (this *Json) Set(name string, val interface{}) *Json {
	if this == nil {
		log.Println("Json setname on nil")
		return nil
	}
	this.val = nil
	if this.m == nil {
		this.m = make(map[string]*Json)
	}
	j := New()
	this.m[name] = j
	j.Setv(val)
	return this
}

func (this *Json) Setv(val interface{}) {
	if this == nil {
		log.Println("Json set on nil val")
		return
	}
	switch v := val.(type) {
	// Numbers to float as json parser
	case nil:
		this.val = nil
	case float64:
		this.val = v
	case float32:
		this.val = float64(v)
	case int:
		vf := float64(v)
		this.val = vf
	case int64:
		vf := float64(v)
		this.val = vf
	case uint:
		vf := float64(v)
		this.val = vf
	case int32:
		vf := float64(v)
		this.val = vf
	case uint32:
		vf := float64(v)
		this.val = vf
	// Non numeric
	case bool:
		this.val = v
	case string:
		this.val = v
	case []*Json:
		this.val = v
	case *Json:
		if len(v.m) > 0 {
			this.val = nil
			this.m = v.m
		} else {
			this.val = v.val
			this.m = nil
		}
	default:
		log.Printf("Json set unknown '%v'\n", v)
	}
}
