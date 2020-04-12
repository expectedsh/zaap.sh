import "./assets/stylesheets/index.scss"

import React from "react"
import { useSelector } from "react-redux"
import Auth from "~/views/Auth"
import Dashboard from "~/views/Dashboard"

function App() {
  const token = useSelector(state => state.authentication.token)

  return token ? <Dashboard/> : <Auth/>
}

export default App
