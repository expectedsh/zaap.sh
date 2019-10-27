export interface State {
  token: string | undefined
}

export interface SetTokenAction {
  type: "SET_TOKEN"
  token: string | undefined
}

export type Action = SetTokenAction | { type: "LOL" }
