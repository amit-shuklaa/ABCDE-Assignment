import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';       // We'll create this file soon
import ItemList from './components/ItemList'; // We'll create this file soon

function App() {
  // Basic authentication check: returns true if a token exists in localStorage
  const isAuthenticated = () => {
    return localStorage.getItem('token') !== null;
  };

  return (
    <Router>
      <div className="App">
        <Routes>
          {/* Public routes */}
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Login />} /> {/* Default route to Login screen */}

          {/* Protected route - requires authentication */}
          <Route
            path="/items"
            element={isAuthenticated() ? <ItemList /> : <Navigate to="/login" replace />}
          />

          {/* Catch-all route for undefined paths, redirects to Login */}
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;