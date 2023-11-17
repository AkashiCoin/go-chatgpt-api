import {UserState, UserAction, UserActionTypes} from "@/types/user.ts";
import {Reducer} from "react";

export const reducer: Reducer<UserState, UserAction> = (state: UserState, action: UserAction): UserState => {
  switch (action.type) {
    case UserActionTypes.LOGIN:
      return {
        ...state,
        user: action.payload
      };
    case UserActionTypes.LOGOUT:
      return {
        ...state,
        user: undefined
      };

    default:
      return state;
  }
};

export const initialState: UserState = {
  user: undefined
};