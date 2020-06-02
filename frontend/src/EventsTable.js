import React from 'react';
import { Table, Header, Segment, Label } from 'semantic-ui-react'
import EventDetail from './EventDetail';
import MessageHistory from './MessageHistory';

export default function EventsTable({results, historyBtn}) {
    console.log(results)

    const rows = results.map(((result, index) => {
        return (
            <Table.Row key={ index }>
                <Table.Cell>{ result.id }</Table.Cell>
                <Table.Cell>{ ( 'timestamp' in result.event_data ) ? result.event_data.timestamp : result.mail.timestamp }</Table.Cell>
                <Table.Cell>{ result.eventType }</Table.Cell>
                <Table.Cell>{ result.mail.commonHeaders.subject }</Table.Cell>
                <Table.Cell>{ result.recipients }</Table.Cell>
                <Table.Cell><EventDetail eventData={ result }/></Table.Cell>
                { historyBtn ? <Table.Cell><MessageHistory message={ result.messageId }/></Table.Cell> : '' }
            </Table.Row>
        );
    }));

    return (
        <div className="ui container">
            <Segment>
                <Header>Activity </Header>
                <Table striped>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>ID</Table.HeaderCell>
                            <Table.HeaderCell>Time (UTC)</Table.HeaderCell>
                            <Table.HeaderCell>Event</Table.HeaderCell>
                            <Table.HeaderCell>Subject</Table.HeaderCell>
                            <Table.HeaderCell>Recipients</Table.HeaderCell>
                            <Table.HeaderCell>Details</Table.HeaderCell>
                            { historyBtn ? <Table.HeaderCell>History</Table.HeaderCell> : '' }
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { rows }
                    </Table.Body>
                </Table>
            </Segment>
        </div>
    );
}
