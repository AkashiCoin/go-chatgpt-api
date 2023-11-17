
export let base_path: string = ""
export const setBasePath = (path: string) => {
    base_path = path
    if (!base_path.startsWith("/")) {
        base_path = "/" + base_path
    }
    if (base_path.endsWith("/")) {
        base_path = base_path.slice(0, -1)
    }
}
if (window.WEBAPP.base_path) {
    setBasePath(window.WEBAPP.base_path)
}

export let api : string= import.meta.env.VITE_API_URL as string
if (window.WEBAPP.api) {
    api = window.WEBAPP.api
}
if (api === "/") {
    api = location.origin + base_path
}
if (api.endsWith("/")) {
    api = api.slice(0, -1)
}
