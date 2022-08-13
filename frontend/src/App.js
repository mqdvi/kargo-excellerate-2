import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import './App.css';
import Example from './container/index';
import Navbar from './components/navbar/index';

import Login from './container/login/index';
import ReadTransporter from './container/transporter/read';
import ReadDetailTransporter from './container/transporter/read-detail';
import UpdateTransporter from './container/transporter/update';
import CreateTransporter from './container/transporter/create';

function App() {
    return (
        <>
            <Navbar />
            <BrowserRouter>
                <Routes>
                    <Route path="/" exact element={<Example />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/Transporter/trucks" element={<ReadTransporter />} />
                    <Route path="/Transporter/trucks/:id" element={<ReadDetailTransporter />} />
                    <Route path="/Transporter/trucks/create" element={<CreateTransporter />} />
                    <Route path="/Transporter/trucks/update/:id" element={<UpdateTransporter />} />
                </Routes>
            </BrowserRouter>
        </>
    );
}

export default App;