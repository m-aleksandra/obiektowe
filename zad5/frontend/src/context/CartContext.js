import React, { createContext, useContext, useEffect, useState } from "react";

const CartContext = createContext();
export const useCart = () => useContext(CartContext);

export const CartProvider = ({ children }) => {
  const [cartId, setCartId] = useState(localStorage.getItem("cartId"));
  const [items, setItems] = useState([]);

  const createCart = async () => {
    const res = await fetch("/api/cart", { method: "POST" });
    const data = await res.json();
    setCartId(data.id);
    localStorage.setItem("cartId", data.id);
    setItems([]);
    console.log("Utworzono nowy koszyk:", data.id);
  };

  const addToCart = async (productId) => {
    if (!cartId) await createCart();
    const currentCartId = localStorage.getItem("cartId");
    const res = await fetch(`/api/cart/${currentCartId}/add`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ productId, quantity: 1 }),
    });
    if (res.ok) {
      console.log(`Dodano produkt ${productId} do koszyka ${currentCartId}`);
    }
  };

  const clearCartState = () => {
    setItems([]);
  };

  useEffect(() => {
    if (!cartId) return;
    fetch(`/api/cart/${cartId}`).then((res) => {
      if (res.status === 404) {
        localStorage.removeItem("cartId");
        setCartId(null);
        createCart();
      }
    });
  }, [cartId]);

  return (
    <CartContext.Provider value={{ cartId, addToCart, clearCartState }}>
      {children}
    </CartContext.Provider>
  );
};
