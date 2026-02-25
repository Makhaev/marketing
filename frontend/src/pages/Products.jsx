import React, { useEffect, useState } from "react";
import styles from "./ProductList.module.css"; // <-- CSS Modules

export default function ProductList() {
  const [products, setProducts] = useState([]);
  const [name, setName] = useState("");
  const [category, setCategory] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/products")
      .then(res => res.json())
      .then(setProducts)
      .catch(console.error);
  }, []);

  const addProduct = async (e) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/products", {
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

  return (
    <div className={styles.container}>
      <h2>Продукты</h2>
      <form onSubmit={addProduct}>
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
        <button type="submit">Добавить продукт</button>
      </form>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Название</th>
            <th>Категория</th>
          </tr>
        </thead>
        <tbody>
          {products.map(p => (
            <tr key={p.id}>
              <td>{p.id}</td>
              <td>{p.name}</td>
              <td>{p.category}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}