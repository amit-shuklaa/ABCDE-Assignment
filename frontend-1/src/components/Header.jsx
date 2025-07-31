import React from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../App.css';
import { FiShoppingCart, FiClock, FiLogOut, FiCheckCircle } from 'react-icons/fi';

function Header() {
  const navigate = useNavigate();

  const getAuthHeaders = () => {
    const token = localStorage.getItem('token');
    return {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };
  };

  const handleApiCall = async (apiCallFunc) => {
    try {
      await apiCallFunc();
    } catch (error) {
      console.error('API call failed:', error.response ? error.response.data : error.message);
      if (error.response && error.response.status === 401) {
        localStorage.removeItem('token');
        navigate('/login');
      }
    }
  };

  const handleCheckout = async () => {
    await handleApiCall(async () => {
      await axios.post('http://localhost:8080/orders', {}, getAuthHeaders());
      // Using a more modern notification
      showNotification('success', 'Order placed successfully!');
    });
  };

  const handleViewCart = async () => {
    await handleApiCall(async () => {
      const response = await axios.get('http://localhost:8080/carts', getAuthHeaders());
      const cartData = response.data.cart;

      if (cartData?.items?.length > 0) {
        const cartDetails = cartData.items
          .map(item => `• ${item.item_name || 'Item'} (ID: ${item.item_id})`)
          .join('\n');
        showNotification('info', `Your Cart:\n${cartDetails}`);
      } else {
        showNotification('info', 'Your cart is empty');
      }
    });
  };

  const handleViewOrderHistory = async () => {
    await handleApiCall(async () => {
      const response = await axios.get('http://localhost:8080/orders', getAuthHeaders());
      const orderIDs = response.data.order_ids;

      if (orderIDs?.length > 0) {
        showNotification('info', `Your Orders:\n${orderIDs.map(id => `• Order #${id}`).join('\n')}`);
      } else {
        showNotification('info', 'No orders placed yet');
      }
    });
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
    showNotification('success', 'Logged out successfully!');
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

  return (
    <header className="app-header">
      <div className="logo">ShopEasy</div>
      <nav className="nav-actions">
        <button className="nav-btn checkout-btn" onClick={handleCheckout}>
          <FiCheckCircle className="icon" />
          <span>Checkout</span>
        </button>
        <button className="nav-btn cart-btn" onClick={handleViewCart}>
          <FiShoppingCart className="icon" />
          <span>Cart</span>
        </button>
        <button className="nav-btn orders-btn" onClick={handleViewOrderHistory}>
          <FiClock className="icon" />
          <span>Orders</span>
        </button>
        <button className="nav-btn logout-btn" onClick={handleLogout}>
          <FiLogOut className="icon" />
          <span>Logout</span>
        </button>
      </nav>
    </header>
  );
}

export default Header;