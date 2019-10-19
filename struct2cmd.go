package struct2cmd

import (
    "strings"
    "flag"
    "reflect"
    "os"
)


func extractMethod(Ptr interface{}) []string {
    ref := reflect.TypeOf(Ptr).Elem()
    methodNames := make([]string, 0)
    for i:=0; i < ref.NumMethod(); i++ {
        method := ref.Method(i)
        methodNames = append(methodNames, method.Name)
    }
    return methodNames
}

func isBasicKind(k reflect.Kind) bool {
    switch (k) {
    case reflect.Bool:
        return true
    case reflect.String:
        return true
    case reflect.Int:
        return true
    }
    return false
}

func extractField(Ptr interface{}) map[string]reflect.Type{
    ref := reflect.TypeOf(Ptr).Elem()
    fieldName2Type := make(map[string]reflect.Type)
    for i:=0; i < ref.NumField(); i++ {
        field := ref.Field(i)
        if isBasicKind(field.Type.Kind()) {
            fieldName2Type[field.Name] = field.Type
        }
    }
    return fieldName2Type
}

func field2Flag(fields map[string]reflect.Type) map[string]interface{} {
    args := make(map[string]interface{})
    for field, Type := range fields {
        argName := strings.ToLower(field)
        switch Type.Kind() {
        case reflect.String:
            var arg string
            flag.StringVar(&arg, argName, "", field)
            args[field] = &arg
        case reflect.Int:
            var arg int
            flag.IntVar(&arg, argName, 0, field)
            args[field] = &arg
        case reflect.Bool:
            var arg bool
            flag.BoolVar(&arg, argName, true, field)
            args[field]= &arg
        }
    }
    return args
}

func setFlag2Field(Ptr interface{}, Flag map[string]interface{}) {
    ref := reflect.ValueOf(Ptr)
    for field, valuePtr := range Flag {
        ref.Elem().FieldByName(field).Set(reflect.ValueOf(valuePtr).Elem())
    }
}

var method string

func setupMethod(Ptr interface{}) {
    methods := extractMethod(Ptr)
    usage := "method should in [" + strings.Join(methods, ",") + "]"
    flag.StringVar(&method, "method", "", usage)
}

func Run(Ptr interface{}) {
    fields := extractField(Ptr)
    args := field2Flag(fields)
    setupMethod(Ptr)
    flag.Parse()
    setFlag2Field(Ptr, args)
    ref := reflect.ValueOf(Ptr).Elem()
    M, ok := ref.Type().MethodByName(method)
    if !ok {
        flag.Usage()
        os.Exit(1)
    }
    M.Func.Call([]reflect.Value{ref})
}
