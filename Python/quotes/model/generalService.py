import grpc

from dataclasses import asdict
from lib.proto.py.general_expression_pb2_grpc import GeneralServiceStub
from lib.proto.py.general_expression_pb2 import HeaderRequest, InfoRequest, DescribeRequest

class GeneralService(object):
    def __init__(
            self, 
            url:str
        )->None:
        self.__channel = grpc.insecure_channel(url)
        self.__stub = GeneralServiceStub(self.__channel)

    def headerResponse(self)->dict:
        with self.__channel:
            request_message = HeaderRequest()
            response = self.__stub.HeaderEvent(request_message)
            return asdict(response)
        
    def infoResponse(self)->dict:
        with self.__channel:
            request_message = InfoRequest()
            response = self.__stub.InfoEvent(request_message)
            return asdict(response)
        
    def describlerResponse(self)->dict:
        with self.__channel:
            request_message = DescribeRequest()
            response = self.__stub.DescriblerEvent(request_message)
            return asdict(response)

