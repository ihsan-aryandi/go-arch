package util

import (
	"fmt"
	"strings"
)

type dsnCreator struct {
	dsn []string
}

/*
	CreateDSN is user for creating data source name string
*/
func NewDSNCreator() *dsnCreator {
	return &dsnCreator{}
}

func (dc *dsnCreator) DBName(value string) *dsnCreator {
	dc.dsn = append(dc.dsn, keyValue("dbname", value))
	return dc
}

func (dc *dsnCreator) User(value string) *dsnCreator {
	dc.dsn = append(dc.dsn, keyValue("user", value))
	return dc
}

func (dc *dsnCreator) Password(value string) *dsnCreator {
	dc.dsn = append(dc.dsn, keyValue("password", value))
	return dc
}

func (dc *dsnCreator) Host(value string) *dsnCreator {
	dc.dsn = append(dc.dsn, keyValue("host", value))
	return dc
}

func (dc *dsnCreator) SSLMode(value string) *dsnCreator {
	dc.dsn = append(dc.dsn, keyValue("sslmode", value))
	return dc
}

func (dc *dsnCreator) Schema(value string) *dsnCreator {
	dc.dsn = append(dc.dsn, keyValue("search_path", value))
	return dc
}

func (dc *dsnCreator) Create() string {
	return strings.Join(dc.dsn, " ")
}

func keyValue(key string, value string) string {
	return fmt.Sprintf("%v=%v", key, value)
}