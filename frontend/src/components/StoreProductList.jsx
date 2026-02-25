import React, { useEffect, useState } from "react";
import styles from "./StoreProductList.module.css";

export default function StoreProductList() {
  const [storeProducts, setStoreProducts] = useState([]);
  const [stores, setStores] = useState([]);
  const [products, setProducts] = useState([]);
  const [storeId, setStoreId] = useState("");
  const [productId, setProductId] = useState("");
  const [price, setPrice] = useState("");
  const [isPromo, setIsPromo] = useState(false);
  const [imageUrl, setImageUrl] = useState("");

  useEffect(() => {
    fetch("/api/store-products")
      .then(res => res.json())
      .then(setStoreProducts)
      .catch(console.error);

    fetch("/api/stores")
      .then(res => res.json())
      .then(setStores)
      .catch(console.error);

    fetch("/api/products")
      .then(res => res.json())
      .then(setProducts)
      .catch(console.error);
  }, []);

 const addStoreProduct = async (e) => {
  e.preventDefault();
  const response = await fetch("/api/store-products", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      store_id: parseInt(storeId),
      product_id: parseInt(productId),
      price: parseFloat(price),
      is_promo: isPromo,
      image_url: imageUrl || null
    }),
  });

  if (response.ok) {
    const newSP = await response.json();
    setStoreProducts([...storeProducts, newSP]);
    setStoreId("");
    setProductId("");
    setPrice("");
    setIsPromo(false);
    setImageUrl("");
  } else {
    const err = await response.json();
    alert("Ошибка: " + err.error);
  }
};

  return (
    <div className={styles.container}>
      <h2>Товары в магазинах</h2>

      <form onSubmit={addStoreProduct}>
        <select
          value={storeId}
          onChange={e => setStoreId(e.target.value)}
          required
        >
          <option value="">Выберите магазин</option>
          {stores.map(s => (
            <option key={s.id} value={s.id}>
              {s.name}
            </option>
          ))}
        </select>

        <select
          value={productId}
          onChange={e => setProductId(e.target.value)}
          required
        >
          <option value="">Выберите продукт</option>
          {products.map(p => (
            <option key={p.id} value={p.id}>
              {p.name}
            </option>
          ))}
        </select>

        <input
          type="number"
          step="0.01"
          placeholder="Цена"
          value={price}
          onChange={e => setPrice(e.target.value)}
          required
        />

        <label>
          <input
            type="checkbox"
            checked={isPromo}
            onChange={e => setIsPromo(e.target.checked)}
          />
          Акция
        </label>

        <input
          type="text"
          placeholder="URL изображения"
          value={imageUrl}
          onChange={e => setImageUrl(e.target.value)}
        />

        <button type="submit">Добавить товар</button>
      </form>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Магазин</th>
            <th>Продукт</th>
            <th>Цена</th>
            <th>Акция</th>
            <th>Изображение</th>
          </tr>
        </thead>
        <tbody>
          {storeProducts.map(sp => (
            <tr key={sp.id}>
              <td>{sp.id}</td>
              <td>{sp.store_id}</td>
              <td>{sp.product_id}</td>
              <td>{sp.price}</td>
              <td>{sp.is_promo ? "Да" : "Нет"}</td>
              <td>
                {sp.image_url ? (
                  <img
                    src={sp.image_url}
                    alt="product"
                    style={{
                      width: "60px",
                      height: "60px",
                      objectFit: "cover"
                    }}
                  />
                ) : (
                  "-"
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}