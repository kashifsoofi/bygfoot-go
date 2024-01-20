package domain

// A struct representing an option or a constant.
type Option struct {
	Name        string
	StringValue string
	Value       int
}

// An array of name-value pairs with an associated keyed data list for quick access.
type OptionList struct {
	list     []Option
	dataList map[string]Option
}

func (ol *OptionList) Length() int {
	return len(ol.list)
}

func (ol *OptionList) ListItem(i int) Option {
	return ol.list[i]
}

func (ol *OptionList) Add(opt Option) {
	ol.list = append(ol.list, opt)
}
