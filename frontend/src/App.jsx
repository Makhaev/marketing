import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Stores from "./pages/Stores";
import Products from "./pages/Products";
import StoreProducts from "./pages/StoreProducts";
import AdminUsers from "./pages/AdminUsers"; // <-- импортируем админку

export default function App() {
  return (
    <Router>
      <header style={{ padding: "10px", borderBottom: "1px solid #ccc" }}>
        <Link to="/stores"><button>Магазины</button></Link>
        <Link to="/products"><button>Продукты</button></Link>
        <Link to="/store-products"><button>Товары в магазинах</button></Link>
        <Link to="/admin/users"><button>Админка</button></Link> {/* <-- сюда */}
      </header>

      <main style={{ padding: "20px" }}>
        <Routes>
          <Route path="/stores" element={<Stores />} />
          <Route path="/products" element={<Products />} />
          <Route path="/store-products" element={<StoreProducts />} />
          <Route path="/admin/users" element={<AdminUsers />} /> {/* <-- сюда */}
          <Route path="*" element={<Stores />} /> {/* редирект на стартовую страницу */}
        </Routes>
      </main>
    </Router>
  );
}