package middleware

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

func AutoMigrate() {
	connection, err := repository.NewDbContext()

	if err != nil {
		err.Error()
		println("automigrate repo patladı")
	} else {
		connection.Debug().AutoMigrate(
			&entity.Vehicle{},
			&entity.City{},
			&entity.CreditCard{},
			&entity.Customer{},
			&entity.Invoice{},
			&entity.Order{},
			&entity.Payment{},
			&entity.VehicleBrand{},
		)
	}

}
