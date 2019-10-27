import "./index.css"
import React from "react"
import ReactDOM from "react-dom"
import { Provider } from "react-redux"
import { ThemeProvider } from "emotion-theming"
import { createBrowserHistory } from "history"
import { Router, Switch, Route } from "react-router-dom"
import { theme } from "utils/theme"
import store from "store"
import Home from "views/home"
import SignIn from "views/auth/signIn"
import SignUp from "views/auth/signUp"

const history = createBrowserHistory()

const App = () => {
  return (
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <Router history={history}>
          <Switch>
            <Route path="/" component={Home}/>
            <Route path="/sign_in" component={SignIn}/>
            <Route path="/sign_up" component={SignUp}/>
          </Switch>
        </Router>
      </ThemeProvider>
    </Provider>
  )
}

ReactDOM.render(<App />, document.querySelector('#root'))
