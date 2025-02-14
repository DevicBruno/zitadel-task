import React, { useEffect } from 'react';
import authConfig from '../authConfig';
import EndpointButtons from './EndpointButtons';

const Callback = ({ auth, setAuth, userManager, userInfo, setUserInfo, handleLogout }) => {

  useEffect(() => {
    if (auth === null) {
      userManager.signinRedirectCallback().then((user) => {
        if (user) {
          setAuth(true);
          const access_token = user.access_token;
          // Make a request to the user info endpoint using the access token
          fetch(authConfig.userinfo_endpoint, {
            headers: {
              'Authorization': `Bearer ${access_token}`
            }
          })
            .then(response => response.json())
            .then(userInfo => {
              setUserInfo({ ...userInfo, access_token });
            });
        } else {
          setAuth(false);
        }
      }).catch((error) => {
        setAuth(false);
      });
    }
  }, [auth, userManager, setAuth]);

  if (auth === true && userInfo) {
    return (
      <div>
        <h1>Welcome, {userInfo.name}!</h1>
        <h2>Your ZITADEL Profile Information</h2>
        <h3>Name:  {userInfo.name}</h3>
        <h3>Email: {userInfo.email}</h3>
        <h3>Email Verified: {userInfo.email_verified ? "Yes" : "No"}</h3>
        <h3>Locale: {userInfo.locale}</h3>
        <button onClick={handleLogout}>Log out</button>
        <hr style={{ margin: '30px 0' }} />
        <EndpointButtons userInfo={userInfo} />
      </div>
    );
  }
  else {
    return <div>Loading...</div>;
  }
};

export default Callback;
