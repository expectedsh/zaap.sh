import {combineReducers, createStore, compose, applyMiddleware,} from 'redux'
import thunk from 'redux-thunk'
import authentication from './authentication/reducer'
import user from './user/reducer'
import applications from './applications/reducer'
import application from './application/reducer'
import runners from './runners/reducer'

const reducer = combineReducers({
  authentication,
  user,
  applications,
  application,
  runners,
})

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export default createStore(reducer, composeEnhancers(applyMiddleware(thunk)))
