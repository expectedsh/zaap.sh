import React from 'react'
import PropTypes from 'prop-types'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'

function AppList({ loading, error, applications }) {

  return (
    <>
      <Header preTitle="Overview" title="Applications">
        <Button outline as="a" href="/apps/new" noMargin>
          New application
        </Button>
      </Header>
    </>
  )
}

AppList.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.object,
  applications: PropTypes.array,
}

export default AppList
