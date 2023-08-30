package interfaces

import "grpcmodel/models"

type ICustomer interface{
	CreateCustomer(Info *models.CustomerRequest)(*models.CustomerResponse,error)
}