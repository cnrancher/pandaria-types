package schema

import (
	"github.com/rancher/norman/types"
	v3 "github.com/rancher/types/apis/mgt.pandaria.io/v3"
	"github.com/rancher/types/factory"
)

var (
	Version = types.APIVersion{
		Version:          "v3",
		Group:            "mgt.pandaria.io",
		Path:             "/v3/pandaria",
		SubContext:       true,
		SubContextSchema: "/v3/schemas/pandaria",
	}

	Schemas = factory.Schemas(&Version).
		Init(filterTypes)
)

func filterTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.Filter{}).
		MustImport(&Version, v3.SensitiveFilter{})
}
