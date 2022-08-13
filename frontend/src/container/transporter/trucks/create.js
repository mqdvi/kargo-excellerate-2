import React, { Component } from 'react'
import TrucksService from '../../../services/trucks-service';

export class CreateTransporterTrucks extends Component {
    constructor(props) {
        super(props)

        this.state = {
            trucksType: {},
            'licenseNumber': '',
            'plateType': '',
            'truckTypeId': '',
            'productionYear': '',
            'stnk': '',
            'kir': ''
        }
    }

    componentDidMount = async() => {
        const items = localStorage.getItem('user')
        if (items !== "Transporter") {
            window.open("/", "_self");
        }

        await TrucksService.getTrucksType()
        .then((res) => {
            this.setState({
                trucksType: res.data
            });
        });
    }

    handlerChange = (e) => {
        this.setState({ [e.target.name]: e.target.value })
    }

    onFileChange = (e) => {
	    this.setState({ stnk: e.target.files[0] });
        this.setState({ kir: e.target.files[0] });
	};

    handlerSubmit = async (event) => {
        event.preventDefault()
        
        const formData = {};
        formData["licenseNumber"] = this.state.licenseNumber;
        formData["plateType"] = this.state.plateType;
        formData["truckTypeId"] = this.state.truckTypeId;
        formData["productionYear"] = this.state.productionYear;
        // formData["stnk"] = null;
        // formData["kir"] = null;

        console.log(formData)
       await TrucksService.createTrucks(formData);
        window.open("/Transporter/trucks", "_self")
    }

    render() {
        const { data } = this.state.trucksType
        const { items } = data || {}

        return (
            <div className="background p-10">
                <h1 className="header-background text-xl">CREATE TRUCKS</h1>

                <form class="rounded p-10" onSubmit={this.handlerSubmit}>
                    <div class="mb-6s">
                        <label class="block text-sm font-header mb-2 uppercase" for="licenseNumber">License Number</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="licenseNumber" type="text" placeholder="licenseNumber" onChange = {this.handlerChange} required/>
                    </div>
                    <div class="inline-block relative w-64 mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="plateType">plateType</label>
                        <select class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="plateType" placeholder="plateType" onChange = {this.handlerChange} required>
                            <option name="plateType">Black</option>
                            <option name="plateType">Yellow</option>
                        </select>
                    </div>
                    <div class="inline-block relative w-64 mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="truckTypeId">truckType</label>
                        <select class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="truckTypeId" placeholder="truckTypeId" onChange = {this.handlerChange} required>
                            <option name="tronton">Tronton</option>
                            <option name="broton">Broton</option>

                        {(items || []).map((item, index) => (
                            <option name="truckTypeId">{item.id}</option>
                        ))}
                        </select>
                    </div>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="productionYear">productionYear</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="productionYear" type="number" placeholder="productionYear" onChange = {this.handlerChange} required/>
                    </div>

                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="stnk">stnk</label>
                        <input className="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline bg-white" name="foto" type="file" accept='.png, .jpg, .jpeg' onChange = {this.onFileChange} />
                    </div>
                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="kir">kir</label>
                        <input className="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline bg-white" name="kir" type="file" accept='.png, .jpg, .jpeg' onChange = {this.onFileChange} />
                    </div>
                    

                    <div class="mb-6">
                        <input 
                            type = 'submit'
                            value = 'Save Unit'
                            className='bg-yellow-100 px-5 py-2'
                        />
                    </div>

                    
                </form>
            </div>
        )
    }
}

export default CreateTransporterTrucks