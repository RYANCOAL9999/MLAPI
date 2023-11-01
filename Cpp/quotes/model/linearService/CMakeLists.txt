cmake_minimum_required(VERSION 3.8)

project(quotes C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(ls_proto "../../../../../../../AIProto/lib/enum_expression.proto" ABSOLUTE)
get_filename_component(ls_proto "../../../../../../../AIProto/lib/np_library.proto" ABSOLUTE)
get_filename_component(ls_proto "../../../../../../../AIProto/linear_expression.proto" ABSOLUTE)
get_filename_component(ls_proto_path "${ls_proto}" PATH)

# Generated sources
set(enum_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.cc")
set(enum_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/enum_expression.pb.h")
set(np_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.cc")
set(np_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/lib/np_library.pb.h")
set(ls_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/linear_expression.pb.cc")
set(ls_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/linear_expression.pb.h")
set(ls_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/linear_expression.grpc.pb.cc")
set(ls_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/linear_expression.grpc.pb.h")

add_custom_command(
    OUTPUT "${enum_proto_srcs}" "${enum_proto_hdrs}" "${np_proto_srcs}" "${np_proto_srcs}" "${ls_proto_srcs}" "${ls_proto_hdrs}" "${ls_grpc_srcs}" "${ls_grpc_hdrs}"
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
    --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
    -I "${ls_proto_path}"
    --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
    "${ls_proto}"
    DEPENDS "${ls_proto}"
)

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# ls_grpc_proto
add_library(
    ls_grpc_proto
    ${enum_proto_srcs}
    ${enum_proto_hdrs} 
    ${np_proto_srcs} 
    ${np_proto_srcs}
    ${ls_grpc_srcs}
    ${ls_grpc_hdrs}
    ${ls_proto_srcs}
    ${ls_proto_hdrs}
)

target_link_libraries(
    ls_grpc_proto
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF}
)

# add_library(
#   helper
#   "helper.h"
#   "helper.cc"
# )
# target_link_libraries(
#   helper
#   ls_grpc_proto
#   ${_REFLECTION}
#   ${_GRPC_GRPCPP}
#   ${_PROTOBUF_LIBPROTOBUF}
# )

# Targets route_guide_(client|server)
foreach(
    _target
    linearServiceClient
    # linearServiceServer
)
    add_executable(
        ${_target}
        "${_target}.cc"
    )
    target_link_libraries(
        ${_target}
        ls_grpc_proto
        # helper
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
    )
endforeach()














