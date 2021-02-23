import React from 'react';
import {Input, Button } from 'semantic-ui-react'

export default class SupessionForm extends React.Component{
    constructor(props){
        super(props)
        this.state = {
            searchQuery: ""
        }
    }

    handleInputChanged(event){
        this.setState({
            searchQuery:event.target.value
        });
    }

    handleCheckButtonClicked() {
        var searchQuery = this.state.searchQuery;
	fetch(`http://localhost:8080/api/suppression/check?email=${searchQuery}`)
            .then((response) => response.json())
	    .then((response) => alert(response))
      }
      

    handleAddButtonClicked() {
        var searchQuery = this.state.searchQuery;
        alert('Added: ' + searchQuery)
        console.log(searchQuery)
    }

    handleDeleteButtonClicked() {
        var searchQuery = this.state.searchQuery;
        alert('Deleted: ' + searchQuery)
        console.log(searchQuery)
    }  

    render(){
        return (
            <div>
                <div>
                    <h3> Suppression List Form </h3>
                </div>
            <Input type="text" placeholder = "Enter email" style={{paddingBottom :'15px', width: '400px'}} value={this.state.searchQuery} onChange={this.handleInputChanged.bind(this)}/>
            <div>
                <Button onClick={this.handleCheckButtonClicked.bind(this)} style={{padding :'15px'}}>
                    Check
                </Button>
                <Button onClick={this.handleAddButtonClicked.bind(this)} style={{padding :'15px'}}>
                    Add
                </Button>
                <Button onClick={this.handleDeleteButtonClicked.bind(this)} style={{padding :'15px'}}>
                    Delete
                </Button>
            </div>
            </div>
          );
    }

}

