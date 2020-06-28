import React from 'react'
import { select } from '@storybook/addon-knobs'
import RunnerStatusBadge from '~/components/utils/RunnerStatusBadge'

export default {
  title: 'Utils/RunnerStatusBadge',
  component: RunnerStatusBadge,
}

const status = ['online', 'offline', 'unknown']

export const base = () => (
  <RunnerStatusBadge status={select('Status', status, 'online')} />
)

export const all = () => (
  <>
    {status.map((c) => (
      <div style={{ marginBottom: 16 }}>
        <RunnerStatusBadge status={c} />
      </div>
    ))}
  </>
)
