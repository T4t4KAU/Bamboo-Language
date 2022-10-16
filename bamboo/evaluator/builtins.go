package evaluator

import (
	"fmt"
	"monkey/object"
	"os"
)

// 建立内置函数映射表
var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect(), " ")
			}
			fmt.Println()
			return NULL
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1 or 0", len(args))
			} else if len(args) == 1 {
				switch arg := args[0].(type) {
				case *object.Integer:
					os.Exit(int(arg.Value))
				default:
					return newError("argument to `exit` not supported, got %s", args[0].Type())
				}

			} else {
				os.Exit(0)
			}
			return NULL
		},
	},
}
