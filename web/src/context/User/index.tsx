// contexts/User/index.jsx

import React, {createContext, ReactNode, useReducer} from "react"
import { reducer, initialState } from "./reducer"
import {UserAction} from "@/types/user.tsx";
import {UserState} from "@/types/user.ts";

type UserProviderProps = {
    children: ReactNode;
};

type UserContextType = {
    state: UserState;
    dispatch: React.Dispatch<UserAction>;
};

export const UserContext = createContext<UserContextType>({
    state: initialState,
    dispatch: (() => null) as React.Dispatch<UserAction>
})

export const UserProvider: React.FC<UserProviderProps> = ({ children }) => {
    const [state, dispatch] = useReducer(reducer, initialState)

    return (
        <UserContext.Provider value={{ state, dispatch }}>
            { children }
        </UserContext.Provider>
    )
}