import React, { Suspense, lazy } from 'react'
import { Redirect, Route, Switch } from 'react-router-dom'
import Auth from '~/components/pages/Auth/Auth'

const SignInRoute = lazy(() => import('~/components/pages/AuthSignIn/AuthSignInCont'))
const SignUpRoute = lazy(() => import('~/components/pages/AuthSignUp/AuthSignUpCont'))

function AuthCont() {
  return (
    <Auth>
      <Suspense fallback={<p>error</p>}>
        <Switch>
          <Route path="/sign_in" component={SignInRoute} />
          <Route path="/sign_up" component={SignUpRoute} />
          <Redirect to="/sign_in" />
        </Switch>
      </Suspense>
    </Auth>
  )
}

export default AuthCont
