import {QuerySelection} from './ProfileSelector';
import {ProfileSelection, ProfileSelectionFromParams, SuffixParams} from '@parca/profile';
import ProfileExplorerSingle from './ProfileExplorerSingle';
import ProfileExplorerCompare from './ProfileExplorerCompare';
import {QueryServiceClient} from '@parca/client';

type NavigateFunction = (path: string, queryParams: any) => void;

interface ProfileExplorerProps {
  queryClient: QueryServiceClient;
  queryParams: any;
  navigateTo: NavigateFunction;
}

const ProfileExplorer = ({
  queryClient,
  queryParams,
  navigateTo,
}: ProfileExplorerProps): JSX.Element => {
  /* eslint-disable */
  // Disable eslint due to params being snake case
  const {
    expression_a,
    from_a,
    to_a,
    merge_a,
    labels_a,
    time_a,
    time_selection_a,
    compare_a,
    expression_b,
    from_b,
    to_b,
    merge_b,
    labels_b,
    time_b,
    time_selection_b,
    compare_b,
  } = queryParams;
  /* eslint-enable */

  const filterSuffix = (
    o: {[key: string]: string | string[] | undefined},
    suffix: string
  ): {[key: string]: string | string[] | undefined} =>
    Object.fromEntries(Object.entries(o).filter(([key]) => !key.endsWith(suffix)));

  const swapQueryParameters = (o: {
    [key: string]: string | string[] | undefined;
  }): {[key: string]: string | string[] | undefined} => {
    Object.entries(o).forEach(([key, value]) => {
      if (key.endsWith('_b')) {
        o[key.slice(0, -2) + '_a'] = value;
      }
    });
    return o;
  };

  const selectProfileA = (p: ProfileSelection) => {
    return navigateTo('/', {
      ...queryParams,
      ...SuffixParams(p.HistoryParams(), '_a'),
    });
  };

  const selectProfileB = (p: ProfileSelection) => {
    return navigateTo('/', {
      ...queryParams,
      ...SuffixParams(p.HistoryParams(), '_b'),
    });
  };

  // Show the SingleProfileExplorer when not comparing
  if (compare_a !== 'true' && compare_b !== 'true') {
    const query = {
      expression: expression_a as string,
      from: parseInt(from_a as string),
      to: parseInt(to_a as string),
      merge: (merge_a as string) === 'true',
      timeSelection: time_selection_a as string,
    };

    const profile = ProfileSelectionFromParams(
      expression_a as string,
      from_a as string,
      to_a as string,
      merge_a as string,
      labels_a as string[],
      time_a as string
    );

    const selectQuery = (q: QuerySelection) => {
      return navigateTo(
        '/',
        // Filtering the _a suffix causes us to reset potential profile
        // selection when running a new query.
        {
          ...filterSuffix(queryParams, '_a'),
          ...{
            expression_a: q.expression,
            from_a: q.from.toString(),
            to_a: q.to.toString(),
            merge_a: q.merge,
            time_selection_a: q.timeSelection,
          },
        }
      );
    };

    const selectProfile = (p: ProfileSelection) => {
      console.log({
        ...queryParams,
        ...SuffixParams(p.HistoryParams(), '_a'),
      });

      return navigateTo('/', {
        ...queryParams,
        ...SuffixParams(p.HistoryParams(), '_a'),
      });
    };

    const compareProfile = (): void => {
      let compareQuery = {
        compare_a: 'true',
        expression_a: query.expression,
        from_a: query.from.toString(),
        to_a: query.to.toString(),
        merge_a: query.merge,
        time_selection_a: query.timeSelection,

        compare_b: 'true',
        expression_b: query.expression,
        from_b: query.from.toString(),
        to_b: query.to.toString(),
        merge_b: query.merge,
        time_selection_b: query.timeSelection,
      };

      if (profile != null) {
        compareQuery = {
          ...SuffixParams(profile.HistoryParams(), '_a'),
          ...compareQuery,
        };
      }

      void navigateTo('/', compareQuery);
    };

    return (
      <ProfileExplorerSingle
        queryClient={queryClient}
        query={query}
        profile={profile}
        selectQuery={selectQuery}
        selectProfile={selectProfile}
        compareProfile={compareProfile}
      />
    );
  }

  const queryA = {
    expression: expression_a as string,
    from: parseInt(from_a as string),
    to: parseInt(to_a as string),
    merge: (merge_a as string) === 'true',
    timeSelection: time_selection_a as string,
  };
  const queryB = {
    expression: expression_b as string,
    from: parseInt(from_b as string),
    to: parseInt(to_b as string),
    merge: (merge_b as string) === 'true',
    timeSelection: time_selection_b as string,
  };

  const profileA = ProfileSelectionFromParams(
    expression_a as string,
    from_a as string,
    to_a as string,
    merge_a as string,
    labels_a as string[],
    time_a as string
  );
  const profileB = ProfileSelectionFromParams(
    expression_b as string,
    from_b as string,
    to_b as string,
    merge_b as string,
    labels_b as string[],
    time_b as string
  );

  const selectQueryA = (q: QuerySelection) => {
    return navigateTo(
      '/',
      // Filtering the _a suffix causes us to reset potential profile
      // selection when running a new query.
      {
        ...filterSuffix(queryParams, '_a'),
        ...{
          compare_a: 'true',
          expression_a: q.expression,
          from_a: q.from.toString(),
          to_a: q.to.toString(),
          merge_a: q.merge,
          time_selection_a: q.timeSelection,
        },
      }
    );
  };

  const selectQueryB = (q: QuerySelection) => {
    return navigateTo(
      '/',
      // Filtering the _b suffix causes us to reset potential profile
      // selection when running a new query.
      {
        ...filterSuffix(queryParams, '_b'),
        ...{
          compare_b: 'true',
          expression_b: q.expression,
          from_b: q.from.toString(),
          to_b: q.to.toString(),
          merge_b: q.merge,
          time_selection_b: q.timeSelection,
        },
      }
    );
  };

  const closeProfile = (card: string) => {
    let newQueryParameters = queryParams;
    if (card === 'A') {
      newQueryParameters = swapQueryParameters(queryParams);
    }

    return navigateTo('/', {
      ...filterSuffix(newQueryParameters, '_b'),
      ...{
        compare_a: 'false',
      },
    });
  };

  return (
    <ProfileExplorerCompare
      queryClient={queryClient}
      queryA={queryA}
      queryB={queryB}
      profileA={profileA}
      profileB={profileB}
      selectQueryA={selectQueryA}
      selectQueryB={selectQueryB}
      selectProfileA={selectProfileA}
      selectProfileB={selectProfileB}
      closeProfile={closeProfile}
    />
  );
};

export default ProfileExplorer;
