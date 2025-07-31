import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from './Header';
import { useNavigate } from 'react-router-dom';
import '../App.css';
import { FiShoppingBag, FiPlus, FiStar } from 'react-icons/fi';

function ItemList() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const navigate = useNavigate();

  const getAuthHeaders = () => {
    const token = localStorage.getItem('token');
    return {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };
  };

  useEffect(() => {
    const fetchItems = async () => {
      try {
        setLoading(true);
        const response = await axios.get('http://localhost:8080/items', getAuthHeaders());
        setItems(response.data.items);
      } catch (error) {
        console.error('Failed to fetch items:', error.response ? error.response.data : error.message);
        if (error.response?.status === 401) {
          localStorage.removeItem('token');
          navigate('/login');
        }
      } finally {
        setLoading(false);
      }
    };
    fetchItems();
  }, [navigate]);

  const handleAddToCart = async (itemId, itemName) => {
    try {
      await axios.post('http://localhost:8080/carts', { item_id: itemId }, getAuthHeaders());
      showNotification('success', `"${itemName}" added to your cart!`);
    } catch (error) {
      console.error('Failed to add item to cart:', error.response ? error.response.data : error.message);
      if (error.response?.status === 401) {
        localStorage.removeItem('token');
        navigate('/login');
      } else {
        showNotification('error', 'Failed to add item to cart. Please try again.');
      }
    }
  };

  const showNotification = (type, message) => {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;
    notification.textContent = message;
    document.body.appendChild(notification);
    
    setTimeout(() => {
      notification.classList.add('fade-out');
      setTimeout(() => notification.remove(), 500);
    }, 3000);
  };

  const filteredItems = items.filter(item =>
    item.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="app-container">
      <Header />
      
      <main className="main-content">
        <div className="hero-section">
          <h1>Discover Amazing Products</h1>
          <p>Shop the latest trends and exclusive items</p>
          
          <div className="search-container">
            <input
              type="text"
              placeholder="Search products..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
            <button className="search-btn">Search</button>
          </div>
        </div>
        
        {loading ? (
          <div className="loading-state">
            <div className="spinner"></div>
            <p>Loading our finest products...</p>
          </div>
        ) : filteredItems.length === 0 ? (
          <div className="empty-state">
            <FiShoppingBag size={48} />
            <p>{searchTerm ? 'No matching products found' : 'No products available at the moment'}</p>
            {searchTerm && (
              <button onClick={() => setSearchTerm('')} className="clear-search">
                Clear search
              </button>
            )}
          </div>
        ) : (
          <>
            <div className="product-filters">
              <div className="filter-buttons">
                <button className="active">All Items</button>
                <button>Popular</button>
                <button>New Arrivals</button>
                <button>On Sale</button>
              </div>
              <div className="sort-options">
                <span>Sort by:</span>
                <select>
                  <option>Featured</option>
                  <option>Price: Low to High</option>
                  <option>Price: High to Low</option>
                  <option>Customer Rating</option>
                </select>
              </div>
            </div>
            
            <div className="product-grid">
              {filteredItems.map((item) => (
                <div key={item.id} className="product-card">
                  <div className="product-badge">New</div>
                  <div className="product-image">
                    <div className="image-placeholder">
                      {item.name.charAt(0).toUpperCase()}
                    </div>
                  </div>
                  <div className="product-info">
                    <h3 className="product-name">{item.name}</h3>
                    <div className="product-rating">
                      <FiStar className="star filled" />
                      <FiStar className="star filled" />
                      <FiStar className="star filled" />
                      <FiStar className="star filled" />
                      <FiStar className="star" />
                      <span>(24)</span>
                    </div>
                    <p className="product-status">{item.status}</p>
                    <div className="product-footer">
                      <span className="product-price">$29.99</span>
                      <button 
                        className="add-to-cart-btn"
                        onClick={() => handleAddToCart(item.id, item.name)}
                      >
                        <FiPlus className="icon" />
                        Add to Cart
                      </button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </>
        )}
      </main>
      
      <footer className="app-footer">
        <div className="footer-content">
          <div className="footer-section">
            <h3>ShopEase</h3>
            <p>Your premium shopping destination for quality products and exceptional service.</p>
          </div>
          <div className="footer-section">
            <h4>Quick Links</h4>
            <ul>
              <li>Home</li>
              <li>Products</li>
              <li>About Us</li>
              <li>Contact</li>
            </ul>
          </div>
          <div className="footer-section">
            <h4>Customer Service</h4>
            <ul>
              <li>My Account</li>
              <li>Order Tracking</li>
              <li>Returns</li>
              <li>FAQs</li>
            </ul>
          </div>
        </div>
        <div className="footer-bottom">
          <p>&copy; {new Date().getFullYear()} ShopEase. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}

export default ItemList;