import React, { useRef, useState } from "react";
import { useSWRConfig } from "swr";
import { useRequest } from "../networking";
import { Item, Todo } from "../types";
import { capitalizeFirstLetter, detectRegex, tryToFetch } from "../utils/utils";
import { ItemView } from "./ItemView";

const TodoList = ({ uuid }: { uuid: string }) => {
  const [input, setInput] = useState("");

  const [online, setOnline] = useState(true);

  const { data: todo } = useRequest<Todo>(`/api/todo/${uuid}`);
  const { mutate } = useSWRConfig();

  const newItem = async () => {
    if (input.length === 0) {
      return;
    }
    setInput("");

    // Then, try to sync with the server
    try {
      await fetch("/api/todo/" + uuid + "/new?content=" + input);
      mutate(`/api/todo/${uuid}`);
    } catch {
      setOnline(false);
    }
  };

  const switchItem = async (item: Item) => {
    await tryToFetch("/api/todo/" + uuid + "/switch/" + item.ID, setOnline);
    mutate(`/api/todo/${uuid}`);
  };

  const deleteItem = async (item: Item) => {
    await tryToFetch("/api/todo/" + uuid + "/delete/" + item.ID, setOnline);
    mutate(`/api/todo/${uuid}`);
  };

  const deleteItems = () => {
    for (const item of todo?.Items ?? []) {
      if (item.Done) {
        deleteItem(item);
      }
    }
  };

  const sortFunction = (a: Item, b: Item) => {
    if (a.Content.includes(":") === b.Content.includes(":")) {
      if (detectRegex(a.Content) === detectRegex(b.Content)) {
        return a.ID > b.ID ? 1 : -1;
      }
      return a.Content.toLowerCase() > b.Content.toLowerCase() ? 1 : -1;
    }
    return a.Content.includes(":") ? 1 : -1;
  };

  const detected: string[] = [];
  const searchInput = useRef<HTMLInputElement>(null);

  return (
    <>
      <div className="todo-list">
        {!online && (
          <div className="fixed z-10 top-0 right-0 left-0 text-center bg-red-500 text-white">
            Offline
          </div>
        )}

        {!todo?.Items?.length && todo?.Public === false && (
          <>
            ⬆{" "}
            <em>
              Bookmark this URL so you can find it later (only you will be able
              to access it !)
            </em>
          </>
        )}

        <h1>{todo?.Title ?? "Todo"}</h1>

        {todo?.Public === false && (
          <>
            <em>Secret list</em> 🔐
          </>
        )}

        <div className="flex">
          <input
            ref={searchInput}
            autoFocus={true}
            className="flex-1 my-2 text-input"
            placeholder="Add an item"
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                newItem();
              }
            }}
          />
          <button
            className="my-2 ml-2 px-2 rounded w-8 text-white"
            onClick={newItem}
          >
            +
          </button>
        </div>

        <div>
          <ul className="relative">
            {todo?.Items.sort(sortFunction).map((item) => {
              const res = detectRegex(item.Content).toLowerCase();
              if (res && !detected.includes(res)) {
                detected.push(res);
                return (
                  <>
                    <div className="flex sticky top-0 mt-6 -mx-4 backdrop-blur-sm bg-gray-50/90 dark:bg-gray-800/90 border-t border-gray-200">
                      <h3 className="flex-1 pl-6">
                        {capitalizeFirstLetter(res)}{" "}
                      </h3>
                      <button
                        className="px-2 m-2 mr-6 rounded w-8 text-white align-middle"
                        onClick={() => {
                          setInput(res.toLowerCase() + ": ");
                          searchInput.current?.focus();
                        }}
                      >
                        +
                      </button>
                    </div>

                    <li key={item.ID}>
                      <ItemView item={item} switchItem={switchItem} />
                    </li>
                  </>
                );
              }

              return (
                <li>
                  <ItemView item={item} switchItem={switchItem} />
                </li>
              );
            })}
          </ul>
        </div>

        <button className="my-6 px-4 rounded" onClick={deleteItems}>
          🧹 Clean up
        </button>
      </div>
    </>
  );
};

export default TodoList;
