import React from 'react';
import {Input, Button,Label,Dropdown } from 'semantic-ui-react'


const options = [
    { key: 1, text: 'Bounce', value: 'Bounce' },
    { key: 2, text: 'Complaint', value: 'Complaint' },
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
        var searchQuery = this.state.searchQuery;
	fetch(`http://localhost:8080/api/suppression/check?email=${searchQuery}`)
            .then((response) => response.json())
	    .then((response) => alert(response.email_address))
      }
      

    handleAddButtonClicked() {
        var searchQuery = this.state.searchQueryAdd;
        var reason = this.state.reason;
        alert('Added: ' + searchQuery + 'because: ' + reason)
        console.log(searchQuery)
    }

    handleDeleteButtonClicked() {
        var searchQuery = this.state.searchQuery;
        fetch(`http://localhost:8080/api/suppression/delete?email=${searchQuery}`)
            .then((response) => response.json())
	    .then((response) => alert(response.response))
      }
    

    render(){
        return (
            <div>
                <div>
                    <h3> Suppression List Form </h3>
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

