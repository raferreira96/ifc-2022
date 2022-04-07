package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/raferreira96/ifc-2022/codepix/application/usecase"
	"github.com/raferreira96/ifc-2022/codepix/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}
	return transactionUseCase
}
