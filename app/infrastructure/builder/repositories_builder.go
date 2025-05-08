package builder

import (
	"prueba-tecnica-nauta/app/infrastructure/clients"
	"prueba-tecnica-nauta/app/infrastructure/config"
	"prueba-tecnica-nauta/app/infrastructure/repositories"
)

type RepositoriesBuilder struct {
	sqlRepository   *repositories.PostgresRepository
	cacheRepository *repositories.NautaCacheRepository
}

func NewRepositoriesBuilder(
	config *config.Config,
) (*RepositoriesBuilder, error) {
	pool, err := clients.NewPool(config.Database)
	if err != nil {
		return nil, err
	}
	sqlRepository := repositories.NewPostgresRepository(
		pool.FindAllQuery,
		pool.FindSingleQuery,
		pool.ExecuteQuery,
		config.Database.CacheDuration,
		config.Database.MaxFailures,
		config.Database.ResetTimeout,
	)

	cacheRepository, err := repositories.NewNautaCacheRepository(
		sqlRepository.GetBookings,
		sqlRepository.GetContainers,
		sqlRepository.GetOrders,
		sqlRepository.GetInvoices,
		sqlRepository.GetEmailClients,
	)
	if err != nil {
		return nil, err
	}

	return &RepositoriesBuilder{
		sqlRepository:   sqlRepository,
		cacheRepository: cacheRepository,
	}, nil
}
