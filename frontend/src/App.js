import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import './App.css';
import Home from './container/index';
import Navbar from './components/navbar/index';

import Login from './container/login/index';

import ReadTransporterTrucks from './container/transporter/trucks/read';
import ReadDetailTransporterTrucks from './container/transporter/trucks/read-detail';
import UpdateTransporterTrucks from './container/transporter/trucks/update';
import CreateTransporterTrucks from './container/transporter/trucks/create';

import ReadTransporterDrivers from './container/transporter/drivers/read';
import ReadDetailTransporterDrivers from './container/transporter/drivers/read-detail';
import UpdateTransporterDrivers from './container/transporter/drivers/update';
import CreateTransporterDrivers from './container/transporter/drivers/create';

import ReadShipper from './container/shipper/read';

function App() {
    return (
        <>
            <Navbar />
            <BrowserRouter>
                <Routes>
                    <Route path="/" exact element={<Home />} />
                    <Route path="/login" element={<Login />} />

                    <Route path="/Shipper/" element={<ReadShipper />} />

                    <Route path="/Transporter/trucks" element={<ReadTransporterTrucks />} />
                    <Route path="/Transporter/trucks/:id" element={<ReadDetailTransporterTrucks />} />
                    <Route path="/Transporter/trucks/create" element={<CreateTransporterTrucks />} />
                    <Route path="/Transporter/trucks/update/:id" element={<UpdateTransporterTrucks />} />

                    <Route path="/Transporter/drivers" element={<ReadTransporterDrivers />} />
                    <Route path="/Transporter/drivers/:id" element={<ReadDetailTransporterDrivers />} />
                    <Route path="/Transporter/drivers/create" element={<CreateTransporterDrivers />} />
                    <Route path="/Transporter/drivers/update/:id" element={<UpdateTransporterDrivers />} />

                </Routes>
            </BrowserRouter>
        </>
    );
}

export default App;