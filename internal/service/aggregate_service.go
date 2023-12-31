package service

import (
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/aggregate"
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/repository"
)

type AggregateService struct {
	aggregateService repository.IAggregateRepo
}

func NewAggregateService(aggregateService repository.IAggregateRepo) *AggregateService {
	return &AggregateService{
		aggregateService: aggregateService,
	}
}

func (as *AggregateService) GetUsersWithColumnTypes() ([]aggregate.AggregatedData, error) {
	return as.aggregateService.GetUsersWithColumnTypes()
}
