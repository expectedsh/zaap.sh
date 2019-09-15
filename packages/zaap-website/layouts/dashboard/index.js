import 'bootstrap/dist/css/bootstrap-reboot.css'
import 'bootstrap/dist/css/bootstrap-grid.css'
import React from 'react'
import Head from 'next/head'
import { Global, css } from '@emotion/core'
import GlobalLayout from '../global'
import Header from './header'
import Footer from './footer'

const DashboardLayout = ({ children }) => (
  <GlobalLayout>
    <Head>
      <title>Zaap | Dashboard</title>
      <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Material+Icons|Montserrat:400,500,600,700&display=swap" />
    </Head>
    <Header></Header>
    {children}
    <Footer></Footer>
    <Global styles={props => css`
      html {
        position: relative;
        min-height: 100%;
      }
      body {
        margin-bottom: 100px;
        background: ${props.colorGreyLight};
        font-family: ${props.fontFamily};
      }
    `} />
  </GlobalLayout>
)

export default DashboardLayout
