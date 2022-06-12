import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Home from './pages/Home';
import Detail from './pages/Detail';
import Form from './pages/Form';


function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home></Home>}></Route>
          <Route path="/contacts/:id" element={<Detail />}></Route>
          <Route path="/addcontact" element={<Form />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
