import React from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'
import { css } from '@emotion/core'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'
import Callout from '~/components/molecules/Callout'
import SimpleTable from '~/components/molecules/SimpleTable'
import Container from '~/components/atoms/Container'
import EmptyState from '~/components/molecules/EmptyState'

const tableConfig = (apps) => [
  {
    renderHeader: () => 'Name',
    renderCell: (app) => app.name,
    css: css`flex: 1 1 0;`,
  },
  {
    renderHeader: () => 'Status',
    renderCell: (app) => app.status,
    css: css`width: 160px;`,
  },
  {
    renderHeader: () => 'Runner',
    renderCell: (app) => 'Not found',
    css: css`width: 260px;`,
  },
  {
    renderHeader: () => 'Last update',
    renderCell: (app) => moment(app.updatedAt).fromNow(),
    css: css`width: 160px;`,
  },
]

function AppList({ loading, error, applications }) {
  function renderBody() {
    if (loading) {
      return 'Loading...'
    }
    if (error) {
      return (
        <Callout color="danger" block>
          {error.message}
        </Callout>
      )
    }
    return (
      <SimpleTable
        config={tableConfig(applications)}
        dataSource={applications}
        // onRowClick={app => history.push(`/apps/${app.id}`)}
        noData={(
          <EmptyState
            title="You don't have application"
            description="Create an application and it will show up here."
          />
        )}
      />
    )
  }

  return (
    <>
      <Header preTitle="Overview" title="Applications">
        <Button outline as="link" to="/apps/new" noMargin>
          New application
        </Button>
      </Header>

      <Container>
        {renderBody()}
      </Container>
    </>
  )
}

AppList.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.object,
  applications: PropTypes.array,
}

export default AppList
