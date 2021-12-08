
export type Todo = {
    UUID: string;
    Title: string;
    Public: boolean;
    Items: Item[];
}

export type Item = {
    ID: number;
    Content: string;
    Done: boolean;
}
