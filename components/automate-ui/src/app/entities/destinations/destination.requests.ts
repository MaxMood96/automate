import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { of as observableOf, Observable } from 'rxjs';
import { map, mergeMap } from 'rxjs/operators';
import { environment } from 'environments/environment';
import { compact, concat } from 'lodash';
import { Destination, EnableDestination, GlobalConfig } from './destination.model';
import { CreateDestinationPayload, GlobalDataFeedConfigSuccess } from './destination.actions';

export interface DestinationsResponse {
  destinations: Destination[];
}

export interface DestinationResponse {
  destination: Destination;
}
export interface GlobalDataFeedConfigResponse {
  config: GlobalConfig;
}

const DATA_FEED_URL = environment.data_feed_url;
const SECRETS_URL = environment.secrets_url;

interface KVData {
  key?: string;
  value?: string;
}

interface SecretId {
  id: string;
}

interface Secret {
  id?: string;
  name: string;
  type: string;
  data: Array<KVData>;
}

@Injectable()
export class DestinationRequests {

  constructor(private http: HttpClient) { }

  public getDestinations(): Observable<DestinationsResponse> {
    return this.http.post<DestinationsResponse>(encodeURI(
      this.joinToDataFeedUrl(['destinations'])), {});
  }

  public getDestination(id: string): Observable<Destination> {
    return this.http.get<Destination>(this.joinToDataFeedUrl(['destination', id]))
    .pipe(map((destinationsJson: Destination) => destinationsJson));
  }

  public createDestination(destinationData: CreateDestinationPayload,
    headers: string):
    Observable<DestinationResponse> {
    return this.createSecret(destinationData, headers)
      .pipe(mergeMap((secretId: string) => {
        destinationData.secret = secretId;
        return this.http.post<DestinationResponse>(
          this.joinToDataFeedUrl(['destination']), destinationData);
      }));
  }

  public updateDestination(destination: Destination): Observable<DestinationResponse> {
    return this.http.patch<DestinationResponse>(encodeURI(
      this.joinToDataFeedUrl(['destination', destination.id.toString()])), destination);
  }

  public deleteDestination(id: string): Observable<DestinationResponse> {
    return this.http.delete<DestinationResponse>(encodeURI(
      this.joinToDataFeedUrl(['destination', id.toString()])));
  }

  private createSecret(destination: CreateDestinationPayload, headers: string): Observable<string> {
    if ( headers.length > 0 ) {
      const secret = this.newSecret('', destination.name, headers);

      return this.http.post<SecretId>(`${SECRETS_URL}`, secret)
        .pipe(map(secretId => secretId.id));
    } else {
      return observableOf('');
    }
  }

  private newSecret(id: string, name: string, headers: string): Secret {
    return {
      id: id,
      name: name,
      type: 'data_feed',
      data: Array<KVData>(
        {key: 'headers', value: headers},
        {key: 'auth_type', value: 'header_auth'})
    };
  }

  public testDestination(destination: Destination): Observable<Object> {
    if (destination.secret) {
      return this.testDestinationWithSecretId(destination);
    }
  }

  public testDestinationWithUsernamePassword(url: string,
    username: string, password: string): Observable<Object> {
    return this.http.post(encodeURI(
      this.joinToDataFeedUrl(['destinations', 'test'])),
      { url, 'username_password': {username, password} });
  }

  public testDestinationWithHeaders(url: string,
    value: string): Observable<Object> {
    return this.http.post(encodeURI(
      this.joinToDataFeedUrl(['destinations', 'test'])),
      { url, 'header': {value} });
  }

  public testDestinationWithSecretId(destination:Destination): Observable<Object> {
    const SecretIdParams = this.secretIdParams(destination)
    return this.http.post(encodeURI(
      this.joinToDataFeedUrl(['destinations', 'test'])), SecretIdParams);
  }

  public secretIdParams(destination: Destination){
    if(destination.integration_types != ""){
      return { 
        'url': destination.url,
        'secret_id_with_addon': { 
          'id': destination.secret,
          'services': destination.services,
          'integration_types': destination.integration_types,
          'meta_data': destination.meta_data,
          
        } 
      }
    } else {
      return { 
        'url': destination.url,
        'secret_id': { 
          'id': destination.secret
          
        } 
      }
    }
  }

  public testDestinationWithNoCreds(url: string): Observable<Object> {
    return this.http.post(encodeURI(
      this.joinToDataFeedUrl(['destinations', 'test'])), { url, 'none': {}});
  }

  public enableDisableDestinations(destination: EnableDestination): Observable<DestinationResponse> {
    return this.http.patch<DestinationResponse>(encodeURI(
      this.joinToDataFeedUrl(['destination','enable',destination.id.toString()])), destination);
  }
  public globalDataFeedConfig(): Observable<GlobalDataFeedConfigSuccess> {
    return this.http.get<GlobalDataFeedConfigSuccess>(encodeURI(
      this.joinToDataFeedUrl(['config'])));
  }
  // Generate an notifier url from a list of words.
  // falsey values; false, null, 0, "", undefined, and NaN are ignored
  // example: ['foo', 'bar', null, 'baz'] == '${NOTIFIERURL}/foo/bar/baz'
  private joinToDataFeedUrl(words: string[]): string {
    return compact(concat([DATA_FEED_URL], words)).join('/');
  }
}
