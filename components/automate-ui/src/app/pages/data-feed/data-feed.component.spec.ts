import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { waitForAsync, ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpErrorResponse } from '@angular/common/http';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MockComponent } from 'ng2-mock-component';
import { StoreModule, Store } from '@ngrx/store';
import { NgrxStateAtom, ngrxReducers, runtimeChecks } from 'app/ngrx.reducers';
import { DataFeedComponent } from './data-feed.component';
import { DestinationRequests } from 'app/entities/destinations/destination.requests';
import { HttpClient, HttpHandler } from '@angular/common/http';
import { Destination } from 'app/entities/destinations/destination.model';
import { CreateDestinationSuccess, CreateDestinationFailure, CreateDestination } from 'app/entities/destinations/destination.actions';
import { HttpStatus } from 'app/types/types';
import { FeatureFlagsService } from 'app/services/feature-flags/feature-flags.service';
import {
  AuthTypes,
  DataFeedCreateComponent,
  IntegrationTypes,
  StorageIntegrationTypes,
  WebhookIntegrationTypes
} from '../data-feed-create/data-feed-create.component';

describe('DataFeedComponent', () => {
  let component: DataFeedComponent;
  let fixture: ComponentFixture<DataFeedComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [
        DataFeedCreateComponent,
        DataFeedComponent,
        MockComponent({
        selector: 'app-create-data-feed-modal',
        inputs: ['visible', 'creating', 'conflictErrorEvent', 'createForm'],
        outputs: ['close', 'createClicked']
        }),
        MockComponent({ selector: 'app-delete-object-modal',
        inputs: ['default', 'visible', 'objectNoun', 'objectName'],
        outputs: ['close', 'deleteClicked'] }),
        MockComponent({ selector: 'chef-button',
                inputs: ['disabled', 'routerLink'] }),
        MockComponent({ selector: 'chef-error' }),
        MockComponent({ selector: 'chef-form-field' }),
        MockComponent({ selector: 'chef-heading' }),
        MockComponent({ selector: 'chef-icon' }),
        MockComponent({ selector: 'chef-loading-spinner' }),
        MockComponent({ selector: 'mat-select' }),
        MockComponent({ selector: 'mat-option' }),
        MockComponent({ selector: 'chef-page-header' }),
        MockComponent({ selector: 'chef-subheading' }),
        MockComponent({ selector: 'chef-toolbar' }),
        MockComponent({ selector: 'chef-table' }),
        MockComponent({ selector: 'chef-thead' }),
        MockComponent({ selector: 'chef-tbody' }),
        MockComponent({ selector: 'chef-tr' }),
        MockComponent({ selector: 'chef-th' }),
        MockComponent({ selector: 'chef-td' }),
        MockComponent({ selector: 'a', inputs: ['routerLink'] })
      ],
      providers: [
        FeatureFlagsService,
        DestinationRequests,
        HttpClient,
        HttpHandler
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        RouterTestingModule,
        StoreModule.forRoot(ngrxReducers, { runtimeChecks })
      ],
      schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DataFeedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  describe('create data feed', () => {
    let store: Store<NgrxStateAtom>;
    const username = 'test';
    const password = 'test123';
    const token = 'test123';
    const bucketName = 'b123';
    const accessKey = 'test123';
    const secretKey = 'test123';
    const destination = <Destination> {
      id: '1',
      name: 'new data feed',
      secret: 'testSecret',
      url: 'http://foo.com'
    };

    beforeEach(() => {
      store = TestBed.inject(Store);
    });

    it('slidePanel opens slider', () => {
      spyOn(component.createChild, 'slidePanel');
      expect(component.createModalVisible).toBe(false);
      component.slidePanel();
      expect(component.createModalVisible).toBe(true);
    });

    it('on success, closes slider and adds new data feed', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['username'].setValue(username);
      component.createDataFeedForm.controls['password'].setValue(password);
      component.saveDestination({
        auth: AuthTypes.USERNAMEANDPASSWORD,
        name: WebhookIntegrationTypes.SERVICENOW
      });

      store.dispatch(new CreateDestinationSuccess(destination));
      component.sortedDestinations$.subscribe(destinations => {
        expect(destinations).toContain(destination);
      });
    });

    it('on success, closes slider and adds new data feed access token', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['tokenType'].setValue('Bearer');
      component.createDataFeedForm.controls['token'].setValue(token);

      component.saveDestination({
        auth: AuthTypes.ACCESSTOKEN,
        name: WebhookIntegrationTypes.SERVICENOW
      });

      store.dispatch(new CreateDestinationSuccess(destination));
      component.sortedDestinations$.subscribe(destinations => {
        expect(destinations).toContain(destination);
      });
    });

    it('on success, closes slider and adds new data feed minio', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['endpoint'].setValue(destination.url);
      component.createDataFeedForm.controls['bucketName'].setValue(bucketName);
      component.createDataFeedForm.controls['accessKey'].setValue(accessKey);
      component.createDataFeedForm.controls['secretKey'].setValue(secretKey);

      component.saveDestination({
        auth: null,
        name: StorageIntegrationTypes.MINIO
      });

      store.dispatch(new CreateDestinationSuccess(destination));
      component.sortedDestinations$.subscribe(destinations => {
        expect(destinations).toContain(destination);
      });
    });

    it('on conflict error, slider is open with conflict error', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');
      spyOn(component.conflictErrorEvent, 'emit');
      spyOn(component.createChild, 'slidePanel');

      component.slidePanel();
      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['username'].setValue(username);
      component.createDataFeedForm.controls['password'].setValue(password);
      component.saveDestination({
        auth: AuthTypes.USERNAMEANDPASSWORD,
        name: WebhookIntegrationTypes.SERVICENOW
      });

      const conflict = <HttpErrorResponse>{
        status: HttpStatus.CONFLICT,
        ok: false
      };
      store.dispatch(new CreateDestinationFailure(conflict));

      expect(component.createModalVisible).toBe(true);
      expect(component.conflictErrorEvent.emit).toHaveBeenCalled();
    });

    it('on create error, slider is closed with failure banner', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');
      spyOn(component.conflictErrorEvent, 'emit');
      spyOn(component.createChild, 'slidePanel');

      component.slidePanel();
      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['username'].setValue(username);
      component.createDataFeedForm.controls['password'].setValue(password);
      component.saveDestination({
        auth: AuthTypes.USERNAMEANDPASSWORD,
        name: WebhookIntegrationTypes.SERVICENOW
      });

      const error = <HttpErrorResponse>{
        status: HttpStatus.INTERNAL_SERVER_ERROR,
        ok: false
      };
      store.dispatch(new CreateDestinationFailure(error));

      expect(component.createModalVisible).toBe(true);
      expect(component.conflictErrorEvent.emit).toHaveBeenCalledWith(false);
    });

    it('on success, adds new data feed service now and check dispatched store data', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['username'].setValue(username);
      component.createDataFeedForm.controls['password'].setValue(password);

      spyOn(component['store'], 'dispatch');
      component.saveDestination({
        auth: AuthTypes.USERNAMEANDPASSWORD,
        name: WebhookIntegrationTypes.SERVICENOW
      });

      expect(component['store'].dispatch).toHaveBeenCalledWith(
        new CreateDestination({
          name: component.createDataFeedForm.controls['name'].value.trim(),
          url: component.createDataFeedForm.controls['url'].value.trim(),
          integration_types: IntegrationTypes.WEBHOOK,
          services: WebhookIntegrationTypes.SERVICENOW
        }, JSON.stringify({
          Authorization: 'Basic ' + btoa(username + ':' + password)
        }), null)
      );
    });

    it('on success, adds new data feed splunk and check dispatched store data', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['tokenType'].setValue('Bearer');
      component.createDataFeedForm.controls['token'].setValue(token);

      spyOn(component['store'], 'dispatch');
      component.saveDestination({
        auth: AuthTypes.ACCESSTOKEN,
        name: WebhookIntegrationTypes.SPLUNK
      });

      expect(component['store'].dispatch).toHaveBeenCalledWith(
        new CreateDestination({
          name: component.createDataFeedForm.controls['name'].value.trim(),
          url: component.createDataFeedForm.controls['url'].value.trim(),
          integration_types: IntegrationTypes.WEBHOOK,
          services: WebhookIntegrationTypes.SPLUNK
        }, JSON.stringify({
          Authorization: 'Bearer' + ' ' + token
        }), null)
      );
    });

    it('on success, adds new data feed minio and check dispatched store data', () => {
      spyOnProperty(component.createChild, 'saveDone', 'set');

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['endpoint'].setValue(destination.url);
      component.createDataFeedForm.controls['bucketName'].setValue(bucketName);
      component.createDataFeedForm.controls['accessKey'].setValue(accessKey);
      component.createDataFeedForm.controls['secretKey'].setValue(secretKey);

      spyOn(component['store'], 'dispatch');
      component.saveDestination({
        auth: null,
        name: StorageIntegrationTypes.MINIO
      });

      expect(component['store'].dispatch).toHaveBeenCalledWith(
        new CreateDestination({
          name: component.createDataFeedForm.controls['name'].value.trim(),
          url: component.createDataFeedForm.controls['endpoint'].value.trim(),
          integration_types: IntegrationTypes.STORAGE,
          services: StorageIntegrationTypes.MINIO,
          meta_data: [
            {
              key: 'bucket',
              value: component.createDataFeedForm.controls['bucketName'].value.trim()
            }
          ]
        }, null,
        {
          accessKey: component.createDataFeedForm.controls['accessKey'].value.trim(),
          secretKey: component.createDataFeedForm.controls['secretKey'].value.trim()
        })
      );
    });

    it('on test success for Service now and check dispatched store data', () => {

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['username'].setValue(username);
      component.createDataFeedForm.controls['password'].setValue(password);

      spyOn(component['datafeedRequests'], 'testDestinationWithUsernamePassword');
      component.sendTestForDataFeed({
        auth: AuthTypes.USERNAMEANDPASSWORD,
        name: WebhookIntegrationTypes.SERVICENOW
      });

      expect(component['datafeedRequests']
      .testDestinationWithUsernamePassword).toHaveBeenCalledWith(
        destination.url,
        username,
        password
      );
    });

    it('on test success for Splunk and check dispatched store data', () => {

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['url'].setValue(destination.url);
      component.createDataFeedForm.controls['tokenType'].setValue('Bearer');
      component.createDataFeedForm.controls['token'].setValue(token);

      spyOn(component['datafeedRequests'], 'testDestinationWithHeaders');
      component.sendTestForDataFeed({
        auth: AuthTypes.ACCESSTOKEN,
        name: WebhookIntegrationTypes.SPLUNK
      });

      expect(component['datafeedRequests']
      .testDestinationWithHeaders).toHaveBeenCalledWith(
        destination.url,
        JSON.stringify({
          Authorization: 'Bearer' + ' ' + token
        })
      );
    });

    it('on test success for Minio and check dispatched store data', () => {

      component.createDataFeedForm.controls['name'].setValue(destination.name);
      component.createDataFeedForm.controls['endpoint'].setValue(destination.url);
      component.createDataFeedForm.controls['bucketName'].setValue(bucketName);
      component.createDataFeedForm.controls['accessKey'].setValue(accessKey);
      component.createDataFeedForm.controls['secretKey'].setValue(secretKey);

      spyOn(component['datafeedRequests'], 'testDestinationForMinio');
      component.sendTestForDataFeed({
        auth: null,
        name: StorageIntegrationTypes.MINIO
      });

      expect(component['datafeedRequests']
      .testDestinationForMinio).toHaveBeenCalledWith({
        url: destination.url,
        aws: {
          access_key: component.createDataFeedForm.controls['accessKey'].value.trim(),
          secret_access_key: component.createDataFeedForm.controls['secretKey'].value.trim(),
          bucket: component.createDataFeedForm.controls['bucketName'].value.trim()
        }
      });
    });

  });

});
