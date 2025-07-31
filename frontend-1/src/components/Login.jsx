import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../App.css';
import { FiUser, FiLock, FiLogIn, FiShoppingBag } from 'react-icons/fi';


function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
      const response = await axios.post('http://localhost:8080/users/login', {
        username,
        password,
      });

      localStorage.setItem('token', response.data.token);
      navigate('/items');
      showNotification('success', 'Login successful! Welcome back!');

    } catch (error) {
      console.error('Login failed:', error.response ? error.response.data : error.message);
      showNotification('error', 'Invalid credentials. Please try again.');
    } finally {
      setLoading(false);
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

  return (
    <div className="login-container">
      <div className="login-background">
        <div className="login-overlay"></div>
      </div>
      
      <div className="login-card">
        <div className="login-brand">
          <FiShoppingBag className="brand-icon" />
          <h1>ShopEase</h1>
          <p>Your premium shopping destination</p>
        </div>
        
        <form onSubmit={handleLogin} className="login-form">
          <div className="form-group">
            <label htmlFor="username">
              <FiUser className="input-icon" />
              Username
            </label>
            <input
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              placeholder="Enter your username"
            />
          </div>
          
          <div className="form-group">
            <label htmlFor="password">
              <FiLock className="input-icon" />
              Password
            </label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              placeholder="Enter your password"
            />
          </div>
          
          <button type="submit" className="login-btn" disabled={loading}>
            {loading ? (
              <div className="spinner"></div>
            ) : (
              <>
                <FiLogIn className="icon" />
                Sign In
              </>
            )}
          </button>
        </form>
        
        <div className="login-footer">
          <p>New to ShopEase? <span>Create account</span></p>
          <p>Forgot password?</p>
        </div>
      </div>
    </div>
  );
}

export default Login;