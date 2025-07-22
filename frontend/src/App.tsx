import { useState } from 'react'
import './App.css'
import { getApi, initApi } from './shared'
import { tokenStore, userInfoStore } from './store';
import { useNavigate } from "react-router-dom";

function App() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const isLogin = () => {
    return !!tokenStore.token
  }
  let navigate = useNavigate();
  return (
    < div className="userlogin-app">
      <h1>User Login</h1>
      {}
      <div className='login-form'>
        <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
        <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
        <button 
          className='login-button'
          onClick={() => getApi().userLoginPost({ userLoginReq: { username, password } }).then (
          async (response) => {
            tokenStore.token = response.token
            console.log(response.token)
            initApi(response.token)

            const resp = await getApi().userInfoGet()
          
            userInfoStore.nickname =resp.nickName

            navigate("/hello");

          }
          ).catch((err) => { alert(err.message)})
          }
        >Login</button>
      </div>
    </div>
  )
}

export default App
