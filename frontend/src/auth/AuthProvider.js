import React, { createContext, useContext, useState, useEffect } from 'react';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [accessToken, setAccessToken] = useState(null);

  // Handle successful login
  const handleLoginSuccess = (token) => {
    setAccessToken(token);
    setIsAuthenticated(true);
  };

  return (
    <AuthContext.Provider value={{ 
      isAuthenticated, 
      accessToken,
      handleLoginSuccess,
    }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
