import React, { useEffect, useState } from "react";
import { Todo } from "../types";
import { apiFetcher, useRequest } from "../networking";

export const Todos = () => {
  const [input, setInput] = useState("");
  const [newTodoPrivate, setNewTodoPrivate] = useState(false);

  const { data: todos, error, mutate } = useRequest<Todo[]>("/api/todo");

  const newTodo = async () => {
    if (input !== "") {
      setInput("");
      const response = await apiFetcher<{ UUID: string }>("/api/todo", {
        method: "POST",
        queryParams: {
          title: input,
          public: !newTodoPrivate,
        },
      });

      // move to another page
      window.location.href = "/" + response.UUID;
      mutate();
    }
  };

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
        {todos?.map((todo) => {
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
