

export const tryToFetch = async (url: string, setOnline: any) => {
  try {
    await fetch(url)
  } catch {
    setOnline(false)
  }
}
