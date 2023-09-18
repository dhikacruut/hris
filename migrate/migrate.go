package main

import (
	"github.com/dhikacruut/hris/db"
	"github.com/dhikacruut/hris/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(
		&models.Grade{},
		&models.JobDescription{},
		&models.Level{},
		&models.Division{},
		&models.Department{},
		&models.Supervision{},
		&models.PTKP{},
		&models.User{},
		&models.Employee{},
		&models.SalarySlip{},
		&models.SalarySlipDetail{},
		&models.PPH{},
	)
}
