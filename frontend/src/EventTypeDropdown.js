import React from 'react';
import { Dropdown } from 'semantic-ui-react'

const options = [
    { key: 1, text: 'Bounce', value: 'Bounce' },
    { key: 2, text: 'Complaint', value: 'Complaint' },
    { key: 3, text: 'Delivery', value: 'Delivery' },
    { key: 4, text: 'Send', value: 'Send' },
    { key: 5, text: 'Reject', value: 'Reject' },
    { key: 6, text: 'Open', value: 'Open' },
    { key: 7, text: 'Click', value: 'Click' },
    { key: 8, text: 'Failure', value: 'Failure' },
]

const EventTypeDropdown = ({filterCallback}) => (
  <Dropdown 
    clearable 
    options={options} 
    selection
    onChange={filterCallback} />
)

export default EventTypeDropdown
