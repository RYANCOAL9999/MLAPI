package com.quotes.controller;

import io.grpc.channel;
import io.grpc.Grpc;
import io.grpc.InsecureChannelCredentials;
import io.grpc.ManagedChannel;
import io.grpc.StatusRuntimeException;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import java.util.HashMap;
import java.util.Map;
import com.quotes.AIProto;

@RestController
public class PingController {
    @RequestMapping(path = "/ping", method = RequestMethod.GET)
    public Map<String, String> ping() {
        Map<String, String> pong = new HashMap<>();

        // HeaderRequest request = new HeaderRequest();

        // DataFrame response = new DataFrame();

        // resonse = stub.HeaderEvent(request);
        
        pong.put("pong", "Hello, World!");
        return pong;
    }
}
