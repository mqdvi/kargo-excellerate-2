import React, { Component } from 'react'

export class UpdateTransporterDrivers extends Component {
    constructor(props) {
        super(props)

        this.state = {
            'id': '',
            'name': '',
            'phoneNumber': '',
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
    //         phoneNumber: res.data.phoneNumber,
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
        formData.append("name", this.state.name);
        formData.append("phoneNumber", this.state.phoneNumber);
            
        // await ArtikelService.createArtikel('update-artikel.php', formData);
    	
        window.open("/Transporter/drivers", "_self")
    }

    render() {
        const { name, phoneNumber } = this.state

        return (
            <div className="background p-10">
                <h1 className="header-background text-xl">UPDATE DRIVER DATA</h1>

                <form class="rounded p-10" onSubmit={this.handlerSubmit}>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="name">License Number</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="name" type="text" placeholder="name" value={name} onChange = {this.handlerChange} required/>
                    </div>
                    <div class="mb-6">
                        <label class="block text-sm font-header mb-2 uppercase" for="phoneNumber">phoneNumber</label>
                        <input class="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="phoneNumber" type="number" placeholder="phoneNumber" value={phoneNumber} onChange = {this.handlerChange} required/>
                    </div>
                    

                    <div class="mb-6">
                        <input 
                            type = 'submit'
                            value = 'Update Data Drivers'
                            className='bg-yellow-100 px-5 py-2'
                        />
                    </div>

                    
                </form>
            </div>
        )
    }
}

export default UpdateTransporterDrivers;