package devices

import (
	handlerPkg "tanahore/internal/app/devices/handler"
	repositoryPkg "tanahore/internal/app/devices/repository"
	routesPkg "tanahore/internal/app/devices/routes"
	servicePkg "tanahore/internal/app/devices/service"

	"firebase.google.com/go/db"
	"github.com/go-playground/validator"
)

func DeviceSetup(db *db.Client, validate *validator.Validate) routesPkg.DeviceRoute {
	repository := repositoryPkg.NewDeviceRepository(db)
	service := servicePkg.NewDeviceService(repository, validate)
	handler := handlerPkg.NewDeviceHandler(service)
	routes := routesPkg.NewDeviceRoutes(handler)

	return routes
}
