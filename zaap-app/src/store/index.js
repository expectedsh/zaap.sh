import {
  combineReducers, createStore, compose, applyMiddleware,
} from 'redux'
import thunk from 'redux-thunk'
import authentication from './authentication'
import applications from './applications'
import runners from './runners'

const reducer = combineReducers({
  authentication,
  applications,
  runners,
})

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose

export default createStore(reducer, composeEnhancers(applyMiddleware(thunk)))
