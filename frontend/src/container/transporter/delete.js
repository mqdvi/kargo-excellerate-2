import React, { Component } from 'react'
import { Link } from 'react-router-dom';
import ArtikelService from '../../services/artikel-service';

export class DeleteArtikel extends Component {
    constructor(props) {
        super(props)
        this.state = {
            'judul': '',
        }
    }

    componentDidMount = async () => {
        const data = window.location.href.split('/')
        const panjang = data.length
        const id = data[panjang-1]
        const res = await ArtikelService.getArtikelById(id)

        this.setState({
            id: res.data.id,
            judul: res.data.judul
        })
    }

    handlerSubmit = async (event) => {
        event.preventDefault()
        const data = window.location.href.split('/')
        const panjang = data.length
        const id = data[panjang-1]
        await ArtikelService.deleteArtikel(id)
        window.open("/read-artikel", "_self")
    }

    render() {
        return (
            <div className="background">
                <h1 className="header-background">HAPUS ARTIKEL</h1>
                <div className="card-delete">
                    <form onSubmit={this.handlerSubmit}>
                        <p className="pesan-delete pb-5">Apakah kamu yakin ingin menghapus artikel dengan judul "{this.state.judul}" ?</p>
                        <div className="delete-button">
                            <input type = "submit" value = "Delete"/>
                        </div>
                    </form>
                    <Link to={"/admin/read-artikel-admin"} className='cancel-button'>
                        <input type = "submit" value = "Cancel"/>
                    </Link>
                </div>
            </div>
        )
    }
}

export default DeleteArtikel;