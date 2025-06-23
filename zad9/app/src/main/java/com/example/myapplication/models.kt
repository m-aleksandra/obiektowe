package com.example.myapplication

data class Category(val id: Int, val name: String)
data class Product(val id: Int, val name: String, val categoryId: Int)
