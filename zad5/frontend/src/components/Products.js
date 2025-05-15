import { useEffect, useState } from "react";
import { useCart } from "../context/CartContext";

export default function Products() {
  const [products, setProducts] = useState([]);
  const { addItem } = useCart();

  useEffect(() => {
    fetch("/api/products")
      .then((res) => res.json())
      .then(setProducts)
      .catch(console.error);
  }, []);

  const handleAdd = (product) => addItem(product, 1);

  return (
    <div className="container">
      <h2>Produkty</h2>
      {products.length === 0 ? (
        <p>Brak dostępnych produktów.</p>
      ) : (
        <ul>
          {products.map((p) => (
            <li key={p.id}>
              <strong>{p.name}</strong> – {p.price.toFixed(2)} zł
              <br />
              <button onClick={() => handleAdd(p)}>Dodaj do koszyka</button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
