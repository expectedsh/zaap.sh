import { State, Action } from "./types"

const initialState: State = {
  token: undefined,
}

export default (state = initialState, action: Action): State => {
  switch (action.type) {
  case "SET_TOKEN":
    return {
      ...state,
      token: action.token,
    }
  }
  return state
}
