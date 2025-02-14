import React, { useState } from 'react';

const EndpointButtons = ({ userInfo }) => {
  const [publicResponse, setPublicResponse] = useState({ status: null, message: '' });
  const [protectedResponse, setProtectedResponse] = useState({ status: null, message: '' });

  const testPublicEndpoint = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/public');
      const text = await response.text();
      let data;
      try {
        data = JSON.parse(text);
      } catch (e) {
        data = { message: text };
      }
      
      setPublicResponse({
        status: response.status,
        message: data.message || response.statusText
      });
    } catch (error) {
      setPublicResponse({
        status: error.response?.status || 'Network Error',
        message: error.message
      });
    }
  };

  const testProtectedEndpoint = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/private', {
        headers: {
          'Authorization': `Bearer ${userInfo?.access_token}`
        }
      });
      
      const text = await response.text();
      let data;
      try {
        data = JSON.parse(text);
      } catch (e) {
        data = { message: text };
      }

      setProtectedResponse({
        status: response.status,
        message: data.message || response.statusText
      });
    } catch (error) {
      setProtectedResponse({
        status: error.response?.status || 'Network Error',
        message: error.message
      });
    }
  };

  const getStatusStyle = (status) => ({
    color: status === 200 ? '#4CAF50' : 
           status === 401 ? '#FF9800' : 
           status === 403 ? '#FF9800' : 
           '#f44336',
    fontWeight: 'bold'
  });

  const responseBoxStyle = {
    margin: '10px 0',
    padding: '15px',
    borderRadius: '5px',
    backgroundColor: '#f5f5f5',
    border: '1px solid #ddd'
  };

  const getStatusText = (status) => {
    switch (status) {
      case 200: return ' OK';
      case 401: return ' Unauthorized';
      case 403: return ' Forbidden';
      case 404: return ' Not Found';
      case 500: return ' Internal Server Error';
      default: return status === 'Network Error' ? '' : ` ${status}`;
    }
  };

  return (
    <div style={{ margin: '20px', textAlign: 'center' }}>
      <div style={{ marginBottom: '30px' }}>
        <button onClick={testPublicEndpoint}>
          Test Public Endpoint
        </button>
        <div style={responseBoxStyle}>
          <span style={getStatusStyle(publicResponse.status || 0)}>
            {publicResponse.status || '---'} 
            {publicResponse.status ? getStatusText(publicResponse.status) : ''}
          </span>
          <br />
          <span style={{ marginTop: '5px', display: 'block' }}>
            {publicResponse.message || 'Click button to test endpoint'}
          </span>
        </div>
      </div>

      <div>
        <button onClick={testProtectedEndpoint}>
          Test Protected Endpoint
        </button>
        <div style={responseBoxStyle}>
          <span style={getStatusStyle(protectedResponse.status || 0)}>
            {protectedResponse.status || '---'} 
            {protectedResponse.status ? getStatusText(protectedResponse.status) : ''}
          </span>
          <br />
          <span style={{ marginTop: '5px', display: 'block' }}>
            {protectedResponse.message || 'Click button to test endpoint'}
          </span>
        </div>
      </div>
    </div>
  );
};

export default EndpointButtons; 