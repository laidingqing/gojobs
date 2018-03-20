package com.ckmro.cloud.security;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.security.oauth2.provider.endpoint.CheckTokenEndpoint;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.SessionAttributes;

import java.security.Principal;

@SpringBootApplication
@EnableAutoConfiguration
@SessionAttributes("authorizationRequest")
@ComponentScan("com.ckmro.cloud.security")
@RestController
public class OAuthApplication {

    private Logger logger = LoggerFactory.getLogger(OAuthApplication.class);

    public static void main(String[] args) {
        SpringApplication.run(OAuthApplication.class, args);
    }

    @RequestMapping("/user")
    Principal principal(Principal principal) {
        logger.info("sso user is:" + principal);

        return principal;
    }
}
