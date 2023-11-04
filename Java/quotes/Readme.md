# Java Is Not Working GRPC With This Project On Apple Macbook M1 Computer.
It is no solution on build the grpc java file when I try to call command on terminal with auto generated or downloaded protoc-gen-grpc-java exec file.

# Maven Quick Start Spring Boot3:
https://github.com/awslabs/aws-serverless-java-container/wiki/Quick-start---Spring-Boot3

# Start Servless Project:
mvn archetype:generate -DgroupId=com.quotes -DartifactId=quotes -Dversion=1.0-SNAPSHOT -DarchetypeGroupId=com.amazonaws.serverless.archetypes -DarchetypeArtifactId=aws-serverless-springboot-archetype -DarchetypeVersion=1.5.2

# Clean Project Package:
mvn clean package

# Download Protoc-gen-grpc-java: 
https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.59.0/

# Auto Generated Protoc-gen-grpc-java: 
https://github.com/grpc/grpc-java/issues/7690 with glli commented on Aug 14

# Generated-code Reference
https://grpc.io/docs/languages/java/generated-code/

# GRPC Java:
https://github.com/grpc/grpc-java/tree/master

# Bazel:
https://github.com/bazelbuild/bazel

# Gradle:
https://github.com/gradle/gradle

# Rules_proto_grpc Reference:
https://github.com/rules-proto-grpc/rules_proto_grpc/tree/master

# Google Resources Reference:
https://chromium.googlesource.com/infra/goma/server/+/48224f31513604484d7c74afa5fd9d7e621e1691/proto/google

# Stop It On 
src/main/java/com/quotes/controller/PingController

# Example For Call Protobuf:
import com.quotes.AIProto;
HeaderRequest request = new HeaderRequest();
DataFrame response = new DataFrame();
resonse = stub.HeaderEvent(request);

# Someone says that It can work on M2 with Rosetta installed:
https://github.com/grpc/grpc-java/issues/7690 with ejona86 commented 2 weeks ago and venshine commented on May 5
