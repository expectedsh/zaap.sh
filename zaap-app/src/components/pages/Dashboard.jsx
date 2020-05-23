import React from 'react'
import PropTypes from 'prop-types'
import styled from '@emotion/styled'
import Navigation from '~/components/atoms/Navigation'
import Container from '~/components/atoms/Container'
import NavItems from '~/components/atoms/NavItems'

import logo from '~/assets/images/logo.svg'
import NavLink from '~/components/atoms/NavLink'

const StyledRoot = styled.div`
  padding-top: 64px;
  margin-bottom: 64px;
`

const StyledNavigation = styled(Navigation)`
  top: 0;
  position: fixed;
  width: 100%;
  z-index: 10000;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2), 0 1px 1px rgba(0, 0, 0, 0.14), 0 2px 1px -1px rgba(0, 0, 0, 0.12);
`

const StyledNavBrand = styled.img`
  height: 32px;
  padding: 16px 0;
  margin-right: 24px;
`

function Dashboard({ children }) {
  return (
    <StyledRoot>
      <StyledNavigation>
        <Container>
          <StyledNavBrand src={logo} alt="Zaap logo" />
          <NavItems>
            <NavLink to="/apps">Applications</NavLink>
            <NavLink to="/runners">Runners</NavLink>
          </NavItems>
          <NavItems mlAuto>
            <NavLink to="/settings">Settings</NavLink>
          </NavItems>
        </Container>
      </StyledNavigation>
      {children}
    </StyledRoot>
  )
}

Dashboard.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Dashboard
