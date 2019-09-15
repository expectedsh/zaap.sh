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
  font-weight: ${props => props.theme.fontWeightSemiBold};
`

const CreatedAt = styled.div`
  font-size: 14px;
  font-weight: ${props => props.theme.fontWeightSemiBold};
  color: ${props => props.theme.colorTextSecondary};
  text-transform: uppercase;
`

const Dot = styled.div`
  min-height: 16px;
  min-width: 16px;
  max-height: 16px;
  max-width: 16px;
  margin-right: 12px;
  background: ${props => props.online ? props.theme.colorGreen : props.theme.colorRed};
  border-radius: 50%;
`

const ViewButton = styled.a`
  display: flex;
  padding: 6px 12px;
  color: ${props => props.theme.colorPrimary} !important;
  background: #FFF;
  border-radius: .25rem;
  & > .material-icons {
    margin-left: 4px;
  }
  &:hover {
    background: rgba(0, 105, 255, .1);
  }
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
      <ViewButton>
        View <i class="material-icons">arrow_forward</i>
      </ViewButton>
    </ItemsContainer>
  </Wrapper>
)

export default RouteCard
