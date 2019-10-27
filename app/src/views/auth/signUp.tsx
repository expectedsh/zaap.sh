import React, { FormEvent, useState } from "react"
import { Link } from "react-router-dom"
import { useThunkDispatch } from "hooks"
import { TextInput, Button } from "components"
import Layout from "./layout"

const Alternative = () => (
  <>
    Already have an account? <Link to="/sign_in">Sign In</Link>.
  </>
)

const SignUp = () => {
  const dispatch = useThunkDispatch()
  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<Error | undefined>(undefined)

  const onSubmit = (event: FormEvent<HTMLFormElement>) => {

  }

  return (
    <Layout title="Sign Up" alternative={<Alternative />} loading={loading} error={error}>
      <form onSubmit={onSubmit}>
        <TextInput type="text" placeholder="Name" autoComplete="given-name"
          onChange={e => setName(e.target.value)} />
        <TextInput type="email" placeholder="Email Address" autoComplete="email"
          onChange={e => setEmail(e.target.value)} />
        <Button type="submit" variant="success">
          Continue
        </Button>
      </form>
    </Layout>
  )
}

export default SignUp
