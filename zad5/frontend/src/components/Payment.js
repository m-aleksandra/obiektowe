import { useCart } from "../context/CartContext";
import { useState } from "react";

export default function Payment() {
  const { items, clearCart } = useCart();
  const [message, setMessage] = useState("");

  const handlePayment = async () => {
    if (items.length === 0) {
      setMessage("Koszyk jest pusty.");
      return;
    }

    try {
      const cartRes = await fetch("/api/cart", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          items: items.map((item) => ({
            product: item.product,
            quantity: item.quantity,
          })),
        }),
      });

      if (!cartRes.ok) throw new Error("Błąd tworzenia koszyka");
      const cart = await cartRes.json();

      const paymentRes = await fetch("/api/payment", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ cartId: cart.id }),
      });

      if (!paymentRes.ok) throw new Error("Błąd płatności");
      const payment = await paymentRes.json();

      clearCart();
      setMessage(`Płatność zakończona sukcesem! ID: ${payment.id}`);
    } catch (err) {
      setMessage("Błąd: " + err.message);
    }
  };

  return (
    <div className="container">
      <h2>Płatność</h2>
      <button onClick={handlePayment}>Zapłać</button>
      {message && <p>{message}</p>}
    </div>
  );
}
