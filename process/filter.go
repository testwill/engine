package process

import (
	"strings"

	"github.com/mesg-foundation/engine/protobuf/types"
)

// Match returns true if the data match the current list of filters
func (f Process_Node_Filter) Match(data *types.Struct) bool {
	for _, condition := range f.Conditions {
		if !condition.Match(data) {
			return false
		}
	}
	return true
}

// Match returns true the current filter matches the given data
func (f Process_Node_Filter_Condition) Match(data *types.Struct) bool {
	switch f.Predicate {
	case Process_Node_Filter_Condition_EQ:
		return data.Fields[f.Key].Equal(f.Value)
	case Process_Node_Filter_Condition_GT,
		Process_Node_Filter_Condition_GTE,
		Process_Node_Filter_Condition_LT,
		Process_Node_Filter_Condition_LTE:
		n1, ok1 := data.Fields[f.Key].Kind.(*types.Value_NumberValue)
		n2, ok2 := f.Value.Kind.(*types.Value_NumberValue)
		if !ok1 || !ok2 {
			return false
		}
		switch f.Predicate {
		case Process_Node_Filter_Condition_GT:
			return n1.NumberValue > n2.NumberValue
		case Process_Node_Filter_Condition_GTE:
			return n1.NumberValue >= n2.NumberValue
		case Process_Node_Filter_Condition_LT:
			return n1.NumberValue < n2.NumberValue
		case Process_Node_Filter_Condition_LTE:
			return n1.NumberValue <= n2.NumberValue
		}
	case Process_Node_Filter_Condition_CONTAINS:
		switch dataTyped := data.Fields[f.Key].Kind.(type) {
		case *types.Value_ListValue:
			for _, value := range dataTyped.ListValue.Values {
				if value.Equal(f.Value) {
					return true
				}
			}
			return false
		case *types.Value_StringValue:
			filter, ok := f.Value.Kind.(*types.Value_StringValue)
			if !ok {
				return false
			}
			return strings.Contains(dataTyped.StringValue, filter.StringValue)
		default:
			return false
		}
	}
	return false
}
