package clone

import "github.com/jinzhu/copier"

// 结构体属性拷贝工具

// copier 拷贝策略.
var defaultOpt = copier.Option{
	// 忽略零值
	IgnoreEmpty: true,
	// 深拷贝
	DeepCopy: true,
	// 自定义类型转换器
	Converters: nil,
}

// Copy 对象属性深拷贝.
func Copy(dst, src any, option ...copier.Option) error {
	if len(option) > 0 {
		return copier.CopyWithOption(dst, src, option[0])
	}
	return copier.CopyWithOption(dst, src, defaultOpt)
}

// Clone 根据目标类型进行拷贝.
func Clone[T any](src any, option ...copier.Option) (*T, error) {
	var dst T
	var err error
	if len(option) > 0 {
		err = copier.CopyWithOption(&dst, src, option[0])
	} else {
		err = copier.CopyWithOption(&dst, src, defaultOpt)
	}
	if err != nil {
		return nil, err
	}
	return &dst, nil
}
