import React, { useMemo } from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'
import { css } from '@emotion/core'
import Header from '~/components/molecules/Header'
import Button from '~/components/atoms/Button'
import SimpleTable from '~/components/molecules/SimpleTable'
import Container from '~/components/atoms/Container'
import EmptyState from '~/components/molecules/EmptyState'
import RunnerStatusBadge from '~/components/utils/RunnerStatusBadge'
import SimpleStateHandler from '~/components/molecules/SimpleStateHandler'

function RunnerList({ loading, error, runners }) {
  const tableConfig = useMemo(() => [
    {
      renderHeader: () => 'Name',
      renderCell: (runner) => runner.name,
      css: css`flex: 1 1 0;`,
    },
    {
      renderHeader: () => 'Status',
      renderCell: (runner) => <RunnerStatusBadge status={runner.status} />,
      css: css`width: 160px;`,
    },
    {
      renderHeader: () => 'Endpoint',
      renderCell: (runner) => runner.url,
      css: css`width: 260px;`,
    },
    {
      renderHeader: () => 'Last update',
      renderCell: (runner) => moment(runner.updatedAt).fromNow(),
      css: css`width: 160px;`,
    },
  ], [runners])

  return (
    <>
      <Header preTitle="Overview" title="Runners">
        <Button outline as="link" to="/runners/new" noMargin>
          Register runner
        </Button>
      </Header>

      <Container>
        <SimpleStateHandler
          loading={loading}
          error={error}
          onSuccess={(
            <SimpleTable
              config={tableConfig}
              dataSource={runners}
              noData={(
                <EmptyState
                  title="You don't have runner"
                  description="Register a runner and it will show up here."
                />
              )}
            />
          )}
        />
      </Container>
    </>
  )
}

RunnerList.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.object,
  runners: PropTypes.array,
}

export default RunnerList
