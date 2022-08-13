import React, { Component } from 'react'

export class UpdateTransporter extends Component {
    constructor(props) {
        super(props)

        this.state = {
            'id': '',
            'licenseNumber': '',
            'isi': '',
            'productionYear': '',
            'licenceType': ''
        }
    }

    // componentDidMount = async () => {
    //     this.isLogin();
    //     // const id = this.props.match.params.id
    //     const data = window.location.href.split('/')
    //     const panjang = data.length
    //     const id = data[panjang-1]
    //     const res = await ArtikelService.getArtikelById('get-artikel.php?id=' + id)

    //     this.setState({
    //         id: res.data.id,
    //         judul: res.data.judul,
    //         foto: res.data.foto,
    //         isi: res.data.isi,
    //         productionYear: res.data.productionYear,
    //         licenceType: res.data.licenceType
    //     })
    // }

    handlerChange = (e) => {
        this.setState({ [e.target.name]: e.target.value })
    }

    handlerSubmit = async (event) => {
        event.preventDefault()
        
        const formData = new FormData();
        formData.append("id", this.state.id);
        formData.append("licenseNumber", this.state.licenseNumber);
        formData.append("isi", this.state.isi);
        formData.append("productionYear", this.state.productionYear);
        formData.append("licenceType", this.state.licenceType);
            
        // await ArtikelService.createArtikel('update-artikel.php', formData);
    	
        window.open("/Transporter/trucks", "_self")
    }

    render() {
        const { licenseNumber, truckTypeId, productionYear, licenceType } = this.state

        return (
            <div className="background p-10">
                <h1 className="header-background text-xl">ARTIKEL ADMIN PAGE</h1>

                <form class="rounded p-10" onSubmit={this.handlerSubmit}>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="licenseNumber">License Number</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="licenseNumber" type="text" placeholder="licenseNumber" value={licenseNumber} onChange = {this.handlerChange} required/>
                    </div>
                    <div class="inline-block relative w-64 mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="licenceType">licenceType</label>
                        <select class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="licenceType" placeholder="licenceType" value={licenceType} onChange = {this.handlerChange} required>
                            <option name="black">Black</option>
                            <option name="yellow">Yellow</option>
                        </select>
                    </div>
                    <div class="inline-block relative w-64 mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="truckTypeId">truckType</label>
                        <select class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="truckTypeId" placeholder="truckTypeId" value={truckTypeId} onChange = {this.handlerChange} required>
                            <option name="tronton">Tronton</option>
                            <option name="tronton">Tronton</option>
                        </select>
                    </div>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="productionYear">productionYear</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="productionYear" type="number" placeholder="productionYear" value={productionYear} onChange = {this.handlerChange} required/>
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

export default UpdateTransporter;