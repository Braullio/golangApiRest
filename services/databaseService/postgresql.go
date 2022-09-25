package databaseService

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Postgresql struct {
	Host        string
	User        string
	Password    string
	Dbname      string
	Port        string
	Sslmode     string
	Sslrootcert string
}

func BuildPostgresqlSql(host string, user string, password string, dbname string, port string, sslmode string, sslrootcert string) Postgresql {
	var postgresqlStruct Postgresql

	postgresqlStruct.Host = host
	postgresqlStruct.User = user
	postgresqlStruct.Password = password
	postgresqlStruct.Dbname = dbname
	postgresqlStruct.Port = port
	postgresqlStruct.Sslmode = sslmode
	postgresqlStruct.Sslrootcert = sslrootcert

	return postgresqlStruct
}

func DbConnect(postgresqlStruct Postgresql) *gorm.DB {
	if len(postgresqlStruct.Host) == 0 || len(postgresqlStruct.User) == 0 || len(postgresqlStruct.Password) == 0 || len(postgresqlStruct.Dbname) == 0 || len(postgresqlStruct.Port) == 0 {
		log.Fatalf("google.bigquery: Variables not valid")
	}

	var dsn string

	if len(postgresqlStruct.Sslmode) > 0 && len(postgresqlStruct.Sslrootcert) > 0 {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s", postgresqlStruct.Host, postgresqlStruct.User, postgresqlStruct.Password, postgresqlStruct.Dbname, postgresqlStruct.Port, postgresqlStruct.Sslmode, postgresqlStruct.Sslrootcert)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", postgresqlStruct.Host, postgresqlStruct.User, postgresqlStruct.Password, postgresqlStruct.Dbname, postgresqlStruct.Port)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err)
	}

	return db
}
