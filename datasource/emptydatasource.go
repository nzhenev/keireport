package datasource

import (
	"github.com/adlindo/keireport/core"
)

type EmptyDatasource struct {
}

func (o *EmptyDatasource) SetConfig(data map[string]interface{}) error {

	return nil
}

func (o *EmptyDatasource) Next(rpt *core.Keireport) (map[string]interface{}, error) {

	return nil, core.ErrEndOfRow
}

//------------------------------------------------

type EmptyDatasourceBuilder struct {
}

func (o *EmptyDatasourceBuilder) Build(data map[string]interface{}) (core.DataSource, error) {

	ret := &EmptyDatasource{}

	return ret, nil
}

//------------------------------------------------

func init() {

	core.RegisterDatasource("empty", &EmptyDatasourceBuilder{})
}
