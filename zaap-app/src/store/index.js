import {combineReducers, createStore, compose, applyMiddleware,} from 'redux'
import thunk from 'redux-thunk'
import instagramSearch from './instagramSearch/reducer'
import instagramUser from './instagramUser/reducer'
import instagramPosts from './instagramPosts/reducer'
import instagramPost from './instagramPost/reducer'

const reducer = combineReducers({
  instagramSearch,
  instagramUser,
  instagramPosts,
  instagramPost,
})

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export default createStore(reducer, composeEnhancers(applyMiddleware(thunk)))
