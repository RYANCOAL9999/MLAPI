# Java is not working GRPC with this Project on Apple Macbook M1 Computer.
It is no solution on build the grpc java file when I try to call command on terminal with auto generated or downloaded protoc-gen-grpc-java exec file.

# Maven Quick start Spring Boot3:
https://github.com/awslabs/aws-serverless-java-container/wiki/Quick-start---Spring-Boot3

# Start Servless Project:
mvn archetype:generate -DgroupId=com.quotes -DartifactId=quotes -Dversion=1.0-SNAPSHOT -DarchetypeGroupId=com.amazonaws.serverless.archetypes -DarchetypeArtifactId=aws-serverless-springboot-archetype -DarchetypeVersion=1.5.2

# Clean Project Package:
mvn clean package

# Download protoc-gen-grpc-java: 
https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.59.0/

# Auto generated protoc-gen-grpc-java: 
https://github.com/grpc/grpc-java/issues/7690 with glli commented on Aug 14

# Generated-code reference
https://grpc.io/docs/languages/java/generated-code/

# GRPC Java:
https://github.com/grpc/grpc-java/tree/master

#rules_proto_grpc reference:
https://github.com/rules-proto-grpc/rules_proto_grpc/tree/master

# Stop it on 
src/main/java/com/quotes/controller/PingController.
There are example on this reference how to call the protobuff file.

# Someone says that It can work on M2 with Rosetta installed:
https://github.com/grpc/grpc-java/issues/7690 with ejona86 commented 2 weeks ago and venshine commented on May 5