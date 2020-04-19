import React, { useEffect, useReducer } from "react"
import { useDispatch, useSelector } from "react-redux"
import classnames from "classnames/bind"
import moment from "moment"
import { fetchApplicationLogs } from "~/store/application/actions"
import style from "./Logs.module.scss"

const cx = classnames.bind(style)

function addMessage(payload) {
  return { type: 'ADD_MESSAGE', payload }
}

function messageReducer(state, action) {
  switch (action.type) {
  case 'ADD_MESSAGE':
    return [...state, action.payload]
  }
}

function Logs() {
  const dispatch = useDispatch()
  const application = useSelector(state => state.application.application)
  const [state, messageDispatch] = useReducer(messageReducer, [])

  useEffect(() => {
    let _eventSource
    dispatch(fetchApplicationLogs({ id: application.id }))
      .then((eventSource) => {
        _eventSource = eventSource
        eventSource.addEventListener("message", (event) => {
          messageDispatch(addMessage(JSON.parse(event.data)))
        })
      })
      .catch(console.error)
    return () => {
      _eventSource?.close()
    }
  }, [application])

  return (
    <div className={cx('root')}>
      {state.map(v => (
        <div className={cx('log-line')}>
          <div className={cx('time')}>
            {v.time}
          </div>
          <div className={cx('message')}>
            {v.message}
          </div>
        </div>
      ))}
    </div>
  )
}

export default Logs
