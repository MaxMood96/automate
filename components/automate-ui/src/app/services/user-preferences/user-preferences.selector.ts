import { createSelector, createFeatureSelector } from '@ngrx/store';

import { UserPreferencesEntityState } from './user-preferences.reducer';

export const userPreferencesState = createFeatureSelector<UserPreferencesEntityState>('userPreferences');

export const userPreferencesList = createSelector(
  userPreferencesState,
  (userPreferences) => userPreferences.list
);

export const userPreferencesStatus = createSelector(
  userPreferencesState,
  (userPreferences) => userPreferences.status
);

export const userPreferencesError = createSelector(
  userPreferencesState,
  (userPreferences) => userPreferences.error
);

/////////////////////////// Testing with timezone available
export const userPreferencesTimezone = createSelector(
  userPreferencesState,
  (userPreferences) => userPreferences.list.timezone
);
