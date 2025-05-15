import { useCart } from "../context/CartContext";
import { Link } from "react-router-dom";

export default function Cart() {
  const { items } = useCart();

  const total = items.reduce(
    (sum, item) => sum + item.product.price * item.quantity,
    0
  );

  return (
    <div className="container">
      <h2>Koszyk</h2>
      {items.length === 0 ? (
        <p>Koszyk jest pusty.</p>
      ) : (
        <>
          <ul>
            {items.map((item, index) => (
              <li key={index}>
                {item.product.name} – {item.quantity} × {item.product.price.toFixed(2)} zł
              </li>
            ))}
          </ul>
          <p><strong>Razem:</strong> {total.toFixed(2)} zł</p>
          <Link to="/payment">
            <button>Przejdź do płatności</button>
          </Link>
        </>
      )}
    </div>
  );
}
