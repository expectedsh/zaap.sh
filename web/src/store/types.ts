import { ThunkAction } from "redux-thunk"
import { Action as UserAction, State as UserState } from "./user/types"

export interface State {
  user: UserState
}

export type Action = UserAction

export type ThunkResult<R> = ThunkAction<R, State, undefined, Action>
