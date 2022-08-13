import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import { BiEdit } from 'react-icons/bi'
import TrucksService from '../../../services/trucks-service';

class ReadTransporterTrucks extends Component {
    state = {
        trucks: {}
    }

    componentDidMount = async () =>{
        const items = localStorage.getItem('user')
        if (items !== "Transporter") {
            window.open("/", "_self");
        }
        await TrucksService.getTrucks()
        .then((res) => {
            this.setState({
                trucks: res.data
            });
        });
    }

    render() {        
        const { data } = this.state.trucks
        const { items } = data || {}
        
        return (
            <>
                <div className="background p-10">
                    <h1 className="h1 text-2xl py-5">TRANSPORTER PAGE</h1>
                    <div className="create-button py-5">
                        <Link className="bg-yellow-100 px-10 py-2" to='/Transporter/trucks/create'>CREATE TRUCKS</Link>
                    </div>
                    <table className='table-auto'>
                        <thead className='border'>
                            <tr>
                                <th className='px-4 py-5'>License Number</th>
                                <th className='px-4 py-5'>Truck Type</th>
                                <th className='px-4 py-5'>Plat Type</th>
                                <th className='px-4 py-5'>Production Year</th>
                                <th className='px-4 py-5'>Status</th>
                                <th className='px-4 py-5'>Pengaturan</th>

                            </tr>
                        </thead>
                        <tbody className='border'>
                            {(items || []).map((item, index) => (
                                <tr>
                                    <td className='px-4'>
                                        <Link className='px-2' to={"/Transporter/trucks/" + item.id}>{item.licenseNumber}</Link>
                                    </td>
                                    <td className='px-4'>{item.truckType}</td>
                                    <td className='px-4'>{item.plateType}</td>
                                    <td className='px-4'>{item.productionYear}</td>
                                    <td className='px-4'>{item.status}</td>
                                    <td className="flex">
                                        <Link className='px-2' to={"/Transporter/trucks/update" + item.id}><BiEdit /></Link>
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

export default ReadTransporterTrucks