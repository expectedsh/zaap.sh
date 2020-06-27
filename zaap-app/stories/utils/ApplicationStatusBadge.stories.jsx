import React from 'react'
import { select } from '@storybook/addon-knobs'
import ApplicationStatusBadge from '~/components/utils/ApplicationStatusBadge'

export default {
  title: 'Utils/ApplicationStatusBadge',
  component: ApplicationStatusBadge,
}

const status = ['deploying', 'running', 'crashed', 'failed', 'unknown']

export const base = () => (
  <ApplicationStatusBadge status={select('Status', status, 'running')} />
)

export const all = () => (
  <>
    {status.map((c) => (
      <div style={{ marginBottom: 16 }}>
        <ApplicationStatusBadge status={c} />
      </div>
    ))}
  </>
)
