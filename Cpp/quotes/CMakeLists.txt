cmake_minimum_required(VERSION 3.5)
set(CMAKE_CXX_STANDARD 11)

project(quotes LANGUAGES CXX)

find_package(aws-lambda-runtime REQUIRED)
find_package(AWSSDK COMPONENTS core)
find_package(gRPC CONFIG REQUIRED)

add_executable(${PROJECT_NAME} "main.cpp")
target_link_libraries(${PROJECT_NAME} PUBLIC AWS::aws-lambda-runtime ${AWSSDK_LINK_LIBRARIES})
target_link_libraries(${PROJECT_NAME} gRPC::grpc++)

aws_lambda_package_target(${PROJECT_NAME})