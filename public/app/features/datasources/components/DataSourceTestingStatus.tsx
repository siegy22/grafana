import { css, cx } from '@emotion/css';
import React, { HTMLAttributes, ReactNode } from 'react';

import { DataSourceApi, DataSourceSettings as DataSourceSettingsType, GrafanaTheme2 } from '@grafana/data';
import { selectors as e2eSelectors } from '@grafana/e2e-selectors';
import { TestingStatus, config, HealthCheckResultDetails } from '@grafana/runtime';
import { AlertVariant, Alert, useTheme2, Link } from '@grafana/ui';

import { contextSrv } from '../../../core/core';
import { AccessControlAction } from '../../../types';
import { useDataSourceInstance } from '../state/hooks';
import { trackCreateDashboardClicked } from '../tracking';

export type Props = {
  testingStatus?: TestingStatus;
  exploreUrl: string;
  dataSource: DataSourceSettingsType;
};

interface AlertMessageProps extends HTMLAttributes<HTMLDivElement> {
  title: string;
  severity?: AlertVariant;
  exploreUrl: string;
  dataSourceId: string;
  onDashboardLinkClicked: () => void;
}

const getStyles = (theme: GrafanaTheme2, hasTitle: boolean) => {
  return {
    content: css`
      color: ${theme.colors.text.secondary};
      padding-top: ${hasTitle ? theme.spacing(1) : 0};
      max-height: 50vh;
      overflow-y: auto;
    `,
    disabled: css`
      pointer-events: none;
      color: ${theme.colors.text.secondary};
    `,
  };
};

const AlertSuccessMessage = ({ title, exploreUrl, dataSourceId, onDashboardLinkClicked }: AlertMessageProps) => {
  const theme = useTheme2();
  const hasTitle = Boolean(title);
  const styles = getStyles(theme, hasTitle);
  const canExploreDataSources = contextSrv.hasPermission(AccessControlAction.DataSourcesExplore);

  return (
    <div className={styles.content}>
      Next, you can start to visualize data by{' '}
      <Link
        aria-label={`Create a dashboard`}
        href={`/dashboard/new-with-ds/${dataSourceId}`}
        className="external-link"
        onClick={onDashboardLinkClicked}
      >
        building a dashboard
      </Link>
      , or by querying data in the{' '}
      <Link
        aria-label={`Explore data`}
        className={cx('external-link', {
          [`${styles.disabled}`]: !canExploreDataSources,
          'test-disabled': !canExploreDataSources,
        })}
        href={exploreUrl}
      >
        Explore view
      </Link>
      .
    </div>
  );
};

AlertSuccessMessage.displayName = 'AlertSuccessMessage';

export function DataSourceTestingStatus({ testingStatus, exploreUrl, dataSource }: Props) {
  const severity = testingStatus?.status ? (testingStatus?.status as AlertVariant) : 'error';
  const message = testingStatus?.message;

  const healthCheckDetails = testingStatus?.details;
  const detailsMessage = healthCheckDetails?.message;
  const detailsVerboseMessage = healthCheckDetails?.verboseMessage;

  const onDashboardLinkClicked = () => {
    trackCreateDashboardClicked({
      grafana_version: config.buildInfo.version,
      datasource_uid: dataSource.uid,
      plugin_name: dataSource.typeName,
      path: location.pathname,
    });
  };

  const instanceState = useDataSourceInstance(dataSource.uid);

  if (message) {
    return (
      <div className="gf-form-group p-t-2">
        <Alert severity={severity} title={message} aria-label={e2eSelectors.pages.DataSource.alert}>
          {healthCheckDetails && (
            <>
              {detailsMessage}
              {instanceState.value && renderDetailsForDataSource(instanceState.value, healthCheckDetails)}
              {severity === 'success' ? (
                <AlertSuccessMessage
                  title={message}
                  exploreUrl={exploreUrl}
                  dataSourceId={dataSource.uid}
                  onDashboardLinkClicked={onDashboardLinkClicked}
                />
              ) : null}
              {detailsVerboseMessage ? (
                <details style={{ whiteSpace: 'pre-wrap' }}>{String(detailsVerboseMessage)}</details>
              ) : null}
            </>
          )}
        </Alert>
      </div>
    );
  }

  return null;
}

export function renderDetailsForDataSource(
  dataSource: DataSourceApi,
  details: HealthCheckResultDetails
): ReactNode | null {
  const hasDetails = details && Object.entries(details).length > 0;

  if (hasDetails && dataSource.renderHealthCheckDetails) {
    return dataSource.renderHealthCheckDetails(details);
  }

  return null;
}
