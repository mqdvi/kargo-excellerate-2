import React, { Component } from 'react'
import ArtikelService from '../../services/artikel-service';

export class CreateArtikel extends Component {
    constructor(props) {
        super(props)

        this.state = {
            'authorName': '',
            'title': '',
            'body': '',
            'createdAt': '',
            'updatedAt': ''
        }
    }

    handlerChange = (e) => {
        this.setState({ [e.target.name]: e.target.value })
    }

    onFileChange = (e) => {
	    this.setState({ title: e.target.files[0] });
	};

    handlerSubmit = async (event) => {
        event.preventDefault()
        
        const formData = new FormData();
        formData.append("authorName", this.state.authorName);
        formData.append("title", this.state.title);
        formData.append("body", this.state.body);
        formData.append("createdAt", this.state.createdAt);
        formData.append("updatedAt", this.state.updatedAt);

        console.log(this.state.authorName)
            
        await ArtikelService.createArtikel(formData);
        window.open("/read-artikel", "_self")
    }

    render() {
        return (
            <div className="background p-10">
                <h1 className="h1 text-2xl py-5">CREATE ARTIKEL PAGE</h1>

                <form className="rounded p-10" onSubmit={this.handlerSubmit}>
                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="authorName">authorName</label>
                        <input className="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="authorName" type="text" placeholder="authorName" onChange = {this.handlerChange} required/>
                    </div>
                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="title">title</label>
                        <input className="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="title" type="text" placeholder="title" onChange = {this.handlerChange} required/>
                    </div>
                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="body">body</label>
                        <textarea className="shadow appearance-none border rounded-lg w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline h-auto" rows="5" name="body" type="text" placeholder="body" onChange = {this.handlerChange} required/>
                    </div>
                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="createdAt">createdAt</label>
                        <input className="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="createdAt" type="date" placeholder="createdAt" onChange = {this.handlerChange} required/>
                    </div>
                    <div className="mb-6">
                        <label className="block text-sm font-header mb-2 uppercase" for="updatedAt">updatedAt</label>
                        <input className="shadow appearance-none border rounded-full w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" name="updatedAt" type="date" placeholder="updatedAt" onChange = {this.handlerChange} required/>
                    </div>
                    <input 
                        type = 'submit'
                        value = 'submit'
                        className='bg-yellow-100 p-5'
                    />
                </form>
            </div>
        )
    }
}

export default CreateArtikel