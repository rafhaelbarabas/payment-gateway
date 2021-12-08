package factory

import "github.com/rafhaelbarabas/payment-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
