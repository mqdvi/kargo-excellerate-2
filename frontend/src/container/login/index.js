import React, { Component } from "react";

class Login extends Component {
    constructor() {
        super();
        this.state = {
            name: "React"
        };
        this.onValueChange = this.onValueChange.bind(this);
        this.formSubmit = this.formSubmit.bind(this);
    }

    onValueChange(event) {
        this.setState({
            user: event.target.value
        });
    }

    formSubmit(event) {
        event.preventDefault();
        window.open("/"+ (this.state.user).toString() + "/trucks", "_self")
    }

    render() {
        return (
            <div className="background p-10">
                <h1 className="h1 text-2xl py-5">LOGIN PAGE</h1>
                <form onSubmit={this.formSubmit}>
                    <div className="radio">
                        <label>
                            <input
                            type="radio"
                            value="Transporter"
                            checked={this.state.user === "Transporter"}
                            onChange={this.onValueChange}
                            />
                            Transporter
                        </label>
                    </div>
                    <div className="radio">
                        <label>
                            <input
                            type="radio"
                            value="Shipper"
                            checked={this.state.user === "Shipper"}
                            onChange={this.onValueChange}
                            />
                            Shipper
                        </label>
                    </div>
                    <button className="bg-yellow-100 px-5 py-2 m-3" type="submit">
                        Login
                    </button>
                </form>
            </div>
        );
    }
}

export default Login;