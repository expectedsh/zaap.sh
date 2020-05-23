import React from 'react';
import classnames from 'classnames/bind';
import { Redirect, Route, Switch } from 'react-router-dom';
import WithBodyClasses from '~/oldcomponents/WithBodyClasses';
import SignIn from '~/views/Auth/SignIn';
import SignUp from '~/views/Auth/SignUp';
import logo from '~/assets/images/logo.svg';
import style from '~/views/Auth/Auth.module.scss';

const cx = classnames.bind(style);

function Auth() {
  return (
    <WithBodyClasses classNames={[cx('auth-background')]}>
      <div className={cx('root')}>
        <img className={cx('logo')} src={logo} alt="Zaap logo" />
        <Switch>
          <Route path="/sign_in" component={SignIn} />
          <Route path="/sign_up" component={SignUp} />
          <Redirect to="/sign_in" />
        </Switch>
      </div>
    </WithBodyClasses>
  );
}

export default Auth;
