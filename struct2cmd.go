package struct2cmd

import (
	"flag"
	"os"
	"reflect"
	"strings"
)

type Arg struct {
	Type    reflect.Type
	Name    string
	Default interface{}
}

func extractMethod(Ptr interface{}) []string {
	ref := reflect.TypeOf(Ptr)
	methodNames := make([]string, 0)
	for i := 0; i < ref.NumMethod(); i++ {
		method := ref.Method(i)
		if method.Type.NumIn() > 1 {
			continue
		}
		methodNames = append(methodNames, method.Name)
	}
	return methodNames
}

func isBasicKind(k reflect.Kind) bool {
	switch k {
	case reflect.Bool:
		return true
	case reflect.String:
		return true
	case reflect.Int, reflect.Int64:
		return true
	case reflect.Float64:
		return true
	}
	return false
}

func extractField(Ptr interface{}) map[string]Arg {
	ref := reflect.TypeOf(Ptr).Elem()
	fieldName2Type := make(map[string]Arg)
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		if isBasicKind(field.Type.Kind()) && (field.Name[0] >= 'A' && field.Name[0] <= 'Z') {
			arg := Arg{}
			arg.Type = field.Type
			if def, ok := field.Tag.Lookup("default"); ok {
				arg.Name = def
			} else {
				arg.Name = strings.ToLower(field.Name)
			}
			fieldName2Type[field.Name] = arg
		}
	}
	return fieldName2Type
}

func setupFieldArgs(fields map[string]Arg) map[string]interface{} {
	args := make(map[string]interface{})
	for field, arg := range fields {
		switch arg.Type.Kind() {
		case reflect.String:
			var argvar string
			flag.StringVar(&argvar, arg.Name, "", field)
			args[field] = &argvar
		case reflect.Int:
			var argvar int
			flag.IntVar(&argvar, arg.Name, 0, field)
			args[field] = &argvar
		case reflect.Bool:
			var argvar bool
			flag.BoolVar(&argvar, arg.Name, true, field)
			args[field] = &argvar
		case reflect.Int64:
			var argvar int64
			flag.Int64Var(&argvar, arg.Name, 0, field)
			args[field] = &argvar
		case reflect.Float64:
			var argvar float64
			flag.Float64Var(&argvar, arg.Name, 0.0, field)
			args[field] = &argvar
		}
	}
	return args
}

func setField2Struct(Ptr interface{}, Flag map[string]interface{}) {
	ref := reflect.ValueOf(Ptr)
	for field, valuePtr := range Flag {
		ref.Elem().FieldByName(field).Set(reflect.ValueOf(valuePtr).Elem())
	}
}

var method string

func setupMethodArg(Ptr interface{}) {
	methods := extractMethod(Ptr)
	usage := "method should in [" + strings.Join(methods, ",") + "]"
	flag.StringVar(&method, "method", "", usage)
}

func Run(Ptr interface{}) {
	fields := extractField(Ptr)
	args := setupFieldArgs(fields)
	setupMethodArg(Ptr)
	flag.Parse()
	setField2Struct(Ptr, args)
	ref := reflect.ValueOf(Ptr)
	M, ok := ref.Type().MethodByName(method)
	if !ok {
		flag.Usage()
		os.Exit(1)
	}
	M.Func.Call([]reflect.Value{ref})
}
