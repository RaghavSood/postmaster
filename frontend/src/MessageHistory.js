import React from 'react';
import EventsTable from './EventsTable';
import { Button, Header, Modal } from 'semantic-ui-react'

export default class MessageHistory extends React.Component {
    state = {
        results: [],
        queryParams: {
            event_filter: '',
            email_filter: '',
            message_id: this.props.message,
            from: 0,
            direction: "next",
        }
    };

    componentDidMount() {
        this.updateResults(this.props.message)
    }

    componentWillReceiveProps(props) {
        this.updateResults(props.message)
    }


    render() {
        return (
         <Modal trigger={<Button primary>Email History</Button>}>
           <Modal.Header>Message History</Modal.Header>
           <Modal.Content scrolling>
             <Modal.Description>
                <EventsTable results={this.state.results} historyBtn={ false }/>
             </Modal.Description>
           </Modal.Content>
         </Modal>
        );
    }

    updateResults(message_id) {
        fetch(`/api/message?message_id=${message_id}&direction=${this.state.queryParams.direction}`)
            .then((response) => response.json())
            .then((response) => this.setState(response))
    }
}
