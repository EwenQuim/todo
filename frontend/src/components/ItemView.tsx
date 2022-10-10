import React from "react";
import { Item } from "../types";
import { tryToFetch } from "../utils/utils";

const changeItem = async (item: Item, to: string) => {
  tryToFetch(`/api/item/${item.ID}/change?new_content=${to}`, {});
};

export const ItemView = ({
  item,
  switchItem,
}: {
  item: Item;
  switchItem: any;
}) => {
  const [text, setText] = React.useState(item.Content);

  const stringTransform = (str: string) => {
    const res = str.match(/([\w ]+)?: *(.*)/);
    if (res) {
      return res[2];
    }
    return str;
  };

  return (
    <div className="flex items-center">
      <label className="pl-2 pr-4 py-1 my-auto" htmlFor={`item-${item.ID}`}>
        <input
          className="flex-none"
          type="checkbox"
          name="toggle-item"
          id={`item-${item.ID}`}
          onClick={async () => switchItem(item)}
          checked={item.Done}
        />
      </label>
      <div className="flex-1 mt-0 bg-white dark:bg-black">
        {item.Done ? (
          <s>{stringTransform(text)}</s>
        ) : (
          <input
            type="text"
            name=""
            id=""
            className="bg-white border-none w-full dark:bg-black"
            value={stringTransform(text)}
            onChange={(e) => {
              setText(e.target.value);
              changeItem(item, e.target.value);
            }}
          />
        )}
      </div>
    </div>
  );
};
