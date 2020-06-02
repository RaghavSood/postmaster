import React from 'react';
import { Button } from 'semantic-ui-react'

const PageButtons = ({prevCallback, nextCallback}) => (
   <div> 
  <Button 
    primary
    onClick={prevCallback}
  >Prev</Button>
  <Button 
    primary
    onClick={nextCallback}
  >Next</Button>
    </div>
)

export default PageButtons
