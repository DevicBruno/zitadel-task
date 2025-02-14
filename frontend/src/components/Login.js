import { Navigate } from "react-router-dom";
import EndpointButtons from './EndpointButtons';

const Login = ({ auth, handleLogin }) => {
  return (
    <div>
      {auth === null && <div>Loading...</div>}
      {auth === false && (
        <div>
          <h1>Welcome!</h1>
          <button
            onClick={() => {
              handleLogin();
            }}
          >
            Please log in or register.
          </button>
        </div>
      )}
      {auth && <Navigate to="/callback" />}
      <hr style={{ margin: '30px 0' }} />
      <EndpointButtons userInfo={null} />
    </div>
  );
};

export default Login;
