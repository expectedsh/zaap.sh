import React from 'react'
import ReactDOM from 'react-dom'
import { ToastContainer, toast } from "react-toastify"
import { Router } from 'react-router'
import { createBrowserHistory } from 'history'
import { Provider } from 'react-redux'
import App from './App'
import store from './store'

const history = createBrowserHistory()

ReactDOM.render((
  <Provider store={store}>
    <Router history={history}>
      <App />
    </Router>
    <ToastContainer autoClose={4000} position={toast.POSITION.TOP_RIGHT} />
  </Provider>
), document.getElementById('root'))
