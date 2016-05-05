Wrapper for unmarshaled json object for comfortable work.

```
import (
   "fmt"
   "github.com/rekby/json"
)

func main(){
    var js interface{}
    var j *Json

    json.Unmarshal([]byte(`[1,2,"asd",true,null, "null"]`), &js)
    j = FromJson(js)

    fmt.Println(len(j.Av())) // 6
    fmt.Println(j.Av()[3].Sv()) // String ""

    // -----
    json.Unmarshal([]byte(`{"asd":123,"vvv":"test", "null":123}`), &js)
    j = FromJson(js)
    fmt.Println(j.I("asd")) // Int 123
    fmt.Println(j.S("asd")) // String "123"
    fmt.Println(j.I("vvv")) // Int 0 + log warning about bad cast
}
```