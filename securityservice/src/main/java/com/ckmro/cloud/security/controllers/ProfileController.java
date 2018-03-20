package com.ckmro.cloud.security.controllers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.util.Map;

@Controller
public class ProfileController {


    private static final Logger log = LoggerFactory.getLogger(ErrorController.class);

    @RequestMapping("/oauth/profile")
    public String profile(@RequestParam(name="access_token") String accessToken) {


        return "";
    }

}
