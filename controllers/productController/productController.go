package productController

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"golangApiRest/models"
	"golangApiRest/services/databaseService"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

func envValues() (string, string, string, string, string, string, string) {
	Host := os.Getenv("POSTGRESQL_HOST")
	User := os.Getenv("POSTGRESQL_USER")
	Port := os.Getenv("POSTGRESQL_PORT")
	Dbname := os.Getenv("POSTGRESQL_DBNAME")
	Password := os.Getenv("POSTGRESQL_PASSWORD")

	var Sslmode string
	if len(os.Getenv("POSTGRESQL_SSLMODE")) == 0 {
		Sslmode = os.Getenv("POSTGRESQL_SSLMODE")
	} else {
		Sslmode = ""
	}

	var Sslrootcert string
	if len(os.Getenv("POSTGRESQL_SSLMODE")) == 0 {
		Sslrootcert = os.Getenv("POSTGRESQL_SSLROOTCERT")
	} else {
		Sslrootcert = ""
	}

	return Host, User, Password, Dbname, Port, Sslmode, Sslrootcert
}

func Show(c *fiber.Ctx) error {
	host, user, password, dbname, port, sslmode, sslrootcert := envValues()
	postgresqlStruct := databaseService.BuildPostgresqlSql(host, user, password, dbname, port, sslmode, sslrootcert)

	var products []models.Product

	var errors *gorm.DB
	if len(c.Params("id")) > 0 {
		product := searchProduct(c, postgresqlStruct)

		if product.Id > 0 {
			products = append(products, product)
		} else {
			return c.SendStatus(fiber.StatusBadRequest)
		}
	} else {
		errors = databaseService.DbConnect(postgresqlStruct).Find(&products)

		if errors.Error != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(&products)
}

func Create(c *fiber.Ctx) error {
	host, user, password, dbname, port, sslmode, sslrootcert := envValues()
	postgresqlStruct := databaseService.BuildPostgresqlSql(host, user, password, dbname, port, sslmode, sslrootcert)
	databaseService.DbConnect(postgresqlStruct).AutoMigrate(&models.Product{})

	var product models.Product

	err := json.Unmarshal(c.Body(), &product)
	if err != nil {
		log.Fatalln(err)
	}

	errors := databaseService.DbConnect(postgresqlStruct).Create(&product)

	if errors.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}

func Update(c *fiber.Ctx) error {
	host, user, password, dbname, port, sslmode, sslrootcert := envValues()
	postgresqlStruct := databaseService.BuildPostgresqlSql(host, user, password, dbname, port, sslmode, sslrootcert)

	product := searchProduct(c, postgresqlStruct)
	if product.Id == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := json.Unmarshal(c.Body(), &product)
	if err != nil {
		log.Fatalln(err)
	}

	timeNow := time.Now()
	product.UpdatedAt = timeNow

	errors := databaseService.DbConnect(postgresqlStruct).Updates(&product)
	if errors.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}

func Delete(c *fiber.Ctx) error {
	host, user, password, dbname, port, sslmode, sslrootcert := envValues()
	postgresqlStruct := databaseService.BuildPostgresqlSql(host, user, password, dbname, port, sslmode, sslrootcert)

	product := searchProduct(c, postgresqlStruct)
	if product.Id == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	errors := databaseService.DbConnect(postgresqlStruct).Delete(&product)
	if errors.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func searchProduct(c *fiber.Ctx, postgresqlStruct databaseService.Postgresql) models.Product {
	var product models.Product

	id, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	product.Id = uint(id)

	errors := databaseService.DbConnect(postgresqlStruct).First(&product)

	if errors.Error != nil {
		return models.Product{}
	}

	return product
}
