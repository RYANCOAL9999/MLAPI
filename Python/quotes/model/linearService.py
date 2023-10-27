import grpc

from dataclasses import asdict
from lib.proto.py.linear_expression_pb2_grpc import LinearServiceStub
from lib.proto.py.general_expression_pb2 import (
    LinearRegressionRequest,
    LinearRidgeRequest,
    LinearRidgeCVRequest,
    LassoExpressionRequest,
    LassoLarsLassoExpressionRequest,
    BayesianRidgeRequest,
    TweedieRegressorRequest,
    SGDClassifierRequest,
    ElasticNetRequest
) 

class LinearService(object):
    def __init__(
            self, 
            url:str
        )->None:
        self.__channel = grpc.insecure_channel(url)
        self.__stub = LinearServiceStub(self.__channel)

    def linearRegressionResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LinearRegressionRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LinearRegressionEvent(request_message)
            return asdict(response)
        
    def ridgeResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LinearRidgeRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LinearRidgeEvent(request_message)
            return asdict(response)
        
    def ridgeCVResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LinearRidgeCVRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LinearRidgeCVEvent(request_message)
            return asdict(response)
        
    def lassoResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LassoExpressionRequest(
                alpha = data.alpha,
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LassoExpressionEvent(request_message)
            return asdict(response)
        
    def lassoLarsResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LassoLarsLassoExpressionRequest(
                alpha = data.alpha,
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LassoLarsLassoExpressionEvent(request_message)
            return asdict(response)
        
    def BayesianRidgeResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = BayesianRidgeRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.BayesianRidgeEvent(request_message)
            return asdict(response)
        
    def tweedieRegressorResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = TweedieRegressorRequest(
                alpha = data.alpha,
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.TweedieRegressorEvent(request_message)
            return asdict(response)
        
    def sgdClassifierResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = SGDClassifierRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.SGDClassifierEvent(request_message)
            return asdict(response)
        
    def elasticNetResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = ElasticNetRequest(
                alpha = data.alpha,
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.ElasticNetEvent(request_message)
            return asdict(response)

