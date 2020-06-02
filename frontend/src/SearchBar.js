import React from 'react';
import { Input } from 'semantic-ui-react'

const SearchBar = ({filterCallback}) => (
    <Input icon='search' placeholder='Email'
    onChange={filterCallback} />
)

export default SearchBar
