import './App.css';
import './index.css'

import React, { useEffect, useState } from 'react';

import { Todo } from './types';
import TodoList from './components/Todo';

const App = () => {
  const [input, setInput] = useState("")
  const [newTodoPrivate, setNewTodoPrivate] = useState(false)
  const [todos, setTodos] = useState<Todo[]>([])

  const getTodos = () => {
    fetch('/api/todo')
      .then(response => response.json())
      .then(data => {
        console.log(data)
        setTodos(data)
      })
  }


  const newTodo = async () => {
    if (input !== "") {
      setInput("")
      const response = await fetch('/api/todo/new?title=' + input + '&public=' + !newTodoPrivate)
      const jsonResponse: Todo = await response.json()
      console.log(jsonResponse.UUID)
      // move to another page
      window.location.href = '/' + jsonResponse.UUID
      getTodos()

    }
  };

  useEffect(() => {
    getTodos()
  }, [])

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
    <div className="App">

      <h1>TODO APP 📝</h1>
      <h2>Create a new todo list !</h2>

      <label htmlFor="newTodoPrivate">Title </label>
      <input
        id="todoName"
        name="todoName"
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            newTodo();
          }
        }}
      />
      <button className="m-2 p-2 rounded bg-purple-600" onClick={newTodo}>
        Create
      </button>
      <br />

      <label htmlFor="newTodoPrivate">🔐 Secret list </label>
      <input type="checkbox" name="newTodoPrivate" id="newTodoPrivate" checked={newTodoPrivate} onClick={() => setNewTodoPrivate(!newTodoPrivate)} />
      <br />




      <h2>Public todo lists</h2>

      <ul>
        {todos.map(todo => {
          return <li> <a href={todo.UUID}> {todo.Title} </a></li>
        })}
      </ul>

    </div>
  );
}

export default App;
