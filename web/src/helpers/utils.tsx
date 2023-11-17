import { toast } from "react-toastify";

export function getSystemName() {
    const system_name = localStorage.getItem('system_name');
    if (!system_name) return 'Vite React';
    return system_name;
}

export function getFooterHTML() {
    return localStorage.getItem('footer_html');
}


export function showSuccess(message: string) {
    toast.success(message)
}

export function showError(error: Error | string) {
    console.error(error)
    if (typeof error === 'string') {
        toast.error(error)
        return
    }
    toast.error(error.message)
}