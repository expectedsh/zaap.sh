import React from 'react';
import PropTypes from 'prop-types';
import { NavLink as RouterLink } from 'react-router-dom';
import classnames from 'classnames/bind';
import style from './NavigationBar.module.scss';

const cx = classnames.bind(style);

function Link({ to, children }) {
  return (
    <RouterLink exact to={to} className={cx('link')} activeClassName={cx('link-active')}>
      {children}
    </RouterLink>
  );
}

Link.propTypes = {
  to: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
};

function NavigationBar({ children, style }) {
  return (
    <div className={cx('root')} style={style}>
      <div className={cx('wrapper')}>
        {children}
      </div>
    </div>
  );
}

NavigationBar.Link = Link;

NavigationBar.propTypes = {
  style: PropTypes.any,
  children: PropTypes.node,
};

export default NavigationBar;
