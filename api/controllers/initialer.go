package controllers

import (
	"fmt"
)

//Initialer initial interfacer
type Initialer interface {
	initialize(DbUser, DbPasswd, DbPort, DbHost, DbName string) string
}

//PostgresInitialer postgres string
type PostgresInitialer struct {
}

func (p *PostgresInitialer) initialize(DbUser, DbPasswd, DbPort, DbHost, DbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPasswd)
}

//MysqlInitialer mysql string
type MysqlInitialer struct {
}

func (p *MysqlInitialer) initialize(DbUser, DbPasswd, DbPort, DbHost, DbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbHost, DbPort, DbUser, DbName, DbPasswd)
}
