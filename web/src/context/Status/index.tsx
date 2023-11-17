// contexts/User/index.jsx

import React, {ReactNode} from 'react';
import { initialState, reducer } from './reducer';
import {StatusAction, StatusState} from "@/types/status.ts";

type StatusProviderProps = {
  children: ReactNode;
};

type StatusContextType = {
  state: StatusState;
  dispatch: React.Dispatch<StatusAction>;
};

export const StatusContext = React.createContext<StatusContextType>({
  state: initialState,
  dispatch: (() => null) as React.Dispatch<StatusAction>,
});

export const StatusProvider: React.FC<StatusProviderProps> = ({children }) => {
  const [state, dispatch] = React.useReducer(reducer, initialState);
  return (
    <StatusContext.Provider value={{state: state, dispatch: dispatch}}>
      {children}
    </StatusContext.Provider>
  );
};