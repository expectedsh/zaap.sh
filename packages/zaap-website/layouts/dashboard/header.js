import React from 'react'
import styled from '@emotion/styled'

const Wrapper = styled.div`
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.2);
  & > .container {
    height: 64px;
    display: flex;
    flex-direction: row;
    align-items: center;
  }
`

const NavBrand = styled.div`
  font-size: 32px;
  font-weight: ${props => props.theme.fontWeightSemiBold};
  color: ${props => props.theme.colorTextPrimary};
  margin-right: 26px;
`

const NavContainer = styled.div`
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`

const Nav = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
`

const NavLink = styled.div`
  margin-right: 12px;
  &:last-child {
    margin-right: 0;
  }
  & > a {
    color: ${props => props.theme.colorTextPrimary};
    &:hover {
      text-decoration: none;
    }
  }
`

const Header = () => (
  <Wrapper>
    <div className="container">
      <NavBrand>ZAAP</NavBrand>
      <NavContainer>
        <Nav>
          <NavLink active>
            <a href="#">Overview</a>
          </NavLink>
          <NavLink>
            <a href="#">Settings</a>
          </NavLink>
        </Nav>
        <Nav>
          <NavLink>
            <a href="#">Documentation</a>
          </NavLink>
        </Nav>
      </NavContainer>
    </div>
  </Wrapper>
)

export default Header
