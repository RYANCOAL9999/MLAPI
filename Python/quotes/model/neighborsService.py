import grpc

from dataclasses import asdict
from lib.proto.py.nearest_neighbors_pb2_grpc import NeighborsServiceStub
from lib.proto.py.nearest_neighbors_pb2 import NearestNeighborsRequest, KDTreeRequest, NearestCentroidRequest


class NeighborsService(object):
    def __init__(
            self, 
            url:str
        )->None:
        self.__channel = grpc.insecure_channel(url)
        self.__stub = NeighborsServiceStub(self.__channel)

    def nearestNeighborsResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = NearestNeighborsRequest(
                x_drop_data = data.x_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                y_drop_data = data.y_drop_data | None,
                kwargs = data.kwargs
            )
            response = self.__stub.NearestNeighborsEvent(request_message)
            return asdict(response)
        
    def kdTreeResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = KDTreeRequest(
                x_drop_data = data.x_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                y_drop_data = data.y_drop_data | None,
                sample_weight = data.sample_weight | None,
                k = data.k | None,
                leaf_size = data.leaf_size | None,
                metric = data.metric | None,
                returns_distance = data.returns_distance | None,
                dualtree = data.dualtree | None,
                breadth_first = data.breadth_first | None,
                sort_results = data.sort_results | None,
                kwargs = data.kwargs
            )
            response = self.__stub.KDTreeEvent(request_message)
            return asdict(response)
        
    def nearestCentroidResponse(
            self, 
            data
        )->dict:
        with self.__channel:
            request_message = NearestCentroidRequest(
                x_drop_data = data.x_drop_data,
                y_drop_data = data.y_drop_data,
                size = data.size,
                random = data.random,
                key = data.key,
                metric = data.metric | None,
                shrink_threshold = data.shrink_threshold | None
            )
            response = self.__stub.NearestCentroidEvent(request_message)
            return asdict(response)

