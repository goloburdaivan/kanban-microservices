import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from '../homepage_components/App.jsx'
import '../homepage_components/index.css'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
