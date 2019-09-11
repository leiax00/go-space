package conf

import "go.uber.org/fx"

func NewConfModule(ymlFile *string) fx.Option {
	yml = ymlFile
	return fx.Provide(NewConfig)
}
