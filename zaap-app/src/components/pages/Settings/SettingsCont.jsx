import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { toast } from 'react-toastify'
import { userService } from '~/services'
import { setUser } from '~/store/user'
import Settings from './Settings'

function SettingsCont() {
  const dispatch = useDispatch()
  const user = useSelector((s) => s.user.user)
  const [isLoading, setLoading] = useState(true)
  const [error, setError] = useState(undefined)

  useEffect(() => {
    userService.findMe()
      .then((fetchedUser) => dispatch(setUser(fetchedUser)))
      .catch((err) => setError(err))
      .finally(() => setLoading(false))
  }, [])

  function updateProfile(values) {
    return userService.updateMe(values)
      .then(() => {
        toast.success('Profile updated.')
      })
      .catch((err) => {
        if (err.response.status === 422) {
          return error.data
        }
        toast.error(err.response.statusText)
        return undefined
      })
  }

  return (
    <Settings
      loading={isLoading}
      error={error}
      user={user}
      updateProfile={updateProfile}
    />
  )
}

export default SettingsCont
