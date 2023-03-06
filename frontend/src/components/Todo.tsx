import React, { useRef, useState } from "react";
import { apiFetcher, useRequest } from "../networking";
import { Item, Todo } from "../types";
import { capitalizeFirstLetter } from "../utils/utils";
import { ItemView } from "./ItemView";

const TodoList = ({ uuid }: { uuid: string }) => {
  const [input, setInput] = useState("");

  const [online, setOnline] = useState(true);

  const { data: todo, mutate } = useRequest<Todo>(`/api/todo/${uuid}`);

  const newItem = async () => {
    if (input.length === 0) {
      return;
    }

    console.log("Adding item", { content: input, todoID: uuid });

    // Then, try to sync with the server
    try {
      await apiFetcher("/api/todo/item", {
        method: "POST",

        body: { content: input, todoID: uuid },
      });
      mutate();
    } catch {
      setOnline(false);
    }
    setInput("");
  };

  const switchItem = async (item: Item) => {
    await apiFetcher("/api/todo/" + uuid + "/switch/" + item.ID);
    mutate();
  };

  const deleteItem = async (item: Item) => {
    await apiFetcher(`/api/todo/item/${item.ID}`, { method: "DELETE" });
  };

  const deleteItems = async () => {
    for (const item of todo?.Items ?? []) {
      if (item.Done) {
        await deleteItem(item);
      }
    }
    mutate();
  };

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

        <h1 className="text-center md:text-left">{todo?.Title ?? "Todo"}</h1>

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

        <ul>
          {todo?.Groups?.map((group) => (
            <li key={group.Name}>
              {group.Name && (
                <div className="flex sticky top-0 md:top-2 mt-6 -mx-4 backdrop-blur-sm bg-gray-50/90 dark:bg-gray-800/90 md:rounded-md">
                  <h3 className="flex-1 pl-6">
                    {capitalizeFirstLetter(group.Name)}{" "}
                  </h3>
                  <button
                    className="px-2 m-2 mr-6 rounded w-8 text-white align-middle"
                    onClick={() => {
                      setInput(group.Name.toLowerCase() + ": ");
                      searchInput.current?.focus();
                    }}
                  >
                    +
                  </button>
                </div>
              )}
              <ul>
                {group?.Items?.map((item) => (
                  <ItemView key={item.ID} item={item} switchItem={switchItem} />
                ))}
              </ul>
            </li>
          ))}
        </ul>

        <button className="my-6 px-4 rounded" onClick={deleteItems}>
          🧹 Clean up
        </button>
      </div>
    </>
  );
};

export default TodoList;
