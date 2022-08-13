import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import { BiEdit } from 'react-icons/bi'


class ReadShipperTrucks extends Component {
    state = {
        trucks: {}
    }

    componentDidMount = async () =>{
        const items = localStorage.getItem('user')
        if (items !== "Shipper") {
            window.open("/", "_self");
        }
        // await ArtikelService.getArtikel()
        // .then((res) => {
        //     this.setState({
        //         trucks: res.data
        //     });
        // });
    }

    render() {        
        const { data } = this.state.trucks
        const { items } = data || {}
        
        return (
            <>
                <div className="background p-10">
                    <h1 className="h1 text-2xl py-5">SHIPMENTS</h1>
                    <div className="create-button py-5 grid justify-items-end">
                        <Link className="bg-yellow-100 px-10 py-2" to='/Shipper/create'>ADD SHIPMENTS</Link>
                    </div>
                    <table className='table-auto'>
                        <thead className='border'>
                            <tr>
                                <th className='px-4 py-5'>Shipment</th>
                                <th className='px-4 py-5'>License</th>
                                <th className='px-4 py-5'>Driver's Name</th>
                                <th className='px-4 py-5'>Origin</th>
                                <th className='px-4 py-5'>Destinations</th>
                                <th className='px-4 py-5'>Loading Date</th>
                                <th className='px-4 py-5'>Status</th>
                                <th className='px-4 py-5'>Action</th>

                            </tr>
                        </thead>
                        <tbody className='border'>
                            {(items || []).map((item, index) => (
                                <tr>
                                    <td className='px-4'>{item.shipmentNumber}</td>
                                    <td className='px-4'>{item.licenseNumber}</td>
                                    <td className='px-4'>{item.driverName}</td>
                                    <td className='px-4'>{item.originDistrict}</td>
                                    <td className='px-4'>{item.destinationDisctrict}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>            
            </>
        )
    }
}

export default ReadShipperTrucks