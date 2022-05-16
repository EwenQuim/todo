import useSWR from "swr"

export function useRequest<T>(key: string) {
	const { data, error } = useSWR<T>(key, apiFetcher)
	return {
		data,
		isLoading: !error && !data,
		error: error
	}
}

export async function apiFetcher<T=unknown>(url:string, method?: string, body?:unknown) : Promise<T> {

	try {
		const resp = await fetch(url, {
			method: method??"GET",
			body: body ? JSON.stringify(body) : null,
			headers: new Headers({"content-type": "application/json"}),
			credentials: "include" // Send cookies along with request,
		})
		return await resp.json()
	} catch (e) {
		console.log("err", e)
	}

	return {} as T
}
