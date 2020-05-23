import React, { useEffect } from 'react';
import classnames from 'classnames/bind';
import {
  Link, Redirect, Route, Switch,
} from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { logout } from '~/store/authentication/actions';
import { fetchUser } from '~/store/user/actions';
import Alert from '~/oldcomponents/Alert';
import Settings from '~/views/Dashboard/Settings';
import ApplicationList from '~/views/Dashboard/ApplicationList';
import ApplicationNew from '~/views/Dashboard/ApplicationNew';
import ApplicationShow from '~/views/Dashboard/ApplicationShow';
import RunnerList from '~/views/Dashboard/RunnerList';
import RunnerNew from '~/views/Dashboard/RunnerNew';
import logo from '~/assets/images/logo.svg';
import style from './Dashboard.module.scss';

const cx = classnames.bind(style);

function Dashboard() {
  const dispatch = useDispatch();
  const { pending, user, error } = useSelector((state) => state.user);

  useEffect(() => {
    dispatch(fetchUser());
  }, []);

  function renderBody() {
    if (pending) {
      return (
        <div>Loading...</div>
      );
    }
    if (error) {
      return (
        <div className="container">
          <Alert className="alert alert-error" error={error} />
        </div>
      );
    }
    return user ? (
      <Switch>
        <Route path="/settings" component={Settings} />

        <Route path="/apps/new" component={ApplicationNew} />
        <Route path="/apps/:id" component={ApplicationShow} />
        <Route path="/apps" component={ApplicationList} />

        <Route path="/runners/new" component={RunnerNew} />
        <Route path="/runners" component={RunnerList} />

        <Redirect to="/apps" />
      </Switch>
    ) : null;
  }

  return (
    <div className={cx('root')}>
      <div className={cx('navbar')}>
        <div className={cx('container')}>
          <img className={cx('navbar-brand')} src={logo} alt="Zaap logo" />
          <div className={cx('navbar-links')}>
            <Link className={cx('navbar-link')} to="/apps">Applications</Link>
            <Link className={cx('navbar-link')} to="/runners">Runners</Link>
          </div>
          <div className={cx('navbar-links', 'navbar-links-right')}>
            <Link className={cx('navbar-link')} to="/settings">
              Settings
            </Link>
            <div className={cx('navbar-link')} onClick={() => dispatch(logout())}>
              Logout
            </div>
          </div>
        </div>
      </div>
      {renderBody()}
    </div>
  );
}

export default Dashboard;
