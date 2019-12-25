package main
import (
    "fmt"
    "reflect"
)
type order struct {
    ordId    int
    ordName  string
    place    string
}
type employee struct {
    emid   int
    emname string
    branch string
    salary int
}
func createQuery(q interface{}) {
    if reflect.ValueOf(q).Kind() == reflect.Struct {
        t := reflect.ValueOf(q)
        nm := reflect.ValueOf(q).NumField()
        query := fmt.Sprintf("value is %v", t)
        for i := 0;i<nm;i++ {
            switch t.Field(i).Kind(){
                case reflect.Int:
                    query = fmt.Sprintf("%s,%d",query,t.Field(i))
                case reflect.String:
                    query = fmt.Sprintf("%s,%s",query,t.Field(i))
            }
        }
        fmt.Println(query)
        return
    }
    fmt.Println("Unsupported type")
}
func main() {
    o := order{
        ordId:   45,
        ordName: "packet",
        place: "jaipur",
    }
    e := employee{
        emid: 5,
        emname: "yash",
        branch: "cs",
        salary: 9,
    }
    createQuery(o)
    createQuery(e)
    i := 9
    createQuery(i)
}
