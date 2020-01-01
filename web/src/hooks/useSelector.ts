import { useSelector, TypedUseSelectorHook } from "react-redux"
import { State } from "store/types"

export default useSelector as TypedUseSelectorHook<State>
