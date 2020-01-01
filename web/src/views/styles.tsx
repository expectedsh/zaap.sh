import styled from "utils/styled"

export const Root = styled.div`
  display: flex;
  flex-direction: row;
`

export const Test = styled.div`
  background: ${props => props.theme.color.blueDark};
  padding: 16px 24px;
  line-height: 24px;
  display: flex;
  flex-direction: row;
  align-self: center;
  justify-content: space-between;
  width: 100%;
  color: ${props => props.theme.color.white};
`

export const Sidebar = styled.div`
  display: flex;
  flex-direction: column;
  width: 250px;
  min-height: 100vh;
  background: ${props => props.theme.color.blue};
`

export const SidebarLogo = styled.img`
  display: flex;
  align-self: center;
  margin: 24px;
`

export const SidebarLinkList = styled.ul`
  list-style: none;
  margin-bottom: 24px;
  padding: 0;

  a {
    display: flex;
    flex-direction: row;
    align-self: center;
    text-decoration: none;
    margin: 0 24px 8px 16px;
    padding: 8px;
    font-size: ${props => props.theme.text.fontSizeLarge};
    line-height: 24px;
    color: ${props => props.theme.color.white};
    border-radius: 4px;

    i {
      padding-right: 8px;
    }

    &:hover {
      background: ${props => props.theme.color.blueDark};
    }
  }
`

export const Content = styled.div`
`
