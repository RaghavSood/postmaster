import React from 'react';
import { Label } from 'semantic-ui-react';
import EventsTable from './EventsTable';
import EventTypeDropdown from './EventTypeDropdown';
import PageButtons from './PageButtons';
import SearchBar from './SearchBar';
import SupessionForm from './SupressionForm';

export default class ConnectedEventsTable extends React.Component {
    state = {
        results: [],
        queryParams: {
            event_filter: '',
            email_filter: '',
            message_id: '',
            from: 0,
            direction: "next",
        },
    };

    componentDidMount() {
        this.updateResults()
    }

    render() {
        return (
            <div>
                <div className="ui container">
                    <SupessionForm/>
                </div>
                <br />
                <div className="ui container">
                    <EventTypeDropdown filterCallback={this.handleEventTypeFilter} />
                    <SearchBar filterCallback={this.handleSearchBar} />
                </div>
                <br />
                <EventsTable results={this.state.results} historyBtn={ true }/>
                <div className="ui container">
                <PageButtons prevCallback={this.prevButtonClick} nextCallback={this.nextButtonClick} />
                </div>
            </div>
        );
    }

    prevButtonClick = (e) => {
        let qp = this.state.queryParams;
        if (this.state.results.length > 0) {
            qp.from = this.state.results[0].id;
            qp.direction = "prev";
        } else if (qp.direction === "next") {
            qp.direction = "last"
        }
        this.setState({queryParams: qp}, () => this.updateResults());
    }

    nextButtonClick = (e) => {
        let qp = this.state.queryParams;
        if (this.state.results.length > 0) {
            qp.from = this.state.results[this.state.results.length -1].id;
            qp.direction = "next";
        } else if (qp.direction === "prev") {
            qp.direction = "first"
        }
        this.setState({queryParams: qp}, () => this.updateResults());
    }

    handleSearchBar = (searchBar, values) => {
        let qp = this.state.queryParams;
        qp.email_filter = values.value;
        this.setState({queryParams: qp}, () => this.updateResults());
    }

    handleEventTypeFilter = (eventFilter, values) => {
        let qp = this.state.queryParams;
        qp.event_filter = values.value;
        this.setState({queryParams: qp}, () => this.updateResults());
    }

    updateResults() {
        fetch(`/api/events?event_filter=${this.state.queryParams.event_filter}&email_filter=${this.state.queryParams.email_filter}&from=${this.state.queryParams.from}&direction=${this.state.queryParams.direction}`)
            .then((response) => response.json())
            .then((response) => this.setState(response))
    }
}
