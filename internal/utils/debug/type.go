package debug

import "reflect"

func TypeOf(v any) string {
	return reflect.TypeOf(v).String()
}
