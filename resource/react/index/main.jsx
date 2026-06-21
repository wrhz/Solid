import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'

const container = document.getElementById('app');
if (!container) {
  throw new Error('cannot find #app view');
}
const root = ReactDOM.createRoot(container);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);