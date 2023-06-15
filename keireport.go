package keireport

import (
	_ "github.com/adlindo/keireport/component"
	"github.com/adlindo/keireport/core"
	_ "github.com/adlindo/keireport/datasource"
	_ "github.com/adlindo/keireport/exporter"
)

func LoadFromFile(fileName string) (*core.Keireport, error) {

	ret := &core.Keireport{}
	err := ret.LoadFromFile(fileName)

	return ret, err
}

func LoadFromString(templateString, baseDir string) (*core.Keireport, error) {

	ret := &core.Keireport{}
	err := ret.LoadFromString(templateString, baseDir)

	return ret, err
}
