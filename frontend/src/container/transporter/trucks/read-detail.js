import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import TrucksService from "../../../services/trucks-service";

export class ReadDetailTransporterTrucks extends Component {
    constructor(props) {
        super(props)

        this.state = {
            trucks:[]
        }
    }

    componentDidMount = async() => {
        const items = localStorage.getItem('user')
        if (items !== "Transporter") {
            window.open("/", "_self");
        }

        let id = window.location.href.split('/')[5]
        await TrucksService.getTrucksById("/"+parseInt(id))
        .then((res) => {
            this.setState({
                trucks: res.data
            });
        });
        console.log(this.state.trucks)

    }

    render() {
        const { data } = this.state.trucks
        const { id, licenseNumber, truckType, plateType, productionYear, status } = data || {}
        return (
            <div className="background">
                <Link to={"/Transporter/trucks"} className='submit-button text-lg py-10'>
                        <input type = "submit" value = "Back"/>
                    </Link>
                <div className="p-10">
                    <p className='text-lg'>ID {id}</p>
                    <p className='text-lg'>licenseNumber {licenseNumber}</p>
                    <p className='text-lg'>truckType {truckType}</p>
                    <p className='text-lg'>plateType {plateType}</p>
                    <p className='text-lg'>productionYear {productionYear}</p>
                    <p className='text-lg'>status {status}</p>
                </div>
            </div>
        )
    }
}

export default ReadDetailTransporterTrucks;