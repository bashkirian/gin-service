import React, { useState } from 'react';
import Maps from './Maps';
import SearchBox from './SearchBox';

import './App.css';

function App() {
  const [selectPosition, setSelectPosition] = useState(null);
  const [listPlace, setListPlace] = React.useState([]);

  return (
    <div className='App'>
      <div className='App-map'>
        <Maps selectPosition={selectPosition} listPlace={listPlace} />
      </div>
      <div className='App-search'>
        <SearchBox selectPosition={selectPosition} setSelectPosition={setSelectPosition}
         listPlace={listPlace} setListPlace={setListPlace} />
      </div>
    </div>
  );
}

export default App;
