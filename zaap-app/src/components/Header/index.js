import React from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames/bind'
import style from './Header.module.scss'

const cx = classnames.bind(style)

function Header({ preTitle, title, centered = false, children }) {
  return (
    <div className={cx('root', { centered })}>
      <div className={cx('wrapper')}>
        <div className={cx('left-pane')}>
          {preTitle && (
            <div className={cx('pre-title')}>
              {preTitle}
            </div>
          )}
          <div className={cx('title')}>{title}</div>
        </div>
        {children && (
          <div className={cx("right-pane")}>
            {children}
          </div>
        )}
      </div>
    </div>
  )
}

Header.propTypes = {
  preTitle: PropTypes.string,
  title: PropTypes.string.isRequired,
  centered: PropTypes.bool,
  children: PropTypes.node,
}

export default Header
