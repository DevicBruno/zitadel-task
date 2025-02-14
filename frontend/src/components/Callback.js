import React, { useEffect } from 'react';
import authConfig from '../authConfig';
import EndpointButtons from './EndpointButtons';

const Callback = ({ auth, setAuth, userManager, userInfo, setUserInfo, handleLogout }) => {

  useEffect(() => {
    const getExistingUser = async () => {
      const user = await userManager.getUser();
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
      }
    };

    if (auth === null) {
      // First try to complete the redirect callback
      userManager.signinRedirectCallback().then((user) => {
        if (user) {
          setAuth(true);
          const access_token = user.access_token;
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
      }).catch(() => {
        // If callback fails, check for existing session
        getExistingUser();
      });
    } else if (!userInfo) {
      // If we have auth but no userInfo (like after refresh), get the user
      getExistingUser();
    }
  }, [auth, userManager, setAuth, setUserInfo]);

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
