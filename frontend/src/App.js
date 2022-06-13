import './App.css';
import Signin from './pages/Signin';
import Signup from './pages/Signup';
import TotalProducts from './pages/TotalProducts';
import {Routes, Route} from "react-router-dom";
import Form from './components/Form';
import Form2 from './components/Form2';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<Signin/>}/>
        <Route path="/signup" element={<Signup />} />
        <Route path="/product" element={<TotalProducts />}/>
        <Route path="/form" element={<Form/>}/>
        <Route path="/add" element={<Form2/>}/>
      </Routes>
    </div>
  );
}

export default App;
