import React, { Component } from 'react';
import axios from 'axios';
 
class App extends Component {
  constructor(props){
              super(props);
              this.state = {
                     data: [],
                     isLoading: false,
               };
               this.click = this.click.bind(this);
        }

    click() {
        this.setState({ isLoading: true });

        axios.get(`/welcome`, {})
            .then((response) => {
                  console.log("response: " + response)
                  this.setState({ data: response.data, isLoading: false });
             })
            .catch((err) => {
                  this.setState({ data: err, isLoading: false });
             });
    }

    render() {
       return  (
            <div>
                <button onClick={this.click} disabled={this.state.isLoading}> click me </button>
                {this.state.data[0]}
            </div>
           );
    }
}
 
export default App;
