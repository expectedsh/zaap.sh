import { client } from "utils/api"
import { ThunkResult } from "store/types"

export type AuthenticateRequest = { provider: "email", email: string }
  | { provider: "github" | "google", code: string }

export type AuthenticateResponse = { token: string }

export const authenticate = (payload: AuthenticateRequest):
  ThunkResult<Promise<string>> => {
  return async (dispatch) => {
    const res = await client.post<AuthenticateResponse>("/authenticate", payload)
    dispatch({ type: "SET_TOKEN", token: res.data.token })
    return res.data.token
  }
}
