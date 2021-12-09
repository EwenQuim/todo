

export const tryToFetch = async (url: string, setOnline: any) => {
  try {
    await fetch(url)
  } catch {
    setOnline(false)
  }
}

export const detectRegex = (s: string): string => s.match(/([\w ]+)?: *(.*)/)?.[1] ?? ""

export const capitalizeFirstLetter = (string: string) => string.charAt(0).toUpperCase() + string.slice(1)
