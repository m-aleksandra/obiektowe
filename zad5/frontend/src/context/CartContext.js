import React, { createContext, useContext, useState } from "react";

const CartContext = createContext();

export const useCart = () => useContext(CartContext);

export const CartProvider = ({ children }) => {
  const [items, setItems] = useState([]);

  const addItem = (product, quantity = 1) => {
    setItems((prev) => {
      const found = prev.find((item) => item.product.id === product.id);
      if (found) {
        return prev.map((item) =>
          item.product.id === product.id
            ? { ...item, quantity: item.quantity + quantity }
            : item
        );
      }
      return [...prev, { product, quantity }];
    });
  };

  const clearCart = () => setItems([]);

  return (
    <CartContext.Provider value={{ items, addItem, clearCart }}>
      {children}
    </CartContext.Provider>
  );
};
