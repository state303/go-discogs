package cmd

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
)

func fixTypesAsPlural(k *koanf.Koanf) {
	types := k.Strings("types")
	typesYaml := "types:\n"
	for i := range types {
		size := len(types[i])
		if types[i][size-1:] != "s" {
			types[i] = types[i] + "s"
		}
		typesYaml += fmt.Sprintf("  - %+v\n", types[i])
	}
	_ = k.Load(rawbytes.Provider([]byte(typesYaml)), yaml.Parser())
}
