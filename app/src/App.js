import React, { Component } from 'react';
import axios from 'axios';
import { config } from "process";
 
class App extends Component {
  constructor(props){
              super(props);
              this.state = {
                     data: [],
		     inputValue: '',
                     isLoading: false,
               };
               this.add = this.add.bind(this);
	       this.show = this.show.bind(this);
               this.handleChange = this.handleChange.bind(this);
        }

    handleChange(event) {
      this.setState({inputValue: event.target.value});
    }

    async add(e) {
	    e.preventDefault();
            this.setState({ data: [] });
	    this.setState({ isLoading: true });

	    const headers = {
		'Content-Type':'application/json'
	    };

	    await axios.post(
		"welcome", 
		 {
		    text: this.state.inputValue,
		},
		{headers}
		).then(response => {
		    console.log("Success ========>", response);
		})
		.catch(error => {
		    console.log("Error ========>", error);
		}
	    )
            this.setState({ inputValue: '' });
	    this.setState({ isLoading: false });
}

    show() {
        this.setState({ isLoading: true });

        axios.get("/welcome", {})
            .then((response) => {
                  console.log("response: " + response)
                  this.setState({ data: response.data, isLoading: false });
		  <p>{this.state.data}</p>
             })
            .catch((err) => {
                  this.setState({ data: err, isLoading: false });
             });
    }

    render() {
       return  (
            <div>
                <form>
		  <input type="text" value={this.state.inputValue} onChange={this.handleChange} />
                  <button onClick={this.add} disabled={this.state.isLoading}> Add </button>
	          <button onClick={this.show} disabled={this.state.isLoading}> Show all </button>
                </form>
            </div>
           );
    }
}
 
export default App;
