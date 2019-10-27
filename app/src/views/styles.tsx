import styled from "utils/styled"

export const Root = styled.div`
  display: flex;
  flex-direction: row;
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
  align-self: flex-start;
  margin: 24px;
`

export const SidebarLinkList = styled.ul`
  list-style: none;
  margin: 0 0 24px 24px;
  padding: 0;

  a {
    display: flex;
    flex-direction: row;
    align-self: center;
    text-decoration: none;
    margin: 0 24px 8px 0;
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
