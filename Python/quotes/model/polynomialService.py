import grpc

from dataclasses import asdict
from lib.proto.py.polynomial_features_pb2_grpc import PolynomialServiceStub
from lib.proto.py.polynomial_features_pb2 import PolynomialFeaturesRequest, PolynomialFeaturesFitTransformRequest

class PolynomialService(object):
    def __init__(
            self, 
            url:str
        )->None:
        self.__channel = grpc.insecure_channel(url)
        self.__stub = PolynomialServiceStub(self.__channel)

    def generalOrder(data) ->str:
        ans = 'C'
        if(data.order.value == 1):
            ans = 'F'
        return ans

    def polynomialFeaturesResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = PolynomialFeaturesRequest(
                x_drop_data = data.x_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                y_drop_data = data.y_drop_data | None,
                degree = data.degree | None,
                interaction_only = data.interaction_only | None ,
                include_bias = data.include_bias | None,
                order = self.generalOrder(data),
                kwargs = data.kwargs | None
            )
            response = self.__stub.PolynomialFeaturesEvent(request_message)
            return asdict(response)
        
    def polynomialFeaturesFitTransformResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = PolynomialFeaturesFitTransformRequest(
                x_drop_data = data.x_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                y_drop_data = data.y_drop_data | None,
                degree = data.degree | None,
                interaction_only = data.interaction_only | None,
                include_bias = data.include_bias | None,
                order = self.generalOrder(data),
                coef_init = data.coef_init | None,
                intercept_init = data.intercept_init | None,
                kwargs = data.kwargs | None
            )
            response = self.__stub.PolynomialFeaturesFitTransformEvent(request_message)
            return asdict(response)

