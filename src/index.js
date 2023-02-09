import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Header from './components/header';
import Question from './components/question';
import Solution from './components/solution';

const router = createBrowserRouter([
  { path: '/', element: <App /> },
  { path: '/questions', element: <Question /> },
  { path: '/solutions', element: <Solution /> },
  { path: '/header', element: <Header /> },
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
