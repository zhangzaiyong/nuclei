package dsl

import (
	"github.com/projectdiscovery/nebula"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/helpers/deserialization"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/runtime"
)

type Options struct {
	Store *runtime.Store
}

func AddGlobalCustomHelpers(options *Options) error {
	_ = nebula.AddFunc("generate_java_gadget", func(args ...interface{}) (interface{}, error) {
		gadget := args[0].(string)
		cmd := args[1].(string)

		var encoding string
		if len(args) > 2 {
			encoding = args[2].(string)
		}
		data := deserialization.GenerateJavaGadget(gadget, cmd, encoding)
		return data, nil
	})

	_ = nebula.AddFunc("nuclei_vars_set", func(key string, value interface{}) {
		options.Store.Set(key, value)
	})

	_ = nebula.AddFunc("nuclei_vars_get", func(key string) (interface{}, error) {
		return options.Store.Get(key), nil
	})

	_ = nebula.AddFunc("nuclei_vars_del", func(key string) {
		options.Store.Del(key)
	})

	_ = nebula.AddFunc("nuclei_vars_len", func(args ...interface{}) int {
		return options.Store.Len()
	})

	_ = nebula.AddFunc("nuclei_vars_has", func(key string) bool {
		return options.Store.Has(key)
	})

	return nil
}