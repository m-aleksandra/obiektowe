import React from "react";
import "./App.css";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Products from "./components/Products";
import Cart from "./components/Cart";
import Payment from "./components/Payment";
import { CartProvider } from "./context/CartContext";

function App() {
  return (
    <CartProvider>
      <Router>
        <nav style={{ padding: "1rem", background: "#eee" }}>
          <Link to="/">Produkty</Link> |{" "}
          <Link to="/cart">Koszyk</Link> |{" "}
          <Link to="/payment">Płatność</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Products />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/payment" element={<Payment />} />
        </Routes>
      </Router>
    </CartProvider>
  );
}

export default App;
