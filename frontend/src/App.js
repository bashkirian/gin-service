import React, { useState } from 'react';
import Header from './Header';
import Maps from './Maps';
import SearchBox from './SearchBox';

import './App.css';

function App() {
  const [selectPosition, setSelectPosition] = useState(null);

  // console.log(selectPosition);
  return (
    <div className='App'>
      <div className='App-header'>
        <Header />
      </div>
      <div className='App-map'>
        <Maps selectPosition={selectPosition} />
      </div>
      <div className='App-search'>
        <SearchBox selectPosition={selectPosition} setSelectPosition={setSelectPosition} />
      </div>
    </div>
  );
}

export default App;
