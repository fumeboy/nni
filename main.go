package nni

import (
	"reflect"
)

type ParseOutput struct {
	url    string
	method string
}

func Parse(h handler) (o ParseOutput) {
	t := reflect.TypeOf(h).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Implements(reflect.TypeOf(new(Method)).Elem()){
			o.method = reflect.New(f.Type).Interface().(Method).withMethod()
		}
	}
	o.url = ParseURL(t)
	return
}

func ParseURL(t reflect.Type) (s string) {
	if t.NumField() == 0{
		return "/"
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Implements(reflect.TypeOf(new(URL)).Elem()){
			s += ParseURL(f.Type) + f.Tag.Get("url")  + "/"
		}else if f.Type.Name() == "string"{
			s += ":"+f.Name  + "/"
		}
	}
	return
}
