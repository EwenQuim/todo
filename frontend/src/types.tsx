export type Todo = {
  UUID: string;
  Title: string;
  Public: boolean;
  Items: Item[];
  Groups: Group[];
};

export type Group = {
  Name: string;
  Items: Item[];
};

export type Item = {
  ID: number;
  Content: string;
  Done: boolean;
};
