from dependency_injector import containers, providers
from domain.usecases.data_processing_usecase import SpectrumProcessingUseCase
from infraestructure.messaging.nats import NatsImp
from infraestructure.databases.mongo import MongoImp


class Container(containers.DeclarativeContainer):
    # infraestructure
    config = providers.Configuration()
    database = providers.Singleton(MongoImp)
    messaging = providers.Singleton(NatsImp)

    # usecases
    spectrum_processing_usecase = providers.Factory(
        SpectrumProcessingUseCase, database=database, messaging=messaging)
