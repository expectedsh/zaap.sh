import React from 'react'
import classnames from "classnames/bind"
import style from "./ListApps.module.scss"

const cx = classnames.bind(style)

function Home() {
  return (
    <div className={cx('root')}>
      Home
    </div>
  )
}

export default Home
