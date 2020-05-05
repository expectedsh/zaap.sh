import React, { useEffect, useReducer, useRef } from "react"
import { useDispatch, useSelector } from "react-redux"
import moment from "moment"
import classnames from "classnames/bind"
import { fetchApplicationLogs } from "~/store/application/actions"
import style from "./ApplicationLogs.module.scss"

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

function ApplicationLogs() {
  const dispatch = useDispatch()
  const ref = useRef(null)
  const application = useSelector(state => state.application.application)
  const [state, messageDispatch] = useReducer(messageReducer, [])

  useEffect(() => {
    let _eventSource
    dispatch(fetchApplicationLogs({ id: application.id }))
      .then((eventSource) => {
        _eventSource = eventSource
        eventSource.addEventListener("message", (event) => {
          messageDispatch(addMessage(JSON.parse(event.data)))
          ref.current.scrollBy(0, ref.current.scrollHeight)
        })
      })
      .catch(console.error)
    return () => {
      _eventSource?.close()
    }
  }, [application, ref])

  return (
    <div className={cx('root')} ref={ref}>
      {state.map((v, index) => (
        <div key={index} className={cx('log-line')}>
          <div className={cx('info')}>
            {moment(v.time).format("MMM DD HH:mm:ss.SSS")}
          </div>
          <div className={cx('info')}>
            {v.pod}
          </div>
          <div>
            {v.message}
          </div>
        </div>
      ))}
    </div>
  )
}

export default ApplicationLogs
