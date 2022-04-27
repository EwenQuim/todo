import React, { useEffect, useState } from "react";
import { Todo } from "../types";

export const Todos = () => {
  const [input, setInput] = useState("");
  const [newTodoPrivate, setNewTodoPrivate] = useState(false);
  const [todos, setTodos] = useState<Todo[]>([]);

  const getTodos = () => {
    fetch("/api/todo")
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setTodos(data);
      });
  };

  const newTodo = async () => {
    if (input !== "") {
      setInput("");
      const response = await fetch(
        "/api/todo/new?title=" + input + "&public=" + !newTodoPrivate
      );
      const jsonResponse: Todo = await response.json();
      console.log(jsonResponse.UUID);
      // move to another page
      window.location.href = "/" + jsonResponse.UUID;
      getTodos();
    }
  };

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <div className="App">
      <h1>TODO APP 📝</h1>
      <h2>Create a new todo list !</h2>

      <div className="flex">
        <input
          id="todoName"
          name="todoName"
          placeholder="Title"
          className="flex-1 my-2 text-input"
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              newTodo();
            }
          }}
        />
        <button
          className="m-2 p-2 rounded bg-purple-600 text-white"
          onClick={newTodo}
        >
          Create
        </button>
      </div>
      <label htmlFor="newTodoPrivate">🔐 Secret list </label>
      <input
        type="checkbox"
        name="newTodoPrivate"
        id="newTodoPrivate"
        checked={newTodoPrivate}
        onClick={() => setNewTodoPrivate(!newTodoPrivate)}
      />
      <br />

      <h2>Public todo lists</h2>

      <ul>
        {todos.map((todo) => {
          return (
            <li>
              {" "}
              <a href={todo.UUID}> {todo.Title} </a>
            </li>
          );
        })}
      </ul>
    </div>
  );
};
