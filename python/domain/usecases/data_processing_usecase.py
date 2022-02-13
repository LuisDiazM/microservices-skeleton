
from typing import List
from infraestructure.messaging.nats import NatsImp
from infraestructure.databases.mongo import MongoImp


class SpectrumProcessingUseCase:

    def __init__(self, database: MongoImp, messaging: NatsImp) -> None:
        self.database = database
        self.messaging = messaging
    
    def apply_return_loss_correction(self) -> List:
        return [1, 2, 3, 4]
