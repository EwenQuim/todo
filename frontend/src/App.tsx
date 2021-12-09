import './App.css';
import './index.css'

import React, { useEffect, useState } from 'react';

import { Todo } from './types';
import TodoList from './components/Todo';
import { Todos } from './components/Todos';

const App = () => {


  if (window.location.pathname !== '/') {
    const path = window.location.pathname.split('/')[1]
    const regex = new RegExp(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$/i)

    if (regex.test(path)) {
      return <TodoList uuid={path} />
    } else {
      return <div>404</div>
    }
  }

  return (
    <Todos />
  );
}

export default App;
