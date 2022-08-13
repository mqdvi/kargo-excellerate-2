import React, { Component } from 'react'
import { Link } from 'react-router-dom';

export class ReadDetailTransporterDrivers extends Component {
    constructor(props) {
        super(props)

        this.state = {
            data:[]
        }
    }

    // componentDidMount() {
    //     const data = window.location.href.split('/')
    //     const panjang = data.length
    //     const id = data[panjang-1]
    //     ArtikelService.getArtikelById('get-artikel.php?id=' + id)
    //     .then((res) => {
    //         this.setState({
    //             data: res.data
    //         });
    //     });
    // }

    render() {
        const { id, name, phoneNumber, createdAt, status } = this.state.data
        return (
            <div className="background">
                <Link to={"/Transporter/trucks"} className='submit-button pl-16'>
                    <input type = "submit" value = "Back"/>
                </Link>
                <div className="lg:p-56 lg:pt-0">
                    <p className='text-lg'>{id}</p>
                    <p className='text-lg'>{name}</p>
                    <p className='text-lg'>{phoneNumber}</p>
                    <p className='text-lg'>{createdAt}</p>
                    <p className='text-lg'>{status}</p>
                </div>
            </div>
        )
    }
}

export default ReadDetailTransporterDrivers;