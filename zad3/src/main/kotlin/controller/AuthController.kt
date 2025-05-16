package com.example.authapp.controller

import com.example.authapp.model.User
import com.example.authapp.service.AuthService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/auth")
class AuthController {

    @Autowired
    private lateinit var authService: AuthService

    @PostMapping("/login")
    fun login(@RequestParam username: String, @RequestParam password: String): String {
        return if (authService.authorize(username, password)) {
            "Authorized"
        } else {
            "Unauthorized"
        }
    }

    @GetMapping("/users")
    fun getUsers(): List<User> = authService.getAllUsers()
}
