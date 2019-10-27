import React from "react"
import { Link } from "react-router-dom"
import * as S from "./styles"
import logo from "assets/images/logo.svg"

const Home = () => {

  return (
    <S.Root>
      <S.Sidebar>
        <S.SidebarLogo src={logo} title="Zaap" alt="Zaap" />
        <S.Test>
          my-project
          <i className="material-icons">keyboard_arrow_down</i>
        </S.Test>
        <S.SidebarLinkList>
          <Link to="/overview">
            <i className="material-icons">dashboard</i>
            Overview
          </Link>
          <Link to="/schema">
            <i className="material-icons">code</i>
            Schema
          </Link>
          <Link to="/settings">
            <i className="material-icons">settings</i>
            Settings
          </Link>
        </S.SidebarLinkList>
      </S.Sidebar>
      <S.Content>
        Hello
      </S.Content>
    </S.Root>
  )
}

export default Home
