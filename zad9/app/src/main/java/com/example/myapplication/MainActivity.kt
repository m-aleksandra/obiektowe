package com.example.myapplication

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import androidx.navigation.NavHostController
import androidx.navigation.compose.*
import com.example.myapplication.ui.theme.MyApplicationTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            MyApplicationTheme {
                val navController = rememberNavController()
                AppNavigation(navController)
            }
        }
    }
}

// --- NAWIGACJA ---
@Composable
fun AppNavigation(navController: NavHostController) {
    NavHost(navController = navController, startDestination = "categories") {
        composable("categories") {
            CategoryListScreen(navController)
        }
        composable("products/{categoryId}") { backStackEntry ->
            val categoryId = backStackEntry.arguments?.getString("categoryId")?.toIntOrNull() ?: 0
            ProductListScreen(categoryId)
        }
    }
}

// --- EKRAN 1: Kategorie ---
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CategoryListScreen(navController: NavController) {
    val categories = listOf(
        Category(1, "Elektronika"),
        Category(2, "Książki"),
        Category(3, "Ubrania")
    )

    Scaffold(
        topBar = {
            TopAppBar(title = { Text("Kategorie") })
        }
    ) { padding ->
        Column(modifier = Modifier
            .padding(padding)
            .padding(16.dp)) {
            categories.forEach { category ->
                Text(
                    text = category.name,
                    style = MaterialTheme.typography.titleMedium,
                    modifier = Modifier
                        .fillMaxWidth()
                        .clickable {
                            navController.navigate("products/${category.id}")
                        }
                        .padding(12.dp)
                )
            }
        }
    }
}

// --- EKRAN 2: Produkty ---
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ProductListScreen(categoryId: Int) {
    val allProducts = listOf(
        Product(1, "Laptop", 1),
        Product(2, "Smartfon", 1),
        Product(3, "Harry Potter", 2),
        Product(4, "Dresy", 3)
    )

    val products = allProducts.filter { it.categoryId == categoryId }

    Scaffold(
        topBar = {
            TopAppBar(title = { Text("Produkty") })
        }
    ) { padding ->
        Column(modifier = Modifier
            .padding(padding)
            .padding(16.dp)) {
            products.forEach { product ->
                Text(
                    text = product.name,
                    style = MaterialTheme.typography.bodyLarge,
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(8.dp)
                )
            }
        }
    }
}
