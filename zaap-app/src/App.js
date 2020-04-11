import React from "react"
import { Redirect, Route, Switch } from "react-router-dom"
import Login from "./views/Login"
import SignUp from "./views/SignUp"

import "./assets/stylesheets/index.scss"

function App() {
  return (
    <>
      <Switch>
        <Route path="/login" component={Login}/>
        <Route path="/sign_up" component={SignUp}/>
        <Redirect exact path="/" to="/login"/>
      </Switch>
    </>
  )
}

export default App
