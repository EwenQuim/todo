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
    <div className="m-2 flex items-center">
      <input
        className="flex-none mr-2"
        type="checkbox"
        name="toggle-item"
        id={`item-${item.ID}`}
        onClick={async () => switchItem(item)}
        checked={item.Done}
      />
      <label className="flex-1 mt-0 bg-white" htmlFor={`item-${item.ID}`}>
        {item.Done ? (
          <s>{stringTransform(text)}</s>
        ) : (
          <input
            type="text"
            name=""
            id=""
            className="bg-white border-none"
            value={stringTransform(text)}
            onChange={(e) => {
              setText(e.target.value);
              changeItem(item, e.target.value);
            }}
          />
        )}
      </label>
    </div>
  );
};
