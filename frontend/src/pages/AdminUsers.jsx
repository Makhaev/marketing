import React, { useState, useEffect } from "react";

export default function AdminUsers() {
  const [users, setUsers] = useState([]);
  const [phone, setPhone] = useState("");
  const [role, setRole] = useState("user");

  useEffect(() => {
    fetch("/api/users")
      .then(res => res.json())
      .then(setUsers)
      .catch(console.error);
  }, []);

  const addUser = async (e) => {
    e.preventDefault();
    const res = await fetch("/api/users", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ phone, role }),
    });
    if (res.ok) {
      const newUser = await res.json();
      setUsers([...users, newUser]);
      setPhone("");
      setRole("user");
    } else {
      alert("Ошибка при создании пользователя");
    }
  };

  return (
    <div>
      <h2>Пользователи (Админка)</h2>
      <form onSubmit={addUser}>
        <input
          placeholder="Телефон"
          value={phone}
          onChange={e => setPhone(e.target.value)}
          required
        />
        <select value={role} onChange={e => setRole(e.target.value)}>
          <option value="user">User</option>
          <option value="owner">Owner</option>
        </select>
        <button type="submit">Добавить пользователя</button>
      </form>

      <ul>
        {users.map(u => (
          <li key={u.id}>{u.id}: {u.phone} ({u.role})</li>
        ))}
      </ul>
    </div>
  );
}