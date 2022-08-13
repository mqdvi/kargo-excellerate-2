import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import { BiEdit } from 'react-icons/bi'
import DriverService from "../../../services/driver-service";


class ReadTransporterDrivers extends Component {
    state = {
        drivers: {}
    }

    componentDidMount = async () =>{
        const items = localStorage.getItem('user')
        if (items !== "Transporter") {
            window.open("/", "_self");
        }
 
        await DriverService.getDriver()
        .then((res) => {
            this.setState({
                drivers: res.data
            });
        });
    }

    render() {        
        const { data } = this.state.drivers
        const { items } = data || {}
        
        return (
            <>
                <div className="background p-10">
                    <h1 className="h1 text-2xl py-5">DRIVERS PAGE</h1>
                    <div className="create-button py-5">
                        <Link className="bg-yellow-100 px-10 py-2" to='/Transporter/drivers/create'>CREATE DRIVERS</Link>
                    </div>
                    <table className='table-auto'>
                        <thead className='border'>
                            <tr>
                                <th className='px-4 py-5'>Driver Name</th>
                                <th className='px-4 py-5'>Phone Number</th>
                                <th className='px-4 py-5'>Created At</th>
                                <th className='px-4 py-5'>Status</th>
                                <th className='px-4 py-5'>Pengaturan</th>

                            </tr>
                        </thead>
                        <tbody className='border'>
                            {(items || []).map((item, index) => (
                                <tr>
                                    <td className='px-4'>
                                        <Link className='px-2' to={"/Transporter/drivers/" + item.id}>{item.name}</Link>
                                    </td>
                                    <td className='px-4'>{item.phoneNumber}</td>
                                    <td className='px-4'>{item.createdAt}</td>
                                    <td className='px-4'>{item.status}</td>
                                    <td className="flex">
                                        <Link className='px-2' to={"/Transporter/drivers/update" + item.id}><BiEdit /></Link>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>            
            </>
        )
    }
}

export default ReadTransporterDrivers
