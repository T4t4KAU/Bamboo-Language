package object

// Environment 将值与名称关联
// 使用关联的名称跟踪值
// 本质上 环境是一个将字符串与对象相关联的哈希映射
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnclosedEnvironment 扩展已有环境
// 将函数的参数添加到一个新的环境中
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment 创建环境
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Get 从环境中获取变量的值
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set 将名称与值关联
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
