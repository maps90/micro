package commons

import (
    "fmt"
    "reflect"
)

//Convent a Struct or Multilevel Struct to Map
func JsMap(in interface{}) (map[string]interface{}, error) {
    out := make(map[string]interface{})

    v := reflect.ValueOf(in)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("JsMap only accepts structs; got %T", v)
    }

    for i := 0; i < v.NumField(); i++ {
        // gets us a StructField
        fi := v.Type().Field(i)
        valueField := v.Field(i)

        tagv := fi.Tag.Get("jsmap")
        //tagv = strings.ToLower(fi.Name)

        if tagv != "" {
            // set key of map to value in struct field
            if valueField.Kind() == reflect.Slice {
                var xx []interface{}

                s := reflect.ValueOf(valueField.Interface())

                for i := 0; i < s.Len(); i++ {
                    fmt.Println(s.Index(i).Interface())
                    x, err := JsMap(s.Index(i).Interface())
                    if err != nil {
                        fmt.Println("error bray")
                    }
                    xx = append(xx, x)
                }

                fmt.Println(xx)

                out[tagv] = xx
            } else if valueField.Kind() == reflect.Struct {
                x, err := JsMap(valueField.Interface())
                if err != nil {
                    fmt.Println("error bray")
                }
                out[tagv] = x
            } else {
                out[tagv] = v.Field(i).Interface()
            }
        }
    }
    return out, nil
}
