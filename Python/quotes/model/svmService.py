import grpc

from dataclasses import asdict
from lib.proto.py.svm_expression_pb2_grpc import SVMServiceStub
from lib.proto.py.svm_expression_pb2 import (
    LinearSVCRequest,
    LinearSVRRequest,    
    SVCRequest
)

class SVMService(object):
    def __init__(self, url:str)->None:
        self.__channel = grpc.insecure_channel(url)
        self.__stub = SVMServiceStub(self.__channel)

    def linearSVCResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LinearSVCRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LinearSVCEvent(request_message)
            return asdict(response)
        
    def linearSVRResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = LinearSVRRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.LinearSVREvent(request_message)
            return asdict(response)
        
    def svcResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = SVCRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                sample_weight = data.sample_weight | None,
                kwargs = data.kwargs
            )
            response = self.__stub.SVCEvent(request_message)
            return asdict(response)

