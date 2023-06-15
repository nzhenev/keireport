package datasource

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/adlindo/keireport/core"
	"github.com/adlindo/keireport/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBDatasource struct {
	DBType  string
	ConnStr string
	Query   string
	RawConn *sql.DB
	Conn    *gorm.DB
	RowNo   int
	Rows    []map[string]interface{}
}

func (o *DBDatasource) init() error {

	var err error

	connBuilder, _ := dbConnBuilderMap[o.DBType]

	if connBuilder == nil {

		return errors.New("Connection builder is not found : " + o.DBType)
	} else {

		o.Conn, err = connBuilder.Build(o)
	}

	return err
}

func (o *DBDatasource) SetConfig(data map[string]interface{}) error {

	var err error

	if o.Query == "" {

		o.Query = util.GetString("query", data)
	}

	return err
}

func (o *DBDatasource) Next(rpt *core.Keireport) (map[string]interface{}, error) {

	var err error

	if o.Rows == nil {

		fmt.Println("=====> ", rpt.SQLDB)
		if rpt.SQLDB != nil {

			fmt.Printf("%T\r\n", rpt.SQLDB.Driver())
			fmt.Println("====>", reflect.TypeOf(rpt.SQLDB.Driver()).String())
		}

		//o.RawConn = rpt.SQLDB
		err = o.init()

		if err == nil {

			o.Rows = []map[string]interface{}{}

			targetSQL := rpt.ReplaceString(o.Query)

			tx := o.Conn

			if rpt.Debug {

				tx = tx.Debug()
			}

			err = tx.Raw(targetSQL).
				Find(&o.Rows).
				Error

			o.RowNo = -1
		}
	}

	if err == nil {

		if (o.RowNo + 1) < len(o.Rows) {

			o.RowNo++
			return o.Rows[o.RowNo], nil
		} else {

			err = core.ErrEndOfRow
		}
	}

	return nil, err
}

//------------------------------------------------

type DBDatasourceBuilder struct {
}

func (o *DBDatasourceBuilder) Build(data map[string]interface{}) (core.DataSource, error) {

	ret := &DBDatasource{
		ConnStr: util.GetString("connStr", data),
		DBType:  util.GetString("dbType", data),
		Query:   util.GetString("query", data),
	}

	return ret, nil
}

//------------------------------------------------

type DBConnBuilder interface {
	Build(ds *DBDatasource) (*gorm.DB, error)
}

type PostgresqlConnBuilder struct {
}

func (o *PostgresqlConnBuilder) Build(ds *DBDatasource) (*gorm.DB, error) {

	var ret *gorm.DB
	var err error

	if ds.RawConn != nil {

		ret, err = gorm.Open(postgres.New(postgres.Config{
			Conn: ds.RawConn,
		}), &gorm.Config{})
	} else {

		ret, err = gorm.Open(postgres.Open(ds.ConnStr), &gorm.Config{})
	}

	return ret, err
}

//------------------------------------------------

var dbConnBuilderMap map[string]DBConnBuilder = map[string]DBConnBuilder{}

func RegisterDBConnBuilder(name string, builder DBConnBuilder) {

	dbConnBuilderMap[name] = builder
}

func NewDBDatasource(dbType string, conn *sql.DB, query string) (core.DataSource, error) {

	ret := &DBDatasource{
		DBType:  dbType,
		RawConn: conn,
	}

	if query != "" {

		ret.Query = query
	}

	return ret, nil
}

func init() {

	core.RegisterDatasource("db", &DBDatasourceBuilder{})

	RegisterDBConnBuilder("postgresql", &PostgresqlConnBuilder{})
}
