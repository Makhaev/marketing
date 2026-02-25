import React, { useEffect, useState } from "react";
import styles from "./StoreList.module.css"; // <-- подключаем CSS

export default function StoreList() {
  const [stores, setStores] = useState([]);
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");

  useEffect(() => {
    fetch("/api/stores")
      .then(res => res.json())
      .then(setStores)
      .catch(console.error);
  }, []);

  const addStore = async (e) => {
  e.preventDefault();
  const response = await fetch("/api/stores", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name, address, owner_id: 1 }) // OwnerID теперь учитывается
  });
  if (response.ok) {
    const newStore = await response.json();
    setStores([...stores, newStore]);
    setName("");
    setAddress("");
  } else {
    const err = await response.json();
    alert("Ошибка: " + err.error);
  }
};

  return (
    <div className={styles.container}>
      <h2>Магазины</h2>
      <form onSubmit={addStore}>
        <input
          type="text"
          placeholder="Название"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Адрес"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          required
        />
        <button type="submit">Добавить магазин</button>
      </form>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Название</th>
            <th>Адрес</th>
          </tr>
        </thead>
        <tbody>
          {stores.map(s => (
            <tr key={s.id}>
              <td>{s.id}</td>
              <td>{s.name}</td>
              <td>{s.address}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}