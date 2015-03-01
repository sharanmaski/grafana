package models

import (
	"errors"
	"time"
)

const (
	DS_GRAPHITE      = "graphite"
	DS_INFLUXDB      = "influxdb"
	DS_INFLUXDB_08   = "influxdb_08"
	DS_ES            = "elasticsearch"
	DS_OPENTSDB      = "opentsdb"
	DS_ACCESS_DIRECT = "direct"
	DS_ACCESS_PROXY  = "proxy"
)

// Typed errors
var (
	ErrDataSourceNotFound = errors.New("Data source not found")
)

type DsAccess string

type DataSource struct {
	Id      int64
	OrgId   int64
	Version int

	Name              string
	Type              string
	Access            DsAccess
	Url               string
	Password          string
	User              string
	Database          string
	BasicAuth         bool
	BasicAuthUser     string
	BasicAuthPassword string
	IsDefault         bool
	JsonData          map[string]interface{}

	Created time.Time
	Updated time.Time
}

// ----------------------
// COMMANDS

// Also acts as api DTO
type AddDataSourceCommand struct {
	OrgId     int64 `json:"-"`
	Name      string
	Type      string
	Access    DsAccess
	Url       string
	Password  string
	Database  string
	User      string
	IsDefault bool

	Result *DataSource
}

// Also acts as api DTO
type UpdateDataSourceCommand struct {
	Id        int64                  `json:"id" binding:"Required"`
	Name      string                 `json:"name" binding:"Required"`
	Type      string                 `json:"type" binding:"Required"`
	Access    DsAccess               `json:"access" binding:"Required"`
	Url       string                 `json:"url"`
	Password  string                 `json:"password"`
	User      string                 `json:"user"`
	Database  string                 `json:"database"`
	IsDefault bool                   `json:"isDefault"`
	JsonData  map[string]interface{} `json:"jsonData"`

	OrgId int64 `json:"-"`
}

type DeleteDataSourceCommand struct {
	Id    int64
	OrgId int64
}

// ---------------------
// QUERIES

type GetDataSourcesQuery struct {
	OrgId  int64
	Result []*DataSource
}

type GetDataSourceByIdQuery struct {
	Id     int64
	OrgId  int64
	Result DataSource
}

type GetDataSourceByNameQuery struct {
	Name   string
	OrgId  int64
	Result DataSource
}

// ---------------------
// EVENTS
type DataSourceCreatedEvent struct {
}