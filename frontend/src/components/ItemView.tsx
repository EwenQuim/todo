import React, { useEffect } from 'react';

import { Item } from "../types"

export const ItemView = ({ item, switchItem, deleteItem }: { item: Item, switchItem: any, deleteItem: any }) => {

  const stringTransform = (str: string) => {
    const res = str.match(/([\w ]+)?: *(.*)/)
    if (res) {
      return res[2]
    }
    return str
  }

  return (
    <div className="m-2 flex items-center">
      <input className="flex-none mr-2" type="checkbox" name="toggle-item" id={`item-${item.ID}`} onClick={async () => switchItem(item)} checked={item.Done} />
      <label className="flex-1 mt-0" htmlFor={`item-${item.ID}`}>
        {item.Done ? <s>{stringTransform(item.Content)}</s> : stringTransform(item.Content)}
      </label>

      {" "}
      <button className="flex-none p-1 rounded w-8 h-8" onClick={async () => deleteItem(item)}>
        🗑
      </button>
    </div>
  )
}
