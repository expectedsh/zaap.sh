import React from 'react'
import { css } from '@emotion/core'
import styled from '@emotion/styled'

const Wrapper = styled.div`
  border: 1px solid ${props => props.theme.colorGrey};
  border-radius: .25rem;
  background: #FFF;
  padding: 24px;
  margin-bottom: 32px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12);
`

const ItemsContainer = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
`

const Title = styled.div`
  font-size: 24px;
`

const CreatedAt = styled.div`
  font-size: 14px;
  font-weight: ${props => props.theme.fontWeightBold};
  color: ${props => props.theme.colorTextSecondary};
  text-transform: uppercase;
`

const Dot = styled.div`
  min-height: 16px;
  min-width: 16px;
  max-height: 16px;
  max-width: 16px;
  background: ${props => props.online ? props.theme.colorGreen : props.theme.colorRed};
  border-radius: 50%;
`

const RouteCard = ({ online }) => (
  <Wrapper>
    <ItemsContainer>
      <Title>Get Articles</Title>
      <Dot online={online} />
    </ItemsContainer>
    <CreatedAt>Created a month ago</CreatedAt>
    <ItemsContainer style={{ marginTop: 12 }}>
      <div>
        Path: GET Method: /articles
      </div>
      <div>
        View
      </div>
    </ItemsContainer>
  </Wrapper>
)

export default RouteCard
