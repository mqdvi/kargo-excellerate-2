import React, { Component } from 'react'

export class CreateTransporter extends Component {
    constructor(props) {
        super(props)

        this.state = {
            'licenseNumber': '',
            'licenceType': '',
            'truckTypeId': '',
            'productionYear': '',
        }
    }

    handlerChange = (e) => {
        this.setState({ [e.target.name]: e.target.value })
    }

    handlerSubmit = async (event) => {
        event.preventDefault()
        
        const formData = new FormData();
        formData.append("licenseNumber", this.state.licenseNumber);
        formData.append("licenceType", this.state.licenceType);
        formData.append("truckTypeId", this.state.truckTypeId);
        formData.append("productionYear", this.state.productionYear);
            
        // await ArtikelService.createArtikel(formData);
        window.open("/Transporter/trucks", "_self")
    }

    render() {
        return (
            <div className="background p-10">
                <h1 className="header-background text-xl">ARTIKEL ADMIN PAGE</h1>

                <form class="rounded p-10" onSubmit={this.handlerSubmit}>
                    <div class="mb-6s">
                        <label class="block text-sm font-header mb-2 uppercase" for="licenseNumber">License Number</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="licenseNumber" type="text" placeholder="licenseNumber" onChange = {this.handlerChange} required/>
                    </div>
                    <div class="inline-block relative w-64 mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="licenceType">licenceType</label>
                        <select class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="licenceType" placeholder="licenceType" onChange = {this.handlerChange} required>
                            <option name="black">Black</option>
                            <option name="yellow">Yellow</option>
                        </select>
                    </div>
                    <div class="inline-block relative w-64 mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="truckTypeId">truckType</label>
                        <select class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="truckTypeId" placeholder="truckTypeId" onChange = {this.handlerChange} required>
                            <option name="tronton">Tronton</option>
                            <option name="tronton">Tronton</option>
                        </select>
                    </div>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="productionYear">productionYear</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="productionYear" type="number" placeholder="productionYear" onChange = {this.handlerChange} required/>
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

export default CreateTransporter