import '~/assets/stylesheets/reset.scss'

import React from 'react'
import ReactDOM from 'react-dom'
import { ToastContainer, toast } from 'react-toastify'
import { Router } from 'react-router'
import { createBrowserHistory } from 'history'
import { Provider } from 'react-redux'
import ThemeProvider from '~/style/themeProvider'
import store from '~/store'
import AppCont from '~/containers/AppCont'

const history = createBrowserHistory()

ReactDOM.render((
  <Provider store={store}>
    <ThemeProvider>
      <Router history={history}>
        <AppCont />
      </Router>
    </ThemeProvider>
    <ToastContainer autoClose={2000} position={toast.POSITION.TOP_RIGHT} />
  </Provider>
), document.getElementById('root'))
