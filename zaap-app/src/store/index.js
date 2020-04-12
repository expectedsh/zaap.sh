import {combineReducers, createStore, compose, applyMiddleware,} from 'redux'
import thunk from 'redux-thunk'
import authentication from './authentication/reducer'
import user from './user/reducer'

const reducer = combineReducers({
  authentication,
  user,
})

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export default createStore(reducer, composeEnhancers(applyMiddleware(thunk)))
