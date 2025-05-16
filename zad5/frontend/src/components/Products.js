// ✅ src/components/Products.js
import { useEffect, useState } from "react";
import { useCart } from "../context/CartContext";
import "../styles.css";

export default function Products() {
  const [products, setProducts] = useState([]);
  const { addToCart } = useCart();

  useEffect(() => {
    fetch("/api/products")
      .then((res) => res.json())
      .then(setProducts);
  }, []);

  return (
    <div className="container">
      <h2>Products</h2>
      <ul>
        {products.map((p) => (
          <li key={p.id}>
            {p.name} – {p.price.toFixed(2)} zł
            <button className="button" onClick={() => addToCart(p.id)}>Add to cart</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
