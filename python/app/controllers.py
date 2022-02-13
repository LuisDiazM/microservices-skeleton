from dependency_injector.wiring import Provide, inject
from flask_restful import Resource
from domain.usecases.data_processing_usecase import SpectrumProcessingUseCase
from app.di import Container


class SpectrumHandler(Resource):

    @inject
    def __init__(self, spectrum_usecase: SpectrumProcessingUseCase = Provide[Container.spectrum_processing_usecase]) -> None:
        self.spectrum_usecase = spectrum_usecase
        super().__init__()

    def get(self):
        spectrum = self.spectrum_usecase.apply_return_loss_correction()
        return {'spectrum': spectrum}, 200
