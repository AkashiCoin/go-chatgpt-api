import {StatusAction, StatusActionTypes, StatusState} from "@/types/status.ts";
import {Reducer} from "react";

export const reducer: Reducer<StatusState, StatusAction> = (state: StatusState, action: StatusAction) => {
  switch (action.type) {
    case StatusActionTypes.SET:
      return {
        ...state,
        status: action.payload,
      };
    case StatusActionTypes.UNSET:
      return {
        ...state,
        status: undefined,
      };
    default:
      return state;
  }
};

export const initialState: StatusState = {
  status: undefined,
};
