import { useEffect, useState } from "react";
import { useCart } from "../context/CartContext";
import "../styles.css";

export default function Cart() {
  const { cartId } = useCart();
  const [items, setItems] = useState([]);
  const [detailedItems, setDetailedItems] = useState([]);

  useEffect(() => {
    if (!cartId) return;

    fetch(`/api/cart/${cartId}`)
      .then((res) => res.json())
      .then(async (data) => {
        setItems(data.items || []);

        const promises = data.items.map(async (item) => {
          const res = await fetch(`/api/products/${item.productId}`);
          const product = await res.json();
          return {
            productId: item.productId,
            name: product.name,
            price: product.price,
            quantity: item.quantity,
          };
        });

        const result = await Promise.all(promises);
        setDetailedItems(result);
      });
  }, [cartId]);

  const removeItem = async (productId) => {
    await fetch(`/api/cart/${cartId}/item/${productId}`, {
      method: "DELETE",
    });
    window.location.reload();
  };

  const total = detailedItems.reduce(
    (sum, i) => sum + i.price * i.quantity,
    0
  ).toFixed(2);

  return (
    <div className="container">
      <h2>Your Cart</h2>
      {detailedItems.length === 0 ? (
        <p>The cart is empty.</p>
      ) : (
        <>
          <ul>
            {detailedItems.map((item) => (
              <li key={item.productId}>
                {item.name} – {item.quantity} × {item.price.toFixed(2)} zł
                <button className="button" onClick={() => removeItem(item.productId)}>
                  Remove
                </button>
              </li>
            ))}
          </ul>
          <p>
            <strong>Total:</strong> {total} zł
          </p>
        </>
      )}
    </div>
  );
}
