import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Products from "./components/Products";
import Payment from "./components/Payment";
import Cart from "./components/Cart";
import { CartProvider } from "./context/CartContext";
import "./styles.css";

function App() {
  return (
    <CartProvider>
      <Router>
        <nav className="container">
          <Link to="/">Products</Link> | <Link to="/cart">Cart</Link> | <Link to="/payment">Payment</Link>
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
