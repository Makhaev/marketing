// ProductList.jsx
import React, { useEffect, useState } from "react";
import Button from "./Button";
import styles from "./ProductList.module.css";

export default function ProductList() {
  const [products, setProducts] = useState([]);
  const [name, setName] = useState("");
  const [category, setCategory] = useState("");

  useEffect(() => {
    fetch("/api/products")
      .then(res => res.json())
      .then(data => setProducts(data))
      .catch(console.error);
  }, []);

  const addProduct = async (e) => {
    e.preventDefault();
    const response = await fetch("/api/products", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, category }),
    });
    if (response.ok) {
      const newProduct = await response.json();
      setProducts([...products, newProduct]);
      setName("");
      setCategory("");
    } else {
      alert("Ошибка при создании продукта");
    }
  };

  const deleteProduct = async (id) => {
    const response = await fetch(`/api/products/${id}`, {
      method: "DELETE",
    });
    if (response.ok) {
      setProducts(products.filter(p => p.id !== id));
    } else {
      alert("Ошибка при удалении продукта");
    }
  };

  return (
    <div className={styles.container}>
      <h2>Продукты</h2>

      <form onSubmit={addProduct} className={styles.form}>
        <input
          type="text"
          placeholder="Название"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Категория"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
          required
        />
        <Button type="submit">Добавить продукт</Button>
      </form>

      <div className={styles.cardList}>
        {products.map(p => (
          <div key={p.id} className={styles.card}>
            <h3>{p.name}</h3>
            <p>Категория: {p.category}</p>
            <Button onClick={() => deleteProduct(p.id)} style={{backgroundColor:'#ef4444'}}>Удалить</Button>
          </div>
        ))}
      </div>
    </div>
  );
}