const authConfig = {
    authority: 'https://bruno-devic-interview-task-instance-ujvch7.us1.zitadel.cloud',
    client_id: '306808089301086298',
    redirect_uri: 'http://localhost:3000/callback',
    response_type: 'code',
    scope: 'openid profile email',
    post_logout_redirect_uri: 'http://localhost:3000',
    userinfo_endpoint: 'https://bruno-devic-interview-task-instance-ujvch7.us1.zitadel.cloud/oidc/v1/userinfo',
    response_mode: 'query',
    code_challenge_method: 'S256',
  };

 export default authConfig;
