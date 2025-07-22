import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { userInfoStore } from './store';


import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { Nickname } from './components/Nickname.tsx';

let router = createBrowserRouter([
  {
    path: "/",
    element: <App />
  },
  {
    path: '/hello',
    element: <Nickname/>
  }
]);


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
