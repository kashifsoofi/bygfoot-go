package models

// A struct representing an option or a constant.
type Option struct {
	name, string_value string
	value              int
}

// An array of name-value pairs with an associated keyed data list for quick access.
type OptionList struct {
	list     []Option
	dataList map[string]Option
}
