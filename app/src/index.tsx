import "./index.css"
import React from "react"
import ReactDOM from "react-dom"
import { ThemeProvider } from "emotion-theming"
import { createBrowserHistory } from "history"
import { Router, Switch, Route } from "react-router-dom"
import { theme } from "utils/theme"
import Login from "views/login"

const history = createBrowserHistory()

const App = () => {

  return (
    <ThemeProvider theme={theme}>
      <Router history={history}>
        <Switch>
          <Route path="/login" component={Login}/>
          <Route path="/sign_up" render={() => <p>Sign Up</p>}/>
          <Route path="/home" render={() => <p>Home</p>}/>
        </Switch>
      </Router>
    </ThemeProvider>
  )
}

ReactDOM.render(<App />, document.querySelector('#root'))
