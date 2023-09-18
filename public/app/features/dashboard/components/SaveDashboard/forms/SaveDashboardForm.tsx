import { css } from '@emotion/css';
import React, { useMemo, useState } from 'react';

import { GrafanaTheme2 } from '@grafana/data';
import { selectors } from '@grafana/e2e-selectors';
import { Stack } from '@grafana/experimental';
import { config } from '@grafana/runtime';
import { Button, Checkbox, Form, TextArea, useStyles2 } from '@grafana/ui';
import { DashboardModel } from 'app/features/dashboard/state';

import { GenAIDashboardChangesButton } from '../../GenAI/GenAIDashboardChangesButton';
import { SaveDashboardData, SaveDashboardOptions } from '../types';

interface FormDTO {
  message: string;
}

export type SaveProps = {
  dashboard: DashboardModel; // original
  isLoading: boolean;
  saveModel: SaveDashboardData; // already cloned
  onCancel: () => void;
  onSuccess: () => void;
  onSubmit?: (clone: DashboardModel, options: SaveDashboardOptions, dashboard: DashboardModel) => Promise<any>;
  options: SaveDashboardOptions;
  onOptionsChange: (opts: SaveDashboardOptions) => void;
};

export const SaveDashboardForm = ({
  dashboard,
  isLoading,
  saveModel,
  options,
  onSubmit,
  onCancel,
  onSuccess,
  onOptionsChange,
}: SaveProps) => {
  const hasTimeChanged = useMemo(() => dashboard.hasTimeChanged(), [dashboard]);
  const hasVariableChanged = useMemo(() => dashboard.hasVariableValuesChanged(), [dashboard]);

  const [saving, setSaving] = useState(false);
  const styles = useStyles2(getStyles);

  return (
    <Form
      onSubmit={async (data: FormDTO) => {
        if (!onSubmit) {
          return;
        }
        setSaving(true);
        options = { ...options, message: data.message };
        const result = await onSubmit(saveModel.clone, options, dashboard);
        if (result.status === 'success') {
          if (options.saveVariables) {
            dashboard.resetOriginalVariables();
          }
          if (options.saveTimerange) {
            dashboard.resetOriginalTime();
          }
          onSuccess();
        } else {
          setSaving(false);
        }
      }}
    >
      {({ register, errors }) => {
        const messageProps = register('message');
        return (
          <Stack gap={2} direction="column" alignItems="flex-start">
            {hasTimeChanged && (
              <Checkbox
                checked={!!options.saveTimerange}
                onChange={() =>
                  onOptionsChange({
                    ...options,
                    saveTimerange: !options.saveTimerange,
                  })
                }
                label="Save current time range as dashboard default"
                aria-label={selectors.pages.SaveDashboardModal.saveTimerange}
              />
            )}
            {hasVariableChanged && (
              <Checkbox
                checked={!!options.saveVariables}
                onChange={() =>
                  onOptionsChange({
                    ...options,
                    saveVariables: !options.saveVariables,
                  })
                }
                label="Save current variable values as dashboard default"
                aria-label={selectors.pages.SaveDashboardModal.saveVariables}
              />
            )}
            <div className={styles.changesDescription}>
              {config.featureToggles.dashgpt && (
                <GenAIDashboardChangesButton
                  onGenerate={(text) => {
                    onOptionsChange({
                      ...options,
                      message: text,
                    });
                    messageProps.onChange({ currentTarget: { value: text } });
                  }}
                  diff={saveModel.diff}
                />
              )}
              <TextArea
                {...messageProps}
                aria-label="message"
                value={options.message}
                onChange={(e) => {
                  onOptionsChange({
                    ...options,
                    message: e.currentTarget.value,
                  });
                  messageProps.onChange(e);
                }}
                placeholder="Add a note to describe your changes."
                autoFocus
                rows={5}
              />
            </div>

            <Stack alignItems="center">
              <Button variant="secondary" onClick={onCancel} fill="outline">
                Cancel
              </Button>
              <Button
                type="submit"
                disabled={!saveModel.hasChanges || isLoading}
                icon={saving ? 'fa fa-spinner' : undefined}
                aria-label={selectors.pages.SaveDashboardModal.save}
              >
                {isLoading ? 'Saving...' : 'Save'}
              </Button>
              {!saveModel.hasChanges && <div>No changes to save</div>}
            </Stack>
          </Stack>
        );
      }}
    </Form>
  );
};

function getStyles(theme: GrafanaTheme2) {
  return {
    changesDescription: css`
      display: flex;
      align-items: end;
      flex-direction: column;
      width: 100%;
    `,
  };
}
