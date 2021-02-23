import React from 'react';
import {Input, Button,Label,Dropdown } from 'semantic-ui-react'

    function handleErrors(response) {
        if (response.error != undefined) {
            throw Error(response.error);
        }
        return response;
    }


const options = [
    { key: 1, text: 'Bounce', value: 'BOUNCE' },
    { key: 2, text: 'Complaint', value: 'COMPLAINT' },
]

export default class SupessionForm extends React.Component{
    constructor(props){
        super(props)
        this.state = {
            searchQueryCheckDelete: "",
            searchQueryAdd: "",
            reason:""
        }
    }

    handleInputChanged(event){
        this.setState({
            searchQueryCheckDelete:event.target.value
        });
    }

    handleAddInputChanged(event){
        this.setState({
            searchQueryAdd:event.target.value
        });
    }

    handleReasonChanged = (event,data) => {
        this.setState({
            reason : data.value
        });
    };

    handleReasonChanged(event){
        this.setState({
            reason:event.target.value
        });
    }

    handleCheckButtonClicked() {
        var searchQuery = this.state.searchQueryCheckDelete;
	fetch(`/api/suppression/check?email=${searchQuery}`)
            .then((response) => response.json())
	    .then(handleErrors)
	    .then((response) => alert(response.results.result_message))
	    .catch((error) => alert(error))
      }
      

    handleAddButtonClicked() {
        var searchQuery = this.state.searchQueryAdd;
        var reason = this.state.reason;
        fetch(`/api/suppression/add?email=${searchQuery}&reason=${reason}`, {method: 'POST'})
            .then((response) => response.json())
	    .then(handleErrors)
	    .then((response) => alert(response.results.result_message))
	    .catch((error) => alert(error))
    }

    handleDeleteButtonClicked() {
        var searchQuery = this.state.searchQueryCheckDelete;
        fetch(`/api/suppression/delete?email=${searchQuery}`, {method: 'POST'})
            .then((response) => response.json())
	    .then(handleErrors)
	    .then((response) => alert(response.results.result_message))
	    .catch((error) => alert(error))
      }
    

    render(){
        return (
            <div>
                <div>
                    <h3>Suppression List Management</h3>
                </div>
            <Label>    
            <Input type="text" placeholder = "Enter email" style={{paddingBottom :'15px', width: '400px'}} value={this.state.searchQueryCheckDelete} onChange={this.handleInputChanged.bind(this)}/>
            <div>
                <Button onClick={this.handleCheckButtonClicked.bind(this)} style={{padding :'15px'}}>
                    Check
                </Button>
                <Button onClick={this.handleDeleteButtonClicked.bind(this)} style={{padding :'15px'}}>
                    Delete
                </Button>
            </div>
            </Label>
            <Label>
            <Input type="text" placeholder = "Enter email" style={{paddingBottom :'15px', width: '200px'}} value={this.state.searchQueryAdd} onChange={this.handleAddInputChanged.bind(this)}/>
            <Dropdown 
                options={options} 
                selection
                value = {this.state.reason}
                onChange={this.handleReasonChanged} />
            <div>
                <Button onClick={this.handleAddButtonClicked.bind(this)} style={{padding :'15px'}}>
                    Add
                </Button>
            </div>
            </Label>
            </div>
          );
    }

}

