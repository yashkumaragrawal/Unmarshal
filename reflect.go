package main
import (
    "encoding/json"
    "fmt"
    "reflect"
)

func main() {
    p := `{"product1":{"id":11,"product_id":1,"title":"One Plus 7 Pro","color":"Grey","ram":"4GB","storage":"64GB","price":"34000"},"product2":{"id":12,"product_id":1,"title":"One Plus 7 Pro","color":"Grey","ram":"8GB","storage":"64GB","price":"36000"},"product3":{"id":13,"product_id":1,"title":"One Plus 7 Pro","color":"Blue","ram":"4GB","storage":"64GB","price":"33000"},"product4":{"id":14,"product_id":1,"title":"One Plus 7 Pro","color":"Blue","ram":"8GB","storage":"64GB","price":"35000"}}`
    var data interface{}
    bytes := []byte(p)
    json.Unmarshal(bytes, &data)
    tt := reflect.ValueOf(data)
    if tt.Kind() == reflect.Map {
        for _, key := range tt.MapKeys() {
            strct := tt.MapIndex(key)
            if reflect.ValueOf(strct).Kind() == reflect.Struct {
                fmt.Println(reflect.ValueOf(strct))
                fmt.Println(reflect.TypeOf(strct))
                fmt.Println(strct)
            }
            fmt.Println()
        }
    }
}
