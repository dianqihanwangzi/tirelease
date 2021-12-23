import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
// import App from './App';
import reportWebVitals from './reportWebVitals';
// import Dashboard from './dashboard/Dashboard'
import Dashboard from './components/dashboard-layout'

ReactDOM.render(
  // <React.StrictMode>
  //   <App />
  // </React.StrictMode>,
  <React.StrictMode>
    <Dashboard />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
