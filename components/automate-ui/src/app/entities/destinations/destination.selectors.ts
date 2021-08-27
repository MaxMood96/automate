import { createSelector, createFeatureSelector } from '@ngrx/store';

import { 
  DestinationEntityState, 
  destinationEntityAdapter,
  GlobalConfigEntityState,
  globalConfigEntityAdapter
 } from './destination.reducer';
import { routeParams } from 'app/route.selectors';
import { find } from 'lodash/fp';

export const destinationState = createFeatureSelector<DestinationEntityState>('destinations');
export const globalDataFeedConfigState = createFeatureSelector<GlobalConfigEntityState>('globalDataFeedConfig');

export const {
  selectAll: allDestinations,
  selectEntities: destinationEntities
} = destinationEntityAdapter.getSelectors(destinationState);

export const {
  selectAll: allglobalDataFeedConfig,
  selectEntities: globalDataFeedConfigEntities
} = globalConfigEntityAdapter.getSelectors(globalDataFeedConfigState);

export const destinationStatus = createSelector(
  destinationState,
  (state) => state.status
);

export const saveStatus = createSelector(
  destinationState,
  (state) => state.saveStatus
);

export const saveError = createSelector(
  destinationState,
  (state) => state.saveError
);

export const getStatus = createSelector(
  destinationState,
  (state) => state.getStatus
);

export const destinationFromRoute = createSelector(
  destinationEntities,
  routeParams,
  (state, { id }) => find({ id }, state)
);

export const updateStatus = createSelector(
  destinationState,
  (state) => state.updateStatus
);

export const deleteStatus = createSelector(
  destinationState,
  (state) => state.deleteStatus
);

export const destinationEnableStatus = createSelector(
  destinationState,
  (state) => state.enableStatus
);

export const globalDataFeedStatus = createSelector(
  globalDataFeedConfigState,
  (state) => state.globalConfigStatus
);
export const globalDataFeed = createSelector(
  globalDataFeedConfigState,
  (state) => state.globalConfig
);
