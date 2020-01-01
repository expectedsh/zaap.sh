import "./index.css"
import React from "react"
import ReactDOM from "react-dom"
import { Provider } from "react-redux"
import { ThemeProvider } from "emotion-theming"
import { createBrowserHistory } from "history"
import { Router, Switch, Route, Redirect } from "react-router-dom"
import { theme } from "utils/theme"
import { useSelector } from "hooks"
import store from "store"
import Home from "views/home"
import SignIn from "views/auth/signIn"
import SignUp from "views/auth/signUp"

const history = createBrowserHistory()

const App = () => {
  const isLogged = useSelector(state => !!state.user.token)

  return (
    <Router history={history}>
      <Switch>
        {isLogged ? (
          <>
            <Route path="/" component={Home} />
            <Redirect to="/" />
          </>
        ) : (
          <>
            <Route path="/sign_in" component={SignIn} />
            <Route path="/sign_up" component={SignUp} />
            <Redirect to="/sign_in" />
          </>
        )}
      </Switch>
    </Router>
  )
}

ReactDOM.render((
  <Provider store={store}>
    <ThemeProvider theme={theme}>
      <App />
    </ThemeProvider>
  </Provider>
), document.querySelector('#root'))
