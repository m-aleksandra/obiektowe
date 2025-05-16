// ✅ src/components/Payment.js
import { useState } from "react";
import { useCart } from "../context/CartContext";
import "../styles.css";

export default function Payment() {
  const { cartId, clearCartState } = useCart();
  const [cardNumber, setCardNumber] = useState("");
  const [message, setMessage] = useState("");

  const pay = async () => {
    const res = await fetch("/api/payment", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ cartId: parseInt(cartId), cardNumber }),
    });

    const data = await res.json();
    setMessage(data.message || data.error);

    if (data.message === "Payment processed successfully") {
      await fetch(`/api/cart/${cartId}`, { method: "DELETE" });
      clearCartState();
    }
  };

  return (
    <div className="container">
      <h2>Payment</h2>
      <input
        className="input"
        value={cardNumber}
        onChange={(e) => setCardNumber(e.target.value)}
        placeholder="Card number"
      />
      <button className="button" onClick={pay}>Pay</button>
      {message && <p className="message">{message}</p>}
    </div>
  );
}