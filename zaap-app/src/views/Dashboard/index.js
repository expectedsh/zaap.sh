import React, { useEffect } from "react"
import classnames from "classnames/bind"
import { Link, Redirect, Route, Switch } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { logout } from "~/store/authentication/actions"
import { fetchUser } from "~/store/user/actions"
import Alert from "~/components/Alert"
import Settings from "~/views/Dashboard/Settings"
import ListApps from "~/views/Dashboard/ListApps"
import NewApp from "~/views/Dashboard/NewApp"
import logo from "~/assets/images/logo.svg"
import style from "./Dashboard.module.scss"

const cx = classnames.bind(style)

function Dashboard() {
  const dispatch = useDispatch()
  const { pending, user, error } = useSelector(state => state.user)

  useEffect(() => {
    dispatch(fetchUser())
  }, [])

  function renderBody() {
    if (pending) {
      return (
        <div>Loading...</div>
      )
    }
    if (error) {
      return (
        <div className="container">
          <Alert className="alert alert-error" error={error} />
        </div>
      )
    }
    return user ? (
      <Switch>
        <Route path="/settings" component={Settings}/>
        <Route path="/apps/new" component={NewApp}/>
        <Route path="/apps" component={ListApps}/>
        <Redirect to="/apps"/>
      </Switch>
    ) : null
  }

  return (
    <>
      <div className={cx("navbar")}>
        <div className={cx("container")}>
          <img className={cx("navbar-brand")} src={logo} alt="Zaap logo"/>
          <div className={cx("navbar-links")}>
            <Link className={cx("navbar-link")} to="/apps">Applications</Link>
          </div>
          <div className={cx("navbar-links", "navbar-links-right")}>
            <Link className={cx("navbar-link")} to="/settings">
              Settings
            </Link>
            <div className={cx("navbar-link")} onClick={() => dispatch(logout())}>
              Logout
            </div>
          </div>
        </div>
      </div>
      {renderBody()}
    </>
  )
}

export default Dashboard
