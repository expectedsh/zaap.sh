import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'
import { css } from '@emotion/core'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'
import SimpleTable from '~/components/molecules/SimpleTable'
import Container from '~/components/atoms/Container'
import EmptyState from '~/components/molecules/EmptyState'
import ApplicationStatusBadge from '~/components/utils/ApplicationStatusBadge'
import Link from '~/components/atoms/Link'
import SimpleStateHandler from '~/components/molecules/SimpleStateHandler'

function AppList({
  loading, error, applications, runners,
}) {
  const tableConfig = useMemo(() => [
    {
      renderHeader: () => 'Name',
      renderCell: (app) => app.name,
      css: css`flex: 1 1 0;`,
    },
    {
      renderHeader: () => 'Status',
      renderCell: (app) => <ApplicationStatusBadge status={app.status} />,
      css: css`width: 160px;`,
    },
    {
      renderHeader: () => 'Runner',
      renderCell: (app) => {
        const runner = runners?.find((r) => r.id === app.runnerId)
        return runner
          ? <Link to="/runners">{runner.name}</Link>
          : 'Not found'
      },
      css: css`width: 260px;`,
    },
    {
      renderHeader: () => 'Last update',
      renderCell: (app) => moment(app.updatedAt).fromNow(),
      css: css`width: 160px;`,
    },
  ], [runners])

  return (
    <>
      <Header preTitle="Overview" title="Applications">
        <Button outline as="link" to="/apps/new" noMargin>
          New application
        </Button>
      </Header>

      <Container>
        <SimpleStateHandler
          loading={loading}
          error={error}
          onSuccess={(
            <SimpleTable
              config={tableConfig}
              dataSource={applications}
              // onRowClick={app => history.push(`/apps/${app.id}`)}
              noData={(
                <EmptyState
                  title="You don't have application"
                  description="Create an application and it will show up here."
                />
              )}
            />
          )}
        />
      </Container>
    </>
  )
}

AppList.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.object,
  applications: PropTypes.array,
  runners: PropTypes.array,
}

export default AppList
