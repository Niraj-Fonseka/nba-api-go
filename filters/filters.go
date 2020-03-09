package filters

type Filter struct {
	UserFields map[string]string
	Filters    []string
}

func (f *Filter) NewFilter(filters ...string) *Filter {
	return &Filter{
		UserFields: make(map[string]string),
		Filters:    filters,
	}
}

func (f *Filter) SetFilter(filterName string, value string) {
	f.UserFields[filterName] = value
}

func (f *Filter) GetFilters() []string {
	return f.Filters
}
