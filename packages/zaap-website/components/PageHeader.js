import React from 'react'
import styled from '@emotion/styled'

const Wrapper = styled.div`
  margin: 42px 0 32px;
`

const Title = styled.div`
  font-size: 28px;
  font-weight: ${props => props.theme.fontWeightSemiBold};
  padding-bottom: 8px;
  border-bottom: 1px solid ${props => props.theme.colorGrey};
`

const PageHeader = ({ title }) => (
  <Wrapper>
    <Title>{title}</Title>
  </Wrapper>
)

export default PageHeader
