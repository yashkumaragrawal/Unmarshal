package main
import (
    "encoding/json"
    "fmt"
    "reflect"
)
func main() {
    p := `{"product1":{"id":11,"product_id":1,"title":"One Plus 7 Pro","color":"Grey","ram":"4GB","storage":"64GB","price":"34000"},"product2":{"id":12,"product_id":1,"title":"One Plus 7 Pro","color":"Grey","ram":"8GB","storage":"64GB","price":"36000"},"product3":{"id":13,"product_id":1,"title":"One Plus 7 Pro","color":"Blue","ram":"4GB","storage":"64GB","price":"33000"},"product4":{"id":14,"product_id":1,"title":"One Plus 7 Pro","color":"Blue","ram":"8GB","storage":"64GB","price":"35000"}}`
    var data map[string](map[string]interface{})
    bytes := []byte(p)
    json.Unmarshal(bytes, &data)
    t := reflect.ValueOf(data)
    if t.Kind() == reflect.Map {
        fmt.Println(t.Kind())
    }
    for index,element := range data{
        fmt.Print(index," - ")
        fmt.Println(element["title"]," ",element["color"]," ",element["ram"]," ",element["storage"]," ",element["price"])
    }
}
