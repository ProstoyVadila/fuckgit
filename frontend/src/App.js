import './App.css';
import React from 'react';
import Body from './components/body';
import Header from './components/header';
import Footer from './components/footer';
import About from './components/about';


function App() {
  return (
    <div className="App">
      <Header />
      <Body />
      <About />
      <Footer />
    </div>
  );
}

export default App;
