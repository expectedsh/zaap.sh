import { useDispatch } from "react-redux"
import { ThunkDispatch } from "redux-thunk"
import { State, Action } from "store/types"

export default () => useDispatch<ThunkDispatch<State, any, Action>>()
