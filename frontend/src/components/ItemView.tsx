import { Item } from "../types"
import React from 'react';

export const ItemView = ({ item, switchItem, deleteItem }: { item: Item, switchItem: any, deleteItem: any }) => {

  return (
    <span>
      <input type="checkbox" name="toggle-item" id={`item-${item.ID}`} onClick={async () => switchItem(item)} checked={item.Done} />
      <label htmlFor={`item-${item.ID}`}>
        {item.Done ? <s>{item.Content}</s> : item.Content}
      </label>

      {" "}
      <button onClick={async () => deleteItem(item)}>
        x
      </button>
    </span>
  )
}
