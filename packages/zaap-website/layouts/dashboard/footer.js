import React from 'react'
import styled from '@emotion/styled'

const Wrapper = styled.div`
  position: absolute;
  bottom: 0;
  left: 0;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100px;
  border-top: 1px solid ${props => props.theme.colorGrey};
`

const Copyright = styled.div`
  color: ${props => props.theme.colorTextSecondary};
  margin-right: 16px;
`

const Links = styled.div`
  a {
    color: ${props => props.theme.colorTextSecondary};
    text-decoration: underline;
    margin-right: 12px;
    &:last-child {
      margin-right: 0;
    }
  }
`

const Footer = () => (
  <Wrapper>
    <Copyright>
      © 2019 Zaap, Inc.
    </Copyright>
    <Links>
      <a href="#">Support</a>
      <a href="#">Terms</a>
      <a href="#">Privacy</a>
    </Links>
  </Wrapper>
)

export default Footer
