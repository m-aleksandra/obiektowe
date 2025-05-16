package com.example.authapp.service

import com.example.authapp.model.User
import jakarta.annotation.PostConstruct
import org.springframework.stereotype.Service

@Service
class AuthService {

    private val users = mutableListOf<User>()

    @PostConstruct
    fun init() {
        users.add(User("admin", "admin123"))
        users.add(User("user", "user123"))
    }

    fun authorize(username: String, password: String): Boolean {
        return users.any { it.username == username && it.password == password }
    }

    fun getAllUsers(): List<User> = users
}
