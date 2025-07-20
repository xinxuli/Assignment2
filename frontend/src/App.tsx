import { useState } from 'react'
import './App.css'
import { Api } from './shared'

function App() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [token, setToken] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isRegistering, setIsRegistering] = useState(false)

  return (
    < div className="userlogin-app">
      <h1>User Login</h1>
      {}
      <div className='login-form'>
        <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
        <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
        <button 
        className='login-button'
        onClick={() => Api.userLoginPost({username, password}).then (
          (response) => {
            setToken(response.token)
            setIsLoggedIn(true)
          }
        ).catch((err) => { alert(err.message)})
        }
        >Login</button>
      </div>
    </div>
  )
}

export default App
