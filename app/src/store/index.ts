import { createStore, combineReducers, compose, applyMiddleware } from "redux"
import thunk from "redux-thunk"
import { State, Action } from "./types"
import userReducer from "./user/reducer"

const reducer = combineReducers<State, Action>({
  user: userReducer,
})

const composeEnhancers = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
  ? (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({})
  : compose

const enhancer = composeEnhancers(applyMiddleware(thunk))

export default createStore(reducer, enhancer)
