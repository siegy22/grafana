import { act, render, screen } from '@testing-library/react';
import React, { useEffect, useRef, useState } from 'react';

import { LogRowModel, dateTimeForTimeZone } from '@grafana/data';
import { convertRawToRange } from '@grafana/data/src/datetime/rangeutil';
import { config } from '@grafana/runtime';
import { LogsSortOrder } from '@grafana/schema';

import { InfiniteScroll, Props, SCROLLING_THRESHOLD } from './InfiniteScroll';
import { createLogRow } from './__mocks__/logRow';

jest.mock('lodash', () => ({
  ...jest.requireActual('lodash'),
  debounce: (callback: () => void) => {
    function debounced() {
      callback();
    }
    debounced.cancel = jest.fn();
    return debounced;
  }
}));

const defaultTz = 'browser';

const absoluteRange = {
  from: 1702578600000,
  to: 1702578900000,
};
const defaultRange = convertRawToRange({
  from: dateTimeForTimeZone(defaultTz, absoluteRange.from),
  to: dateTimeForTimeZone(defaultTz, absoluteRange.to),
});

const defaultProps: Omit<Props, 'children'> = {
  loading: false,
  loadMoreLogs: jest.fn(),
  range: defaultRange,
  rows: [],
  sortOrder: LogsSortOrder.Descending,
  timeZone: 'browser',
};

function ScrollWithWrapper({ children, ...props }: Props) {
  const [initialized, setInitialized] = useState(false);
  const scrollRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    // Required to get the ref
    if (scrollRef.current && !initialized) {
      setInitialized(true);
    }
  }, [initialized]);

  return (
    <div style={{ height: 40, overflowY: 'scroll' }} ref={scrollRef} data-testid="scroll-element">
      {initialized && (
        <InfiniteScroll {...props} scrollElement={scrollRef.current!}>
          {children}
        </InfiniteScroll>
      )}
    </div>
  );
}

beforeAll(() => {
  config.featureToggles.logsInfiniteScrolling = true;
});
afterAll(() => {
  config.featureToggles.logsInfiniteScrolling = false;
});

function setup(loadMoreMock: () => void, rows: LogRowModel[], order: LogsSortOrder, startPosition: number) {
  const { element, events } = getMockElement(startPosition);

  function scrollTo(position: number) {
    element.scrollTop = position;
    act(() => {
      events['scroll'](new Event('scroll'));
    });
  }

  function simulateScrollTo(position: number) {
    const deltaY = element.scrollTop > position ? -5 : 5;

    // Based on real user-scrolling events in the browser
    // Scroll to the position
    scrollTo(position);
    // Stay on edge
    act(() => {
      const event = new WheelEvent('wheel', { deltaY });
      events['wheel'](event);
    });
    // Trigger the actual scroll
    act(() => {
      const event = new WheelEvent('wheel', { deltaY });
      events['wheel'](event);
    });
  }

  render(
    <InfiniteScroll
      {...defaultProps}
      sortOrder={order}
      rows={rows}
      scrollElement={element as unknown as HTMLDivElement}
      loadMoreLogs={loadMoreMock}
    >
      <div data-testid="contents" style={{ height: 100 }} />
    </InfiniteScroll>
  );
  return { element, events, simulateScrollTo, scrollTo };
}

describe('InfiniteScroll', () => {
  test('Wraps components without adding DOM elements', async () => {
    const { container } = render(
      <ScrollWithWrapper {...defaultProps}>
        <div data-testid="contents" />
      </ScrollWithWrapper>
    );

    expect(await screen.findByTestId('contents')).toBeInTheDocument();
    expect(container).toMatchInlineSnapshot(`
      <div>
        <div
          data-testid="scroll-element"
          style="height: 40px; overflow-y: scroll;"
        >
          <div
            data-testid="contents"
          />
        </div>
      </div>
`);
  });

  describe.each([LogsSortOrder.Descending, LogsSortOrder.Ascending])(
    'When the sort order is %s',
    (order: LogsSortOrder) => {
      let rows: LogRowModel[];
      beforeEach(() => {
        rows = createLogRows(absoluteRange.from + 2 * SCROLLING_THRESHOLD, absoluteRange.to - 2 * SCROLLING_THRESHOLD);
      });

      test.each([
        ['up', 10, 0],
        ['down', 50, 60],
      ])(
        'Requests more logs when moving the mousewheel %s',
        async (_: string, startPosition: number, endPosition: number) => {
          const loadMoreMock = jest.fn();
          const { simulateScrollTo } = setup(loadMoreMock, rows, order, startPosition);

          expect(await screen.findByTestId('contents')).toBeInTheDocument();

          simulateScrollTo(endPosition);

          expect(loadMoreMock).toHaveBeenCalled();
          expect(await screen.findByTestId('Spinner')).toBeInTheDocument();
        }
      );

      test('Does not request more logs when there is no scroll', async () => {
        const loadMoreMock = jest.fn();
        const { element, simulateScrollTo } = setup(loadMoreMock, rows, order, 0);

        expect(await screen.findByTestId('contents')).toBeInTheDocument();
        element.clientHeight = 40;
        element.scrollHeight = element.clientHeight;

        simulateScrollTo(0);
        simulateScrollTo(40);

        expect(loadMoreMock).not.toHaveBeenCalled();
        expect(screen.queryByTestId('Spinner')).not.toBeInTheDocument();
      });

      test.each([
        ['up', 10, 0],
        ['down', 50, 60],
      ])(
        'Does not request more logs when the user is not "on edge"',
        async (_: string, startPosition: number, endPosition: number) => {
          const loadMoreMock = jest.fn();
          const { scrollTo } = setup(loadMoreMock, rows, order, startPosition);

          expect(await screen.findByTestId('contents')).toBeInTheDocument();

          scrollTo(endPosition);

          expect(loadMoreMock).not.toHaveBeenCalled();
          expect(screen.queryByTestId('Spinner')).not.toBeInTheDocument();
        }
      );

      test('Requests newer logs from the most recent timestamp', async () => {
        const startPosition = order === LogsSortOrder.Descending ? 10 : 50; // Scroll top
        const endPosition = order === LogsSortOrder.Descending ? 0 : 60; // Scroll bottom

        const loadMoreMock = jest.fn();
        const { simulateScrollTo } = setup(loadMoreMock, rows, order, startPosition);

        expect(await screen.findByTestId('contents')).toBeInTheDocument();

        simulateScrollTo(endPosition);

        expect(loadMoreMock).toHaveBeenCalledWith({
          from: rows[rows.length - 1].timeEpochMs,
          to: absoluteRange.to,
        });
      });

      test('Requests older logs from the oldest timestamp', async () => {
        const startPosition = order === LogsSortOrder.Ascending ? 10 : 50; // Scroll top
        const endPosition = order === LogsSortOrder.Ascending ? 0 : 60; // Scroll bottom

        const loadMoreMock = jest.fn();
        const { simulateScrollTo } = setup(loadMoreMock, rows, order, startPosition);

        expect(await screen.findByTestId('contents')).toBeInTheDocument();

        simulateScrollTo(endPosition);

        expect(loadMoreMock).toHaveBeenCalledWith({
          from: absoluteRange.from,
          to: rows[0].timeEpochMs,
        });
      });

      describe('With absolute range matching visible range', () => {
        test.each([
          ['top', 10, 0],
          ['bottom', 50, 60],
        ])(
          'It does not request more when scrolling %s',
          async (_: string, startPosition: number, endPosition: number) => {
            // Visible range matches the current range
            const rows = createLogRows(absoluteRange.from, absoluteRange.to);
            const loadMoreMock = jest.fn();
            const { simulateScrollTo } = setup(loadMoreMock, rows, order, startPosition);

            expect(await screen.findByTestId('contents')).toBeInTheDocument();
            
            simulateScrollTo(endPosition);

            expect(loadMoreMock).not.toHaveBeenCalled();
            expect(screen.queryByTestId('Spinner')).not.toBeInTheDocument();
            expect(await screen.findByTestId('end-of-range')).toBeInTheDocument();
          }
        );
      });

      describe('With relative range matching visible range', () => {
        test.each([
          ['top', 10, 0],
          ['bottom', 50, 60],
        ])(
          'It does not request more when scrolling %s',
          async (_: string, startPosition: number, endPosition: number) => {
            // Visible range matches the current range
            const rows = createLogRows(absoluteRange.from, absoluteRange.to);
            const loadMoreMock = jest.fn();
            const { simulateScrollTo } = setup(loadMoreMock, rows, order, startPosition);

            expect(await screen.findByTestId('contents')).toBeInTheDocument();
            
            simulateScrollTo(endPosition);

            expect(loadMoreMock).not.toHaveBeenCalled();
            expect(screen.queryByTestId('Spinner')).not.toBeInTheDocument();
            expect(await screen.findByTestId('end-of-range')).toBeInTheDocument();
          }
        );
      });
    }
  );
});

function createLogRows(from: number, to: number) {
  const rows = [createLogRow({ entry: 'line1' }), createLogRow({ entry: 'line2' })];
  // Time field
  rows[0].dataFrame.fields[0].values = [from, to];
  rows[0].timeEpochMs = from;
  rows[1].dataFrame.fields[0].values = [from, to];
  rows[1].timeEpochMs = to;
  return rows;
}

// JSDOM doesn't support layout, so we will mock the expected attribute values for the test cases.
function getMockElement(scrollTop: number) {
  const events: Record<string, (e: Event | WheelEvent) => void> = {};
  const element = {
    addEventListener: (event: string, callback: (e: Event | WheelEvent) => void) => {
      events[event] = callback;
    },
    removeEventListener: jest.fn(),
    stopImmediatePropagation: jest.fn(),
    scrollHeight: 100,
    clientHeight: 40,
    scrollTop,
    scrollTo: jest.fn(),
  };

  return { element, events };
}
