import { Item } from "../types"
import React from 'react';

export const ItemView = ({ item, switchItem, deleteItem }: { item: Item, switchItem: any, deleteItem: any }) => {

  return (
    <div className="m-2 flex items-center">
      <input className="flex-none mr-2" type="checkbox" name="toggle-item" id={`item-${item.ID}`} onClick={async () => switchItem(item)} checked={item.Done} />
      <label className="flex-1 mt-0" htmlFor={`item-${item.ID}`}>
        {item.Done ? <s>{item.Content}</s> : item.Content}
      </label>

      {" "}
      <button className="flex-none p-1 rounded w-8 h-8" onClick={async () => deleteItem(item)}>
        🗑
      </button>
    </div>
  )
}
