package lr

type LangSpecs struct {
	Go *GoSpec `yaml:"go"`
}

type GoSpec struct {
	Package                   string            `yaml:"package"`
	Prefix                    string            `yaml:"prefix"`
	ValueTypes                map[string]string `yaml:"value_types"`
	OutputDebugNonKernelItems bool              `yaml:"output_debug_non_kernel_items"`
}

type Param struct {
    Name      string
    ParamType interface{}
}

type ParamList []*Param

func (list ParamList) Len() int {
    return len(list)
}

func (list ParamList) Less(i int, j int) bool {
    return list[i].Name < list[j].Name
}

func (list ParamList) Swap(i int, j int) {
    list[i], list[j] = list[j], list[i]
}
