package vm

import (
	"github.com/multiverse-os/ruby/rubyvm/ast"

	. "github.com/multiverse-os/ruby/rubyvm/interpreter/vm/builtins"
)

func interpretArrayInContext(
	vm *vm,
	arrayNode ast.Array,
	context Value,
) (Value, error) {

	arrayValue, _ := vm.CurrentClasses["Array"].New(vm)
	array := arrayValue.(*Array)

	for _, node := range arrayNode.Nodes {
		value, err := vm.executeWithContext(context, node)
		if err != nil {
			return nil, err
		}

		array.Append(value)
	}

	return array, nil
}
