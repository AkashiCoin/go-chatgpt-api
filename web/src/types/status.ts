export type StatusState = {
    status: object | undefined;
}

export type StatusAction = {
    type: number;
    payload: object | undefined;
}

export enum StatusActionTypes {
    SET,
    UNSET
}