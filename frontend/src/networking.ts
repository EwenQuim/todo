import useSWR from "swr";

export function useRequest<T>(key: string) {
  const { data, error, mutate, isValidating } = useSWR<T>(key, apiFetcher);
  return {
    data,
    isLoading: !error && !data,
    error,
    mutate,
    isValidating,
  };
}

type Options = {
  method?: string;
  body?: Record<string, unknown>;
  headers?: Record<string, string>;
  queryParams?: Record<string, string | number | boolean>;
};

export async function apiFetcher<T = unknown>(
  path: string,
  { method, body, queryParams, headers }: Options | undefined = {}
): Promise<T> {
  const fullURL = new URL(
    path,
    process.env.NODE_ENV === "development"
      ? "http://localhost:8084"
      : document.baseURI
  );

  if (queryParams) {
    Object.keys(queryParams).forEach((key) => {
      fullURL.searchParams.append(key, "" + queryParams[key]);
    });
  }

  try {
    const resp = await fetch(fullURL.toString(), {
      method: method ?? "GET",
      body: body ? JSON.stringify(body) : null,
      headers: new Headers({ "Content-Type": "application/json", ...headers }),
      // credentials: "include", // Send cookies along with request,
    });
    return await resp.json();
  } catch (e) {
    console.log("err", e);
  }

  return {} as T;
}
