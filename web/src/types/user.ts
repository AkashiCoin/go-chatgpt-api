export type User =  {
    id: number;
    username: string;
    email: string;
    token: string;
    created_at: string;
    updated_at: string;
    deleted_at: string | null;
}

export type UserState = {
    user: User | undefined;
}

export type UserAction = {
    type: number;
    payload: User | undefined;
}

export enum UserActionTypes  {
    LOGIN,
    LOGOUT
}