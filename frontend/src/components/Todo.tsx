import React, { useEffect, useState } from 'react';
import { Item, Todo } from '../types';
import ReactModal from 'react-modal';
import { tryToFetch } from '../utils/utils';
import { ItemView } from './ItemView';

const TodoList = ({ uuid }: { uuid: string }) => {
  const [input, setInput] = useState("")
  const [todo, setTodo] = useState<Todo>({ Title: "", UUID: "", Items: [], Public: false })
  const [items, setItems] = useState<Item[]>([])

  const [online, setOnline] = useState(true)

  const getTodo = () => {

    fetch('/api/todo/' + uuid)
      .then(response => response.json())
      .then(data => {
        setTodo(data)
        console.log("set items", data.Items)
        if (data?.Items) {
          setItems(data.Items)
        }
      })
  }

  useEffect(() => {
    getTodo()
  }, [])

  useEffect(() => {
    document.title = todo.Title
    console.log("updated todo :", todo)
  }, [todo])



  const newItem = async () => {
    if (input !== "") {
      setInput("");

      // Change locally (for an impression of speed)
      // Create if empty
      setItems([...items, {
        ID: -1,
        Content: input,
        Done: false
      }])

      // Then, try to sync with the server
      try {
        const response = await fetch('/api/todo/' + uuid + "/new?content=" + input)
        const responseJson = await response.json()
        setItems(items => items.map(item => item.ID === -1 ? { ...item, ID: responseJson.ID } : item))
      } catch {
        setOnline(false)
      }
    }
  };

  const switchItem = async (item: Item) => {
    setItems(items => items.map(i => i.ID === item.ID ? { ...i, Done: !i.Done } : i))
    tryToFetch('/api/todo/' + uuid + "/" + item.ID + "/switch", setOnline)
  }

  const deleteItem = (item: Item) => {
    setItems(items => items.filter(i => i.ID !== item.ID))
    tryToFetch('/api/todo/' + uuid + "/" + item.ID + "/delete", setOnline)
  }

  return (
    <>
      <div className="todo-list">

        {!todo?.Items?.length && !todo?.Public &&
          <>⬆ <em> Bookmark this URL so you can find it later (only you will be able to access it !)</em></>
        }

        <h1>{todo?.Title} {todo && !todo?.Public && "🔐"}</h1>

        <div >
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                newItem();
              }
            }}
          />
          <button className="btn btn-primary" onClick={newItem}>
            Add
          </button>
        </div>

        <div>
          <ul>
            {items.map((item) => (
              <li>
                <ItemView item={item} deleteItem={deleteItem} switchItem={switchItem} />
              </li>
            ))}
          </ul>
        </div>

      </div>
      <ReactModal isOpen={!online} onRequestClose={() => setOnline(true)} contentLabel="Offline">
        <p> ❌ Offline, changes are not synchronized</p>
      </ReactModal>
    </>
  );
};



export default TodoList;
