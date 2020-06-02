import React from 'react';
import { Button, Header, Modal } from 'semantic-ui-react'
import CodeHighlight from './CodeHighlight'

export default function EventDetail({eventData}) {
    console.log(eventData)

    return (
         <Modal trigger={<Button primary>Event Data</Button>}>
           <Modal.Header>Event Details</Modal.Header>
           <Modal.Content scrolling>
             <Modal.Description>
               <Header>Message ID</Header>
               <p>{ eventData.messageId }</p>

               <Header>Event</Header>
               <p>{ eventData.eventType }</p>

               <Header>Subject</Header>
               <p>{ eventData.mail.commonHeaders.subject }</p>

               <Header>Recipients</Header>
               <p>{ eventData.recipients }</p>

               <Header>Received At (UTC)</Header>
               <p>{ eventData.received_at }</p>

               <Header>SNS ID</Header>
               <p>{ eventData.sns_id }</p>

               <Header>Mail</Header>
               <CodeHighlight language="json">{ JSON.stringify(eventData.mail, undefined, 4) }</CodeHighlight>

               <Header>Event Data</Header>
               <CodeHighlight language="json">{ JSON.stringify(eventData.event_data, undefined, 4) }</CodeHighlight>

             </Modal.Description>
           </Modal.Content>
         </Modal> 
    );
};
