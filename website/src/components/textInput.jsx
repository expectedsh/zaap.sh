import React from 'react'
import styles from './textInput.scss'

const TextInput = (props) => {
  return (
    <input {...props} className={styles.root} />
  )
}

export default TextInput
