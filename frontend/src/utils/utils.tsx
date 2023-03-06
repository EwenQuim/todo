const detectRegex = (s: string): string =>
  s.match(/^ *([\w ]+) *: *(.*) *$/)?.[1].trim() ?? "";

export const capitalizeFirstLetter = (string: string) =>
  string.charAt(0).toUpperCase() + string.slice(1);
