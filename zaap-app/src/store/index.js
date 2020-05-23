import {
  combineReducers, createStore, compose, applyMiddleware,
} from 'redux'
import thunk from 'redux-thunk'
import authentication from './authentication'

const reducer = combineReducers({
  authentication,
})

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose

export default createStore(reducer, composeEnhancers(applyMiddleware(thunk)))
