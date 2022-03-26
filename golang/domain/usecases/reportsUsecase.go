package usecases

import (
	"github.com/LuisDiazM/firestore-reports/domain"
	"github.com/LuisDiazM/firestore-reports/domain/models"
)

type ReportsUsecase struct {
	cacheGateway    domain.CacheGateway
	databaseGateway domain.DatabaseGateway
}

func NewReportUsecaseImp(cacheGateway domain.CacheGateway, databaseGateway domain.DatabaseGateway) ReportsUsecase {
	return ReportsUsecase{
		cacheGateway,
		databaseGateway,
	}
}

func (report ReportsUsecase) ExportDataUserResponses() (*[]models.UserResponsesDTO, error) {
	userResponses, err := report.databaseGateway.GetUsersResponses()
	if err != nil {
		return nil, err
	}
	return userResponses, nil
}
