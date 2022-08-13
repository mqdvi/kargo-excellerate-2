import React, { Component } from 'react'
import DriverService from "../../../services/driver-service";

export class CreateTransporterDrivers extends Component {
    constructor(props) {
        super(props)

        this.state = {
            'phoneNumber': '',
            'driverName':'',
        }
    }

    componentDidMount() {
        const items = localStorage.getItem('user')
        if (items !== "Transporter") {
            window.open("/", "_self");
        }
    }

    handlerChange = (e) => {
        this.setState({ [e.target.name]: e.target.value })
    }

    handlerSubmit = async (event) => {
        event.preventDefault()
        
        const formData = new FormData();
        formData.append("phoneNumber", this.state.phoneNumber);
        formData.append("driverName", this.state.driverName);
            
        await DriverService.createDriver(formData);
        window.open("/Transporter/drivers", "_self")
    }

    render() {
        return (
            <div className="background p-10">
                <h1 className="header-background text-xl">ADD NEW DRIVERS</h1>

                <form class="rounded p-10" onSubmit={this.handlerSubmit}>
                    <div class="mb-6s">
                        <label class="block text-sm font-header mb-2 uppercase" for="driverName">Driver Name</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="driverName" type="text" placeholder="driverName" onChange = {this.handlerChange} required/>
                    </div>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="phoneNumber">Phone Number</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="phoneNumber" type="number" placeholder="phoneNumber" onChange = {this.handlerChange} required/>
                    </div>
                    

                    <div class="mb-6">
                        <input 
                            type = 'submit'
                            value = 'Save Drivers'
                            className='bg-yellow-100 px-5 py-2'
                        />
                    </div>

                    
                </form>
            </div>
        )
    }
}

export default CreateTransporterDrivers